// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

func TestGetCertificatePEMBytes(t *testing.T) {
	privKey, certData, err := newMockCertificate()
	if err != nil {
		t.Fatalf("Failed to create mock certificate: %v", err)
	}

	certDER, err := x509.CreateCertificate(rand.Reader, certData, certData, &privKey.PublicKey, privKey)
	if err != nil {
		t.Fatalf("Failed to create DER certificate: %v", err)
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}

	modelCert := appmodels.NewCertificate("Org123", []appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(1))}, cert)

	pemBytes := apputils.GetCertificatePEMBytes(modelCert)

	// Decode the PEM to verify it's correctly encoded
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		t.Fatal("Failed to decode PEM block")
	}

	if block.Type != "CERTIFICATE" {
		t.Errorf("PEM block type is %v, want CERTIFICATE", block.Type)
	}

	reParsedCert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		t.Fatalf("Failed to parse certificate from PEM: %v", err)
	}

	// Compare some fields to ensure the certificate is correctly encoded/decoded
	if reParsedCert.SerialNumber.Cmp(cert.SerialNumber) != 0 {
		t.Errorf("SerialNumber mismatch, got %v, want %v", reParsedCert.SerialNumber, cert.SerialNumber)
	}
	if reParsedCert.Subject.CommonName != cert.Subject.CommonName {
		t.Errorf("CommonName mismatch, got %s, want %s", reParsedCert.Subject.CommonName, cert.Subject.CommonName)
	}
}

func TestGetCertificateDTO(t *testing.T) {

	_, certData, err := newMockCertificate()
	if err != nil {
		t.Fatalf("Failed to create mock certificate: %v", err)
	}

	// Create the certificate model instance
	cert := appmodels.NewCertificate("Org123", []appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(1))}, certData)

	// Generate PEM for comparison
	pemBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certData.Raw,
	}
	pemBytes := pem.EncodeToMemory(pemBlock)

	// Call GetDTO and verify each field
	dto := apputils.ToCertificateDTO(cert)

	if dto.CommonName != cert.GetCommonName() {
		t.Errorf("DTO CommonName mismatch, got %v, want %v", dto.CommonName, cert.GetCommonName())
	}

	if dto.SerialNumber != cert.GetSerialNumber().String() {
		t.Errorf("DTO SerialNumber mismatch, got %v, want %v", dto.SerialNumber, cert.GetSerialNumber().String())
	}

	if dto.SignedBy != cert.GetSignedBy().String() {
		t.Errorf("DTO SignedBy mismatch, got %v, want %v", dto.SignedBy, cert.GetSignedBy().String())
	}

	if dto.Organization != cert.GetOrganizationName() {
		t.Errorf("DTO OrganizationName mismatch, got %v, want %v", dto.Organization, cert.GetOrganizationName())
	}

	if dto.IsCA != cert.IsCA() {
		t.Errorf("DTO IsCA mismatch, got %v, want %v", dto.IsCA, cert.IsCA())
	}

	if dto.IsRootCertificate != cert.IsRootCertificate() {
		t.Errorf("DTO IsRootCertificate mismatch, got %v, want %v", dto.IsRootCertificate, cert.IsRootCertificate())
	}

	if dto.IsIntermediateCertificate != cert.IsIntermediateCertificate() {
		t.Errorf("DTO IsIntermediateCertificate mismatch, got %v, want %v", dto.IsIntermediateCertificate, cert.IsIntermediateCertificate())
	}

	if dto.IsServerCertificate != cert.IsServerCertificate() {
		t.Errorf("DTO IsServerCertificate mismatch, got %v, want %v", dto.IsServerCertificate, cert.IsServerCertificate())
	}

	if dto.IsClientCertificate != cert.IsClientCertificate() {
		t.Errorf("DTO IsClientCertificate mismatch, got %v, want %v", dto.IsClientCertificate, cert.IsClientCertificate())
	}

	if string(dto.Certificate) != string(pemBytes) {
		t.Errorf("DTO PEM mismatch, got %v, want %v", string(dto.Certificate), string(pemBytes))
	}
}

// Helper function to create a new RSA key and a mock x509.Certificate
func newMockCertificate() (*rsa.PrivateKey, *x509.Certificate, error) {
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "example.com",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(365 * 24 * time.Hour), // Valid for one year
		KeyUsage:  x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		IsCA:      true,
	}

	return privKey, cert, nil
}

func TestCreateSignedCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	template := &x509.Certificate{}
	signingCert := &x509.Certificate{}
	publicKey, privateKey := "publicKey", "privateKey"

	// Setup mock expectations
	certBytes := []byte("mockCertBytes")
	mockCert := &x509.Certificate{SerialNumber: big.NewInt(1234)}
	mockManager.On("CreateCertificate", mock.Anything, template, signingCert, publicKey, privateKey).Return(certBytes, nil)
	mockManager.On("ParseCertificate", certBytes).Return(mockCert, nil)

	// Test the function
	cert, err := apputils.CreateSignedCertificate(mockManager, template, signingCert, publicKey, privateKey)
	if err != nil {
		t.Fatalf("CreateSignedCertificate failed: %v", err)
	}
	if cert.SerialNumber.Cmp(big.NewInt(1234)) != 0 {
		t.Errorf("Expected serial number 1234, got %v", cert.SerialNumber)
	}
}

