// Copyright (c) 2024. Heusala roup Oy <info@heusalagroup.fi>. All rights reserved.

package modelcontrollers_test

import (
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/modelutils"
	"io"
	"testing"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

func TestNewOrganizationController(t *testing.T) {

	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	if !controller.UsesOrganizationService(mockService) {
		t.Fatalf("Expected the organization controller to use the mockService, got false")
	}

}

// Now, let's write a test for a successful operation in OrganizationController.
func TestOrganizationController_CreateOrganization_Success(t *testing.T) {
	// Setup
	mockService := &mocks.MockOrganizationService{
		CreateOrganizationFunc: func(certificate models.IOrganization) (models.IOrganization, error) {
			return certificate, nil // Simulate successful creation
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	// Execute
	newOrg := &models.Organization{} // You'd fill this with actual data
	createdOrg, err := controller.CreateOrganization(newOrg)

	// Verify
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if createdOrg != newOrg {
		t.Errorf("Expected created organization to be the same as the input")
	}
}

// Test for handling errors from the service
func TestOrganizationController_CreateOrganization_Error(t *testing.T) {
	// Setup
	expectedError := errors.New("failed to create organization")
	mockService := &mocks.MockOrganizationService{
		CreateOrganizationFunc: func(certificate models.IOrganization) (models.IOrganization, error) {
			return nil, expectedError // Simulate failure
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	// Execute
	newOrg := &models.Organization{} // Assuming this would be populated with test data
	_, err := controller.CreateOrganization(newOrg)

	// Verify
	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}
}

func TestOrganizationController_GetExistingOrganization(t *testing.T) {

	orgId := "testorg"
	expectedModel := &models.Organization{ /* Initialized fields */ }

	mockService := &mocks.MockOrganizationService{
		GetExistingOrganizationFunc: func(id string) (models.IOrganization, error) {
			return expectedModel, nil
		},
	}

	controller := modelcontrollers.NewOrganizationController(mockService)

	organization, err := controller.GetExistingOrganization(orgId)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
	if organization != expectedModel {
		t.Errorf("Expected to receive %v, got %v", expectedModel, organization)
	}
}

func TestOrganizationController_NewRootCertificate(t *testing.T) {

	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)
	organization := "orgID"

	org := models.NewOrganization(organization, []string{"Test Org"})
	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	privateKey, err := modelutils.GenerateRSAPrivateKey(organization, []models.ISerialNumber{serialNumber}, 2048)
	if err != nil {
		t.Fatalf("GenerateRSAPrivateKey failed: %v", err)
	}

	cert, err := controller.NewRootCertificate(org, managers.NewCertificateManager(randomManager), "Test Root CA", privateKey, 365*24*time.Hour)
	if err != nil {
		t.Fatalf("NewRootCertificate failed: %v", err)
	}

	if cert.GetOrganizationID() != "orgID" {
		t.Errorf("Certificate organization ID got = %s, want %s", cert.GetOrganizationID(), "orgID")
	}

	if len(cert.GetCertificate().Subject.Organization) == 0 || cert.GetCertificate().Subject.Organization[0] != "Test Org" {
		t.Errorf("Certificate organization name got = %v, want %v", cert.GetCertificate().Subject.Organization, []string{"Test Org"})
	}

	if !cert.GetCertificate().IsCA {
		t.Errorf("Expected IsCA = true, got false")
	}
}

func TestOrganizationController_NewRootCertificate_CreateCertificateError(t *testing.T) {
	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)
	organization := "orgID"

	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return nil, fmt.Errorf("simulated create certificate error")
		},
	}

	org := models.NewOrganization(organization, []string{"Test Org"})
	serialNumber, _ := modelutils.GenerateSerialNumber(&mocks.MockRandomManager{})
	privateKey, _ := modelutils.GenerateRSAPrivateKey(organization, []models.ISerialNumber{serialNumber}, 2048)

	_, err := controller.NewRootCertificate(org, mockManager, "Test Root CA", privateKey, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewRootCertificate, got nil")
	}
}

func TestOrganizationController_NewIntermediateCertificate(t *testing.T) {

	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	organization := "intermediateOrgID"
	org := models.NewOrganization("intermediateOrgID", []string{"Intermediate Test Org"})

	certManager := managers.NewCertificateManager(managers.NewRandomManager())

	randomManager := managers.NewRandomManager()
	parentSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	parentPrivateKey, _ := modelutils.GenerateRSAPrivateKey(organization, []models.ISerialNumber{parentSerialNumber}, 2048)
	parentCert, _ := controller.NewRootCertificate(org, certManager, "Parent Root CA", parentPrivateKey, 365*24*time.Hour)

	intermediateSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	intermediateCert, err := controller.NewIntermediateCertificate(org, certManager, "Intermediate CA", intermediateSerialNumber, parentCert, parentPrivateKey, 365*24*time.Hour)
	if err != nil {
		t.Fatalf("NewIntermediateCertificate failed: %v", err)
	}

	// Complete assertions for TestOrganizationController_NewIntermediateCertificate
	if !intermediateCert.GetCertificate().IsCA {
		t.Error("Expected intermediate certificate to be a CA")
	}
	if intermediateCert.GetCertificate().MaxPathLenZero != true {
		t.Error("Expected MaxPathLenZero = true for intermediate certificate")
	}
	if intermediateCert.GetCertificate().Subject.CommonName != "Intermediate CA" {
		t.Errorf("Expected CommonName = 'Intermediate CA', got %s", intermediateCert.GetCertificate().Subject.CommonName)
	}

}

