// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apputils_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

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
	dto := apputils.GetCertificateDTO(cert)

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
	publicKey, privateKey := "publicKey", "privateKey" // Simplified for example

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
	parentPrivateKey := &appmocks.MockPrivateKey{}
	commonName := "Server Certificate"
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