func TestCreateSignedCertificate_NilManager(t *testing.T) {
	_, err := apputils.CreateSignedCertificate(
		nil, // manager is nil
		&x509.Certificate{},
		&x509.Certificate{},
		"publicKey",
		"privateKey",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "manager: must be defined")
}

func TestNewIntermediateCertificate(t *testing.T) {

	// Mock the certificate manager
	mockManager := &commonmocks.MockCertificateManager{}

	// Mock inputs
	parentSerialNumber := big.NewInt(10)
	serialNumber := big.NewInt(100)
	organization := &appmocks.MockOrganization{}
	expiration := 365 * 24 * time.Hour
	parentCertificate := &appmocks.MockCertificate{}
	parentPrivateKey := &appmocks.MockPrivateKey{}
	publicKey := &appmocks.MockPublicKey{}
	commonName := "Intermediate CA"

	organizationId := "TestOrg"

	organization.On("GetID").Return(organizationId)
	organization.On("GetName").Return("Test Org")
	organization.On("GetNames").Return([]string{"Test Org"})

	publicKey.On("GetPublicKey").Return(&rsa.PublicKey{})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})
	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	resultCert := &x509.Certificate{SerialNumber: serialNumber}

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.AnythingOfType("*x509.Certificate"), mock.Anything, mock.Anything, mock.Anything).Return([]byte("certBytes"), nil)
	mockManager.On("ParseCertificate", mock.Anything).Return(resultCert, nil)

	// Execute the function under test
	cert, err := apputils.NewIntermediateCertificate(
		mockManager,
		appmodels.NewSerialNumber(serialNumber),
		organization,
		expiration,
		publicKey,
		parentCertificate,
		parentPrivateKey,
		commonName,
	)

	// Assert expectations
	assert.NoError(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, resultCert, cert.GetCertificate())
	assert.Equal(t, organizationId, cert.GetOrganizationID())
	assert.Equal(t, 1, len(cert.GetParents()))
	// assert.Equal(t, cert.GetParents()[0], )
	mockManager.AssertExpectations(t)
}

func TestNewServerCertificate(t *testing.T) {
	// Mock the certificate manager
	mockManager := &commonmocks.MockCertificateManager{}

	// Mock inputs
	parentSerialNumber := big.NewInt(20)
	serialNumber := big.NewInt(200)
	organization := &appmocks.MockOrganization{}
	expiration := 365 * 24 * time.Hour
	parentCertificate := &appmocks.MockCertificate{}
	mockPublicKey := &appmocks.MockPublicKey{}
	mockPublicKey.On("GetPublicKey").Return(&rsa.PublicKey{})

	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "server.example.com"
	dnsNames := []string{"www.example.com", "example.com"}

	organizationId := "TestOrgServer"

	organization.On("GetID").Return(organizationId)
	organization.On("GetName").Return("Test Org Server")
	organization.On("GetNames").Return([]string{"Test Org Server"})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})
	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	resultCert := &x509.Certificate{SerialNumber: serialNumber, DNSNames: dnsNames}

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.AnythingOfType("*x509.Certificate"), mock.Anything, mock.Anything, mock.Anything).Return([]byte("certBytes"), nil)
	mockManager.On("ParseCertificate", mock.Anything).Return(resultCert, nil)

	// Execute the function under test
	cert, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(serialNumber),
		organization,
		expiration,
		mockPublicKey,
		parentCertificate,
		parentPrivateKey,
		commonName,
		dnsNames...,
	)

	// Assert expectations
	assert.NoError(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, resultCert, cert.GetCertificate())
	assert.Equal(t, organizationId, cert.GetOrganizationID())
	assert.ElementsMatch(t, dnsNames, cert.GetCertificate().DNSNames) // Ensure the DNS names match
	mockManager.AssertExpectations(t)
}

func TestNewClientCertificate(t *testing.T) {
	// Mock the certificate manager
	mockManager := &commonmocks.MockCertificateManager{}

	// Mock inputs
	parentSerialNumber := big.NewInt(30)
	serialNumber := big.NewInt(300)
	organization := &appmocks.MockOrganization{}
	expiration := 365 * 24 * time.Hour
	mockPublicKey := appmocks.NewMockRsaPublicKey()
	parentCertificate := &appmocks.MockCertificate{}
	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "Client Certificate"

	organizationID := "TestOrgClient"

	organization.On("GetID").Return(organizationID)
	organization.On("GetName").Return("Test Org Client")
	organization.On("GetNames").Return([]string{"Test Org Client"})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})

	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	resultCert := &x509.Certificate{SerialNumber: serialNumber, ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.AnythingOfType("*x509.Certificate"), mock.Anything, mock.Anything, mock.Anything).Return([]byte("certBytes"), nil)
	mockManager.On("ParseCertificate", mock.Anything).Return(resultCert, nil)

	// Execute the function under test
	cert, err := apputils.NewClientCertificate(
		mockManager,
		appmodels.NewSerialNumber(serialNumber),
		organization,
		expiration,
		mockPublicKey,
		parentCertificate,
		parentPrivateKey,
		commonName,
	)

	// Assert expectations
	assert.NoError(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, resultCert, cert.GetCertificate())
	assert.Equal(t, organizationID, cert.GetOrganizationID())
	assert.Contains(t, cert.GetCertificate().ExtKeyUsage, x509.ExtKeyUsageClientAuth) // Ensure the ExtKeyUsage includes ClientAuth
	mockManager.AssertExpectations(t)
}