func TestOrganizationController_NewIntermediateCertificate_Error(t *testing.T) {

	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			// Simulate an error during certificate creation
			return nil, fmt.Errorf("error creating intermediate certificate")
		},
	}

	org := models.NewOrganization("orgID", []string{"Test Org"})
	serialNumber, _ := modelutils.GenerateSerialNumber(&mocks.MockRandomManager{})
	parentCertificate := &models.Certificate{} // Mock or prepare a parent certificate as needed
	privateKey := &models.PrivateKey{}         // Mock or prepare a private key as needed

	_, err := controller.NewIntermediateCertificate(org, mockManager, "Intermediate CA", serialNumber, parentCertificate, privateKey, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewIntermediateCertificate, got nil")
	}
}

func TestOrganizationController_NewServerCertificate(t *testing.T) {
	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	organization := "serverOrgID"
	org := models.NewOrganization(organization, []string{"Server Test Org"})
	certManager := managers.NewCertificateManager(managers.NewRandomManager())

	randomManager := managers.NewRandomManager()

	parentSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	parentPrivateKey, _ := modelutils.GenerateRSAPrivateKey(organization, []models.ISerialNumber{parentSerialNumber}, 2048)
	parentCert, _ := controller.NewRootCertificate(org, certManager, "Parent CA", parentPrivateKey, 365*24*time.Hour)

	serverSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	dnsNames := []string{"www.example.com", "example.com"}
	serverCert, err := controller.NewServerCertificate(org, certManager, serverSerialNumber, parentCert, parentPrivateKey, dnsNames, 365*24*time.Hour)
	if err != nil {
		t.Fatalf("NewServerCertificate failed: %v", err)
	}

	if serverCert.GetCertificate().DNSNames[0] != dnsNames[0] || serverCert.GetCertificate().DNSNames[1] != dnsNames[1] {
		t.Errorf("DNSNames got = %v, want %v", serverCert.GetCertificate().DNSNames, dnsNames)
	}
	if serverCert.GetCertificate().ExtKeyUsage[0] != x509.ExtKeyUsageServerAuth {
		t.Error("Expected ExtKeyUsage to include ServerAuth")
	}
}

func TestOrganizationController_NewServerCertificate_Error(t *testing.T) {
	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	mockManager := mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return nil, fmt.Errorf("error creating server certificate")
		},
	}

	org := models.NewOrganization("serverOrgID", []string{"Server Test Org"})
	serialNumber, _ := modelutils.GenerateSerialNumber(&mocks.MockRandomManager{})
	parentCertificate := &models.Certificate{} // Prepare a mock parent certificate
	privateKey := &models.PrivateKey{}         // Prepare a mock private key
	dnsNames := []string{"www.example.org"}

	_, err := controller.NewServerCertificate(org, &mockManager, serialNumber, parentCertificate, privateKey, dnsNames, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewServerCertificate, got nil")
	}
}

func TestOrganizationController_NewClientCertificate(t *testing.T) {

	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	certManager := managers.NewCertificateManager(managers.NewRandomManager())

	randomManager := managers.NewRandomManager()

	organization := "clientOrgID"

	org := models.NewOrganization(organization, []string{"Client Test Org"})
	parentSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	parentPrivateKey, _ := modelutils.GenerateRSAPrivateKey(organization, []models.ISerialNumber{parentSerialNumber}, 2048)
	parentCert, _ := controller.NewRootCertificate(org, certManager, "Parent CA", parentPrivateKey, 365*24*time.Hour)

	clientSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	clientCert, err := controller.NewClientCertificate(org, certManager, "Client", clientSerialNumber, parentCert, parentPrivateKey, 365*24*time.Hour)
	if err != nil {
		t.Fatalf("NewClientCertificate failed: %v", err)
	}

	if clientCert.GetCertificate().ExtKeyUsage[0] != x509.ExtKeyUsageClientAuth {
		t.Error("Expected ExtKeyUsage to include ClientAuth")
	}
}

func TestOrganizationController_NewClientCertificate_Error(t *testing.T) {
	mockService := &mocks.MockOrganizationService{}
	controller := modelcontrollers.NewOrganizationController(mockService)

	mockManager := mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return nil, fmt.Errorf("error creating client certificate")
		},
	}

	org := models.NewOrganization("clientOrgID", []string{"Client Test Org"})
	serialNumber, _ := modelutils.GenerateSerialNumber(&mocks.MockRandomManager{})
	parentCertificate := &models.Certificate{} // Prepare a mock parent certificate
	privateKey := &models.PrivateKey{}         // Prepare a mock private key

	_, err := controller.NewClientCertificate(org, &mockManager, "Client", serialNumber, parentCertificate, privateKey, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewClientCertificate, got nil")
	}
}
