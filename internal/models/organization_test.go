// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"crypto/x509"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"io"
	"testing"
	"time"
)

func TestNewOrganization(t *testing.T) {
	orgID := "org123"
	names := []string{"Test Org", "Test Org Department"}
	org := models.NewOrganization(orgID, names)

	if org.GetID() != orgID {
		t.Errorf("GetID() = %s, want %s", org.GetID(), orgID)
	}

	if len(org.GetNames()) != len(names) {
		t.Fatalf("GetNames() returned %d names; want %d", len(org.GetNames()), len(names))
	}

	for i, name := range org.GetNames() {
		if name != names[i] {
			t.Errorf("GetNames()[%d] = %s, want %s", i, name, names[i])
		}
	}
}

func TestOrganization_GetID(t *testing.T) {
	orgID := "org456"
	org := models.NewOrganization(orgID, nil)

	if got := org.GetID(); got != orgID {
		t.Errorf("GetID() = %s, want = %s", got, orgID)
	}
}

func TestOrganization_GetName(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := models.NewOrganization("org789", names)

	if got := org.GetName(); got != names[0] {
		t.Errorf("GetName() = %s, want = %s", got, names[0])
	}
}

func TestOrganization_GetName_NoNames(t *testing.T) {
	org := models.NewOrganization("orgNoNames", []string{})
	if name := org.GetName(); name != "" {
		t.Errorf("GetName() with no names should return an empty string, got: %s", name)
	}
}

func TestOrganization_GetNames(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := models.NewOrganization("org101112", names)

	gotNames := org.GetNames()
	if len(gotNames) != len(names) || gotNames[0] != names[0] || gotNames[1] != names[1] {
		t.Errorf("GetNames() got = %v, want = %v", gotNames, names)
	}
}

func TestOrganization_NewRootCertificate(t *testing.T) {
	org := models.NewOrganization("orgID", []string{"Test Org"})
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	privateKey, err := models.GenerateRSAPrivateKey(serialNumber, 2048)
	if err != nil {
		t.Fatalf("GenerateRSAPrivateKey failed: %v", err)
	}

	cert, err := org.NewRootCertificate(managers.NewCertificateManager(randomManager), "Test Root CA", *privateKey, 365*24*time.Hour)
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

func TestOrganization_NewRootCertificate_CreateCertificateError(t *testing.T) {
	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return nil, fmt.Errorf("simulated create certificate error")
		},
	}

	org := models.NewOrganization("orgID", []string{"Test Org"})
	serialNumber, _ := models.NewSerialNumber(&mocks.MockRandomManager{})
	privateKey, _ := models.GenerateRSAPrivateKey(serialNumber, 2048)

	_, err := org.NewRootCertificate(mockManager, "Test Root CA", *privateKey, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewRootCertificate, got nil")
	}
}

func TestOrganization_NewIntermediateCertificate(t *testing.T) {
	org := models.NewOrganization("intermediateOrgID", []string{"Intermediate Test Org"})

	certManager := managers.NewCertificateManager(managers.NewRandomManager())

	randomManager := managers.NewRandomManager()
	parentSerialNumber, _ := models.NewSerialNumber(randomManager)
	parentPrivateKey, _ := models.GenerateRSAPrivateKey(parentSerialNumber, 2048)
	parentCert, _ := org.NewRootCertificate(certManager, "Parent Root CA", *parentPrivateKey, 365*24*time.Hour)

	intermediateSerialNumber, _ := models.NewSerialNumber(randomManager)
	intermediateCert, err := org.NewIntermediateCertificate(certManager, "Intermediate CA", intermediateSerialNumber, parentCert, *parentPrivateKey, 365*24*time.Hour)
	if err != nil {
		t.Fatalf("NewIntermediateCertificate failed: %v", err)
	}

	// Complete assertions for TestOrganization_NewIntermediateCertificate
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

func TestOrganization_NewIntermediateCertificate_Error(t *testing.T) {
	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			// Simulate an error during certificate creation
			return nil, fmt.Errorf("error creating intermediate certificate")
		},
	}

	org := models.NewOrganization("orgID", []string{"Test Org"})
	serialNumber, _ := models.NewSerialNumber(&mocks.MockRandomManager{})
	parentCertificate := &models.Certificate{} // Mock or prepare a parent certificate as needed
	privateKey := &models.PrivateKey{}         // Mock or prepare a private key as needed

	_, err := org.NewIntermediateCertificate(mockManager, "Intermediate CA", serialNumber, parentCertificate, *privateKey, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewIntermediateCertificate, got nil")
	}
}