func TestNewRootCertificate_Success(t *testing.T) {
	// Setup mock manager, organization, and private key
	mockManager := new(commonmocks.MockCertificateManager)
	mockOrganization := new(appmocks.MockOrganization)
	mockPrivateKey := new(appmocks.MockPrivateKey)
	parentCertificate := new(appmocks.MockCertificate)

	parentSerialNumber := big.NewInt(100)
	serialNumber := appmodels.NewSerialNumber(big.NewInt(1))
	organizationID := "TestOrg"
	expiration := 365 * 24 * time.Hour
	commonName := "Test Root CA"

	expectedCert := &x509.Certificate{ /* expected certificate details */ }

	organizationId := "TestOrg"

	mockOrganization.On("GetID").Return(organizationId)
	mockOrganization.On("GetName").Return("Test Org")
	mockOrganization.On("GetNames").Return([]string{"Test Org"})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})

	mockPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	mockPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	// Mock organization behavior
	mockOrganization.On("GetID").Return(organizationID)
	mockOrganization.On("GetNames").Return([]string{"Test Organization"})

	// Mock certificate manager behavior
	mockManager.On("CreateCertificate", mock.Anything, mock.AnythingOfType("*x509.Certificate"), mock.Anything, mock.Anything, mock.Anything).Return([]byte("certBytes"), nil)
	mockManager.On("ParseCertificate", []byte("certBytes")).Return(expectedCert, nil)

	// Attempt to create a new root certificate
	cert, err := apputils.NewRootCertificate(
		mockManager,
		serialNumber,
		mockOrganization,
		expiration,
		mockPrivateKey,
		commonName,
	)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, cert)
	assert.Equal(t, expectedCert, cert.GetCertificate(), "The generated certificate should match the expected certificate")
	mockManager.AssertExpectations(t)
}

func TestCreateSignedCertificate_CreateCertificateFails(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)

	expectedError := fmt.Errorf("create certificate error")

	// Ensure you return nil for the byte slice when simulating a failure
	mockManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, expectedError)

	template := &x509.Certificate{}

	_, err := apputils.CreateSignedCertificate(
		mockManager,
		template,
		template, // signingCertificate
		nil,
		nil,
	)

	// The focus is on error handling, so check that an error was returned and matches the expected error
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}
	if !strings.Contains(err.Error(), expectedError.Error()) {
		t.Errorf("Expected error to contain %q, got %q", expectedError.Error(), err.Error())
	}

	// Verify that the mock expectations were met
	mockManager.AssertExpectations(t)
}

func TestCreateSignedCertificate_ParseCertificateFails(t *testing.T) {
	// Create a mock certificate manager
	mockManager := new(commonmocks.MockCertificateManager)

	// Setup the mock to return success for CreateCertificate but fail for ParseCertificate
	certBytes := []byte("dummy certificate bytes")
	parseError := fmt.Errorf("parse certificate error")
	mockManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(certBytes, nil)
	mockManager.On("ParseCertificate", certBytes).Return(nil, parseError)

	// Call CreateSignedCertificate with dummy arguments
	_, err := apputils.CreateSignedCertificate(
		mockManager,
		&x509.Certificate{}, // template
		&x509.Certificate{}, // signingCertificate
		nil,                 // signingPublicKey
		nil,                 // signingPrivateKey
	)

	// Verify that the error is what we expect
	if err == nil {
		t.Fatalf("Expected an error, but got nil")
	}

	// Check if the error message matches the expected error
	if !strings.Contains(err.Error(), parseError.Error()) {
		t.Errorf("Expected error message to contain %q, got %q", parseError.Error(), err.Error())
	}
}

func TestNewIntermediateCertificate_NilManager(t *testing.T) {
	_, err := apputils.NewIntermediateCertificate(
		nil, // manager is nil
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Intermediate CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "manager: must be defined")
}

func TestNewIntermediateCertificate_NilSerialNumber(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewIntermediateCertificate(
		mockManager,
		nil, // serialNumber is nil
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Intermediate CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "serialNumber: must be defined")
}

func TestNewIntermediateCertificate_NilOrganization(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewIntermediateCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		nil, // organization is nil
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Intermediate CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization: must be defined")
}

func TestNewIntermediateCertificate_NilParentCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewIntermediateCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		nil, // parentCertificate is nil
		&appmocks.MockPrivateKey{},
		"Intermediate CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentCertificate: must be defined")
}