func TestOrganization_NewServerCertificate(t *testing.T) {
	org := models.NewOrganization("serverOrgID", []string{"Server Test Org"})
	certManager := managers.NewCertificateManager(managers.NewRandomManager())

	randomManager := managers.NewRandomManager()

	parentSerialNumber, _ := models.NewSerialNumber(randomManager)
	parentPrivateKey, _ := models.GenerateRSAPrivateKey(parentSerialNumber, 2048)
	parentCert, _ := org.NewRootCertificate(certManager, "Parent CA", *parentPrivateKey, 365*24*time.Hour)

	serverSerialNumber, _ := models.NewSerialNumber(randomManager)
	dnsNames := []string{"www.example.com", "example.com"}
	serverCert, err := org.NewServerCertificate(certManager, serverSerialNumber, parentCert, *parentPrivateKey, dnsNames, 365*24*time.Hour)
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

func TestOrganization_NewServerCertificate_Error(t *testing.T) {
	mockManager := mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return nil, fmt.Errorf("error creating server certificate")
		},
	}

	org := models.NewOrganization("serverOrgID", []string{"Server Test Org"})
	serialNumber, _ := models.NewSerialNumber(&mocks.MockRandomManager{})
	parentCertificate := &models.Certificate{} // Prepare a mock parent certificate
	privateKey := &models.PrivateKey{}         // Prepare a mock private key
	dnsNames := []string{"www.example.org"}

	_, err := org.NewServerCertificate(&mockManager, serialNumber, parentCertificate, *privateKey, dnsNames, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewServerCertificate, got nil")
	}
}

func TestOrganization_NewClientCertificate(t *testing.T) {
	certManager := managers.NewCertificateManager(managers.NewRandomManager())

	randomManager := managers.NewRandomManager()

	org := models.NewOrganization("clientOrgID", []string{"Client Test Org"})
	parentSerialNumber, _ := models.NewSerialNumber(randomManager)
	parentPrivateKey, _ := models.GenerateRSAPrivateKey(parentSerialNumber, 2048)
	parentCert, _ := org.NewRootCertificate(certManager, "Parent CA", *parentPrivateKey, 365*24*time.Hour)

	clientSerialNumber, _ := models.NewSerialNumber(randomManager)
	clientCert, err := org.NewClientCertificate(certManager, "Client", clientSerialNumber, parentCert, *parentPrivateKey, 365*24*time.Hour)
	if err != nil {
		t.Fatalf("NewClientCertificate failed: %v", err)
	}

	if clientCert.GetCertificate().ExtKeyUsage[0] != x509.ExtKeyUsageClientAuth {
		t.Error("Expected ExtKeyUsage to include ClientAuth")
	}
}

func TestOrganization_NewClientCertificate_Error(t *testing.T) {
	mockManager := mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return nil, fmt.Errorf("error creating client certificate")
		},
	}

	org := models.NewOrganization("clientOrgID", []string{"Client Test Org"})
	serialNumber, _ := models.NewSerialNumber(&mocks.MockRandomManager{})
	parentCertificate := &models.Certificate{} // Prepare a mock parent certificate
	privateKey := &models.PrivateKey{}         // Prepare a mock private key

	_, err := org.NewClientCertificate(&mockManager, "Client", serialNumber, parentCertificate, *privateKey, 365*24*time.Hour)
	if err == nil {
		t.Fatal("Expected an error from NewClientCertificate, got nil")
	}
}