func TestNewIntermediateCertificate_NilParentPrivateKey(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewIntermediateCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		nil, // parentPrivateKey is nil
		"Intermediate CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentPrivateKey: must be defined")
}

func TestNewIntermediateCertificate_EmptyCommonName(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewIntermediateCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"", // commonName is empty
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "commonName: cannot be empty")
}

func TestNewIntermediateCertificate_FailingCreateCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)

	parentSerialNumber := big.NewInt(10)
	organization := &appmocks.MockOrganization{}
	parentCertificate := &appmocks.MockCertificate{}
	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "Intermediate CA"

	organizationId := "TestOrg"

	organization.On("GetID").Return(organizationId)
	organization.On("GetName").Return("Test Org")
	organization.On("GetNames").Return([]string{"Test Org"})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})
	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("creation error"))

	_, err := apputils.NewIntermediateCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		organization,
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		parentCertificate,
		parentPrivateKey,
		commonName,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "creation error")
}

func TestNewServerCertificate_NilManager(t *testing.T) {
	_, err := apputils.NewServerCertificate(
		nil, // manager is nil
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Server Certificate",
		"www.example.com",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "manager: must be defined")
}

func TestNewServerCertificate_NilSerialNumber(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewServerCertificate(
		mockManager,
		nil, // serialNumber is nil
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Server Certificate",
		"www.example.com",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "serialNumber: must be defined")
}

func TestNewServerCertificate_NilOrganization(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		nil, // organization is nil
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Server Certificate",
		"www.example.com",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization: must be defined")
}

func TestNewServerCertificate_NilParentCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		nil, // parentCertificate is nil
		&appmocks.MockPrivateKey{},
		"Server Certificate",
		"www.example.com",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentCertificate: must be defined")
}

func TestNewServerCertificate_NilParentPrivateKey(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		nil, // parentPrivateKey is nil
		"Server Certificate",
		"www.example.com",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentPrivateKey: must be defined")
}

func TestNewServerCertificate_EmptyCommonName(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"", // commonName is empty
		"www.example.com",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "commonName: cannot be empty")
}

func TestNewServerCertificate_NilDnsNames(t *testing.T) {
	var dnsNames []string = nil
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"example.com",
		dnsNames..., // dnsNames is nil
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dnsNames: must be defined")
}

func TestNewServerCertificate_FailingCreateCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)

	parentSerialNumber := big.NewInt(10)
	organization := &appmocks.MockOrganization{}
	parentCertificate := &appmocks.MockCertificate{}
	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "my.example.com"

	organizationId := "TestOrg"

	organization.On("GetID").Return(organizationId)
	organization.On("GetName").Return("Test Org")
	organization.On("GetNames").Return([]string{"Test Org"})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})
	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("creation error"))

	_, err := apputils.NewServerCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		organization,
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		parentCertificate,
		parentPrivateKey,
		commonName,
		commonName,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "creation error")
}

func TestNewClientCertificate_FailingCreateCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)

	parentSerialNumber := big.NewInt(10)
	organization := &appmocks.MockOrganization{}
	parentCertificate := &appmocks.MockCertificate{}
	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "Client CA"

	organizationId := "TestOrg"

	organization.On("GetID").Return(organizationId)
	organization.On("GetName").Return("Test Org")
	organization.On("GetNames").Return([]string{"Test Org"})

	parentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(parentSerialNumber))
	parentCertificate.On("GetCertificate").Return(&x509.Certificate{SerialNumber: parentSerialNumber})
	parentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})
	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("creation error"))

	_, err := apputils.NewClientCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		organization,
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		parentCertificate,
		parentPrivateKey,
		commonName,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "creation error")
}

func TestNewRootCertificate_FailingCreateCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)

	organization := &appmocks.MockOrganization{}
	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "Root CA"

	organizationId := "TestOrg"

	organization.On("GetID").Return(organizationId)
	organization.On("GetName").Return("Test Org")
	organization.On("GetNames").Return([]string{"Test Org"})

	parentPrivateKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	parentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})

	// Set expectations on the mock manager
	mockManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, fmt.Errorf("creation error"))

	_, err := apputils.NewRootCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		organization,
		365*24*time.Hour,
		parentPrivateKey,
		commonName,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "creation error")
}

func TestNewClientCertificate_NilManager(t *testing.T) {
	_, err := apputils.NewClientCertificate(
		nil, // manager is nil
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Client CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "manager: must be defined")
}

func TestNewClientCertificate_NilSerialNumber(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewClientCertificate(
		mockManager,
		nil, // serialNumber is nil
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Client CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "serialNumber: must be defined")
}

func TestNewClientCertificate_NilOrganization(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewClientCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		nil, // organization is nil
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"Client CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization: must be defined")
}

func TestNewClientCertificate_NilParentCertificate(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewClientCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		nil, // parentCertificate is nil
		&appmocks.MockPrivateKey{},
		"Client CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentCertificate: must be defined")
}

func TestNewClientCertificate_NilParentPrivateKey(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewClientCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		nil, // parentPrivateKey is nil
		"Client CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentPrivateKey: must be defined")
}

func TestNewClientCertificate_EmptyCommonName(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewClientCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		appmocks.NewMockRsaPublicKey(),
		&appmocks.MockCertificate{},
		&appmocks.MockPrivateKey{},
		"", // commonName is empty
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "commonName: cannot be empty")
}

func TestNewRootCertificate_NilManager(t *testing.T) {
	_, err := apputils.NewRootCertificate(
		nil, // manager is nil
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		&appmocks.MockPrivateKey{},
		"Root CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "manager: must be defined")
}

func TestNewRootCertificate_NilSerialNumber(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewRootCertificate(
		mockManager,
		nil, // serialNumber is nil
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		&appmocks.MockPrivateKey{},
		"Root CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "serialNumber: must be defined")
}

func TestNewRootCertificate_NilOrganization(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewRootCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		nil, // organization is nil
		365*24*time.Hour,
		&appmocks.MockPrivateKey{},
		"Root CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization: must be defined")
}

func TestNewRootCertificate_NilParentPrivateKey(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewRootCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		nil, // parentPrivateKey is nil
		"Root CA",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "privateKey: must be defined")
}

func TestNewRootCertificate_EmptyCommonName(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, err := apputils.NewRootCertificate(
		mockManager,
		appmodels.NewSerialNumber(big.NewInt(1)),
		&appmocks.MockOrganization{},
		365*24*time.Hour,
		&appmocks.MockPrivateKey{},
		"", // commonName is empty
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "commonName: cannot be empty")
}

func TestToCertificateRevokedDTO(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(big.NewInt(1234))
	serialNumberString := serialNumber.String()
	revocationTime := time.Now()
	expirationTime := revocationTime.Add(365 * 24 * time.Hour)

	mockRevokedCert := new(appmocks.MockRevokedCertificate)
	mockRevokedCert.On("GetSerialNumber").Return(serialNumber)
	mockRevokedCert.On("GetRevocationTime").Return(revocationTime)
	mockRevokedCert.On("GetExpirationTime").Return(expirationTime)

	dto := apputils.ToCertificateRevokedDTO(mockRevokedCert)

	assert.Equal(t, serialNumberString, dto.SerialNumber, "Serial numbers should match")
	assert.Equal(t, revocationTime, dto.RevocationTime, "Revocation times should match")
	assert.Equal(t, expirationTime, dto.ExpirationTime, "Expiration times should match")
}

func TestToRevokedCertificate(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(big.NewInt(1234))
	serialNumberString := serialNumber.String()
	notAfter := time.Now().Add(365 * 24 * time.Hour)
	revocationTime := time.Now()

	mockCert := new(appmocks.MockCertificate)
	mockCert.On("GetSerialNumber").Return(serialNumber)
	mockCert.On("NotAfter").Return(notAfter)

	revokedCert := apputils.ToRevokedCertificate(mockCert, revocationTime)

	assert.Equal(t, serialNumberString, revokedCert.GetSerialNumber().String(), "Serial numbers should match")
	assert.Equal(t, revocationTime, revokedCert.GetRevocationTime(), "Revocation times should match")
	assert.Equal(t, notAfter, revokedCert.GetExpirationTime(), "Expiration times should match")
}

func TestToCertificateCreatedDTO(t *testing.T) {
	var privKey = &rsa.PrivateKey{}
	var der []byte = []byte{1, 2, 3}
	var pemData []byte = []byte("PRIVATE_KEY_PEM")
	mockCertManager := new(commonmocks.MockCertificateManager)
	mockCertificate := new(appmocks.MockCertificate)
	mockPrivateKey := new(appmocks.MockPrivateKey)

	mockPrivateKey.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(123456789)))
	mockPrivateKey.On("GetKeyType").Return(appmodels.RSA_2048)
	mockPrivateKey.On("GetPrivateKey").Return(privKey)

	// Setup mock certificate behavior
	mockCertificate.On("GetCertificate").Return(&x509.Certificate{})
	mockCertificate.On("GetParents").Return([]appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(987654321))})
	mockCertificate.On("GetCommonName").Return("www.example.com")
	mockCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(123456789)))
	mockCertificate.On("GetSignedBy").Return(appmodels.NewSerialNumber(big.NewInt(987654321)))
	mockCertificate.On("GetOrganizationName").Return("Example Org")
	mockCertificate.On("IsCA").Return(false)
	mockCertificate.On("IsRootCertificate").Return(false)
	mockCertificate.On("IsIntermediateCertificate").Return(false)
	mockCertificate.On("IsServerCertificate").Return(true)
	mockCertificate.On("IsClientCertificate").Return(false)

	// Setup mock private key behavior
	privateKeyDTO := appdtos.PrivateKeyDTO{
		Certificate: "123456789",
		Type:        "RSA_2048",
		PrivateKey:  "PRIVATE_KEY_PEM",
	}

	mockCertManager.On("MarshalPKCS1PrivateKey", privKey).Return(der)
	mockCertManager.On("EncodePEMToMemory", mock.Anything).Return(pemData)

	mockCertManager.On("ToPrivateKeyDTO", mock.Anything).Return(privateKeyDTO, nil)

	// Successful conversion
	expectedDTO := appdtos.CertificateCreatedDTO{
		Certificate: appdtos.CertificateDTO{
			CommonName:                "www.example.com",
			SerialNumber:              "123456789",
			Parents:                   []string{"987654321"},
			SignedBy:                  "987654321",
			Organization:              "Example Org",
			IsCA:                      false,
			IsRootCertificate:         false,
			IsIntermediateCertificate: false,
			IsServerCertificate:       true,
			IsClientCertificate:       false,
			Certificate:               "-----BEGIN CERTIFICATE-----\n-----END CERTIFICATE-----\n",
		},
		PrivateKey: privateKeyDTO,
	}

	result, err := apputils.ToCertificateCreatedDTO(mockCertManager, mockCertificate, mockPrivateKey)
	assert.NoError(t, err)
	assert.Equal(t, expectedDTO, result)

	// Test error cases
	t.Run("NilCertManager", func(t *testing.T) {
		_, err := apputils.ToCertificateCreatedDTO(nil, mockCertificate, mockPrivateKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cert manager not defined")
	})

	t.Run("NilCertificate", func(t *testing.T) {
		_, err := apputils.ToCertificateCreatedDTO(mockCertManager, nil, mockPrivateKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "certificate not defined")
	})

	t.Run("NilPrivateKey", func(t *testing.T) {
		_, err := apputils.ToCertificateCreatedDTO(mockCertManager, mockCertificate, nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "private key not defined")
	})

	t.Run("ToPrivateKeyDTOError", func(t *testing.T) {

		mockCertManager.Mock = mock.Mock{}

		var nilbyted []byte = nil
		key2 := &rsa.PrivateKey{}
		mockPrivateKey.On("GetPrivateKey").Return(key2)

		mockCertManager.On("MarshalPKCS1PrivateKey", privKey).Return(nilbyted)
		_, err := apputils.ToCertificateCreatedDTO(mockCertManager, mockCertificate, mockPrivateKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ToCertificateCreatedDTO: failed")
	})
}

func TestToListOfCertificateDTO(t *testing.T) {
	// Setup mock certificates
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)

	commonName1 := "www.example.com"
	commonName2 := "www.test.com"

	mockCert1.On("GetCommonName").Return(commonName1)
	mockCert1.On("GetCertificate").Return(&x509.Certificate{})
	mockCert1.On("GetParents").Return([]appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(987654321))})
	mockCert1.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(123456789)))
	mockCert1.On("GetSignedBy").Return(appmodels.NewSerialNumber(big.NewInt(987654321)))
	mockCert1.On("GetOrganizationName").Return("Example Org")
	mockCert1.On("IsCA").Return(false)
	mockCert1.On("IsRootCertificate").Return(false)
	mockCert1.On("IsIntermediateCertificate").Return(false)
	mockCert1.On("IsServerCertificate").Return(true)
	mockCert1.On("IsClientCertificate").Return(false)

	mockCert2.On("GetCommonName").Return(commonName2)
	mockCert2.On("GetCertificate").Return(&x509.Certificate{})
	mockCert2.On("GetParents").Return([]appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(987654321))})
	mockCert2.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(123456789)))
	mockCert2.On("GetSignedBy").Return(appmodels.NewSerialNumber(big.NewInt(987654321)))
	mockCert2.On("GetOrganizationName").Return("Example Org")
	mockCert2.On("IsCA").Return(false)
	mockCert2.On("IsRootCertificate").Return(false)
	mockCert2.On("IsIntermediateCertificate").Return(false)
	mockCert2.On("IsServerCertificate").Return(true)
	mockCert2.On("IsClientCertificate").Return(false)

	// Assume other necessary mocks here as per ToCertificateDTO usage

	list := []appmodels.ICertificate{mockCert1, mockCert2}

	// Call function
	result := apputils.ToListOfCertificateDTO(list)

	// Assert results
	assert.Len(t, result, 2)
	assert.Equal(t, commonName1, result[0].CommonName)
	assert.Equal(t, commonName2, result[1].CommonName)

	// Ensure mock expectations are met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
}

func TestFilterRootCertificates(t *testing.T) {
	// Setup mock certificates
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)
	mockCert3 := new(appmocks.MockCertificate)

	// Mock behavior: Only mockCert1 and mockCert3 are root certificates
	mockCert1.On("IsRootCertificate").Return(true)
	mockCert2.On("IsRootCertificate").Return(false)
	mockCert3.On("IsRootCertificate").Return(true)

	certificates := []appmodels.ICertificate{mockCert1, mockCert2, mockCert3}

	// Call the function under test
	filteredCerts := apputils.FilterRootCertificates(certificates)

	// Assert that the result contains only the root certificates (mockCert1 and mockCert3)
	assert.Len(t, filteredCerts, 2, "Expected two certificates in the filtered list")
	assert.Contains(t, filteredCerts, mockCert1, "Filtered certificates should include mockCert1")
	assert.Contains(t, filteredCerts, mockCert3, "Filtered certificates should include mockCert3")
	assert.NotContains(t, filteredCerts, mockCert2, "Filtered certificates should not include mockCert2")

	// Ensure mock expectations are met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
	mockCert3.AssertExpectations(t)
}

func TestFilterClientCertificates(t *testing.T) {
	// Setup mock certificates
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)
	mockCert3 := new(appmocks.MockCertificate)

	// Mock behavior: Only mockCert2 is a client certificate
	mockCert1.On("IsClientCertificate").Return(false)
	mockCert2.On("IsClientCertificate").Return(true)
	mockCert3.On("IsClientCertificate").Return(false)

	certificates := []appmodels.ICertificate{mockCert1, mockCert2, mockCert3}

	// Call the function under test
	filteredCerts := apputils.FilterClientCertificates(certificates)

	// Assert that the result contains only the client certificates (mockCert2)
	assert.Len(t, filteredCerts, 1, "Expected one certificate in the filtered list")
	assert.Contains(t, filteredCerts, mockCert2, "Filtered certificates should include mockCert2")

	// Ensure mock expectations are met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
	mockCert3.AssertExpectations(t)
}

func TestFilterServerCertificates(t *testing.T) {
	// Setup mock certificates
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)
	mockCert3 := new(appmocks.MockCertificate)

	// Mock behavior: Only mockCert3 is a server certificate
	mockCert1.On("IsServerCertificate").Return(false)
	mockCert2.On("IsServerCertificate").Return(false)
	mockCert3.On("IsServerCertificate").Return(true)

	certificates := []appmodels.ICertificate{mockCert1, mockCert2, mockCert3}

	// Call the function under test
	filteredCerts := apputils.FilterServerCertificates(certificates)

	// Assert that the result contains only the server certificates (mockCert3)
	assert.Len(t, filteredCerts, 1, "Expected one certificate in the filtered list")
	assert.Contains(t, filteredCerts, mockCert3, "Filtered certificates should include mockCert3")

	// Ensure mock expectations are met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
	mockCert3.AssertExpectations(t)
}

func TestFilterIntermediateCertificates(t *testing.T) {
	// Setup mock certificates
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)
	mockCert3 := new(appmocks.MockCertificate)

	// Mock behavior: mockCert1 and mockCert3 are intermediate certificates
	mockCert1.On("IsIntermediateCertificate").Return(true)
	mockCert2.On("IsIntermediateCertificate").Return(false)
	mockCert3.On("IsIntermediateCertificate").Return(true)

	certificates := []appmodels.ICertificate{mockCert1, mockCert2, mockCert3}

	// Call the function under test
	filteredCerts := apputils.FilterIntermediateCertificates(certificates)

	// Assert that the result contains only the intermediate certificates (mockCert1 and mockCert3)
	assert.Len(t, filteredCerts, 2, "Expected two certificates in the filtered list")
	assert.Contains(t, filteredCerts, mockCert1, "Filtered certificates should include mockCert1")
	assert.Contains(t, filteredCerts, mockCert3, "Filtered certificates should include mockCert3")

	// Ensure mock expectations are met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
	mockCert3.AssertExpectations(t)
}

func TestFilterCertificatesByType(t *testing.T) {
	// Setup mock certificates
	mockRootCert := new(appmocks.MockCertificate)
	mockClientCert := new(appmocks.MockCertificate)
	mockServerCert := new(appmocks.MockCertificate)
	mockIntermediateCert := new(appmocks.MockCertificate)

	// Setup the behavior for the mock certificates
	mockRootCert.On("IsRootCertificate").Return(true)
	mockClientCert.On("IsClientCertificate").Return(true)
	mockServerCert.On("IsServerCertificate").Return(true)
	mockIntermediateCert.On("IsIntermediateCertificate").Return(true)

	// Other types return false
	mockRootCert.On("IsClientCertificate").Return(false)
	mockRootCert.On("IsServerCertificate").Return(false)
	mockRootCert.On("IsIntermediateCertificate").Return(false)

	mockClientCert.On("IsRootCertificate").Return(false)
	mockClientCert.On("IsServerCertificate").Return(false)
	mockClientCert.On("IsIntermediateCertificate").Return(false)

	mockServerCert.On("IsRootCertificate").Return(false)
	mockServerCert.On("IsClientCertificate").Return(false)
	mockServerCert.On("IsIntermediateCertificate").Return(false)

	mockIntermediateCert.On("IsRootCertificate").Return(false)
	mockIntermediateCert.On("IsClientCertificate").Return(false)
	mockIntermediateCert.On("IsServerCertificate").Return(false)

	certificates := []appmodels.ICertificate{mockRootCert, mockClientCert, mockServerCert, mockIntermediateCert}

	tests := []struct {
		name             string
		certificateType  string
		expectedFiltered []appmodels.ICertificate
	}{
		{"Root", "root", []appmodels.ICertificate{mockRootCert}},
		{"Client", "client", []appmodels.ICertificate{mockClientCert}},
		{"Server", "server", []appmodels.ICertificate{mockServerCert}},
		{"Intermediate", "intermediate", []appmodels.ICertificate{mockIntermediateCert}},
		{"Unknown", "unknown", []appmodels.ICertificate{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filteredCerts := apputils.FilterCertificatesByType(certificates, tt.certificateType)
			assert.Equal(t, tt.expectedFiltered, filteredCerts, "Filtered certificates should match expected")
		})
	}

	// Ensure mock expectations are met
	mockRootCert.AssertExpectations(t)
	mockClientCert.AssertExpectations(t)
	mockServerCert.AssertExpectations(t)
	mockIntermediateCert.AssertExpectations(t)
}

func TestToCertificateListDTO(t *testing.T) {
	// Setup mock certificates
	mockCert1 := new(appmocks.MockCertificate)
	mockCert2 := new(appmocks.MockCertificate)

	// Mock behaviors for the certificates
	mockCert1.On("GetCommonName").Return("www.example1.com")
	mockCert1.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(100)))
	mockCert1.On("GetCertificate").Return(&x509.Certificate{})
	mockCert1.On("GetParents").Return([]appmodels.ISerialNumber{})
	mockCert1.On("IsCA").Return(false)
	mockCert1.On("IsRootCertificate").Return(false)
	mockCert1.On("IsIntermediateCertificate").Return(false)
	mockCert1.On("IsServerCertificate").Return(true)
	mockCert1.On("IsClientCertificate").Return(false)
	mockCert1.On("GetSignedBy").Return(appmodels.NewSerialNumber(big.NewInt(987654321)))
	mockCert1.On("GetOrganizationName").Return("Example Org")

	mockCert2.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(100)))
	mockCert2.On("GetCommonName").Return("www.example2.com")
	mockCert2.On("GetCertificate").Return(&x509.Certificate{})
	mockCert2.On("GetParents").Return([]appmodels.ISerialNumber{})
	mockCert2.On("IsCA").Return(false)
	mockCert2.On("IsRootCertificate").Return(false)
	mockCert2.On("IsIntermediateCertificate").Return(true)
	mockCert2.On("IsServerCertificate").Return(false)
	mockCert2.On("IsClientCertificate").Return(false)
	mockCert2.On("GetSignedBy").Return(appmodels.NewSerialNumber(big.NewInt(987654321)))
	mockCert2.On("GetOrganizationName").Return("Example Org")

	certificates := []appmodels.ICertificate{mockCert1, mockCert2}

	// Call function
	result := apputils.ToCertificateListDTO(certificates)

	// Assert result is as expected
	// Note: You might need to adjust these assertions based on the exact output of ToCertificateDTO and NewCertificateListDTO.
	// These are placeholders to illustrate the concept.
	assert.NotNil(t, result, "Resulting CertificateListDTO should not be nil")
	assert.Len(t, result.Payload, 2, "There should be two certificates in the result")
	assert.Equal(t, "www.example1.com", result.Payload[0].CommonName, "The first certificate common name should match")
	assert.Equal(t, "www.example2.com", result.Payload[1].CommonName, "The second certificate common name should match")

	// Ensure mock expectations are met
	mockCert1.AssertExpectations(t)
	mockCert2.AssertExpectations(t)
}

func TestNewServerCertificate_ValidateDNSNamesError(t *testing.T) {
	// Mock dependencies
	mockManager := new(commonmocks.MockCertificateManager)
	mockSerialNumber := new(appmocks.MockSerialNumber)
	mockOrganization := new(appmocks.MockOrganization)
	mockPublicKey := new(appmocks.MockPublicKey)
	mockParentCertificate := new(appmocks.MockCertificate)
	mockParentPrivateKey := new(appmocks.MockPrivateKey)

	// Setup mock returns for the required inputs
	// mockSerialNumber.On("Value").Return(big.NewInt(12345))
	// mockOrganization.On("GetNames").Return([]string{"Test Organization"})
	// mockParentCertificate.On("GetCertificate").Return(&x509.Certificate{})
	// mockPublicKey.On("GetPublicKey").Return(&rsa.PublicKey{})
	// mockParentPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})
	// mockParentCertificate.On("GetParents").Return([]appmodels.ISerialNumber{})
	// mockParentCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(1)))

	// Invalid DNS names to trigger the ValidateDNSNames error
	invalidDNSNames := []string{"!invalid_dns_name"}

	// Attempt to create a new server certificate with invalid DNS names
	_, err := apputils.NewServerCertificate(
		mockManager,
		mockSerialNumber,
		mockOrganization,
		365*24*time.Hour,
		mockPublicKey,
		mockParentCertificate,
		mockParentPrivateKey,
		"www.example.com",
		invalidDNSNames...,
	)

	// Check if the error is not nil and contains the expected message
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "dnsNames: contains invalid characters")

	// Ensure mock expectations are met
	mockManager.AssertExpectations(t)
	mockSerialNumber.AssertExpectations(t)
	mockOrganization.AssertExpectations(t)
	mockPublicKey.AssertExpectations(t)
	mockParentCertificate.AssertExpectations(t)
	mockParentPrivateKey.AssertExpectations(t)
}
