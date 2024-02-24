// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"math/big"
	"testing"
)

// helper function to create a mock x509.Certificate
func newMockX509Certificate(isCA bool, organization []string, serialNumber *big.Int) *x509.Certificate {
	if serialNumber == nil {
		serialNumber = big.NewInt(1)
	}

	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	return &x509.Certificate{
		SerialNumber: serialNumber,
		IsCA:         isCA,
		Subject: pkix.Name{
			Organization: organization,
		},
		PublicKey: key.Public(),
	}
}

func TestNewCertificate(t *testing.T) {
	orgID := "Org123"
	signedBy := models.NewSerialNumber(big.NewInt(2))
	certData := newMockX509Certificate(true, []string{"Test Org"}, nil)

	cert := models.NewCertificate(orgID, signedBy, certData)
	if cert.GetOrganizationID() != orgID {
		t.Errorf("GetOrganizationID() = %v, want %v", cert.GetOrganizationID(), orgID)
	}
	certSignedBy := cert.GetSignedBy()

	if certSignedBy.Cmp(signedBy) != 0 {
		t.Errorf("GetSignedBy() = %v, want %v", cert.GetSignedBy(), signedBy)
	}
	if !cert.IsCA() {
		t.Errorf("IsCA() = false, want true")
	}
	if cert.GetOrganizationName() != "Test Org" {
		t.Errorf("GetOrganizationName() = %v, want %v", cert.GetOrganizationName(), "Test Org")
	}
}

func TestCertificate_IsCA(t *testing.T) {
	caCert := models.NewCertificate("OrgCA", models.NewSerialNumber(big.NewInt(1)), newMockX509Certificate(true, []string{"CA Org"}, nil))
	if !caCert.IsCA() {
		t.Error("IsCA() should return true for CA certificates")
	}

	nonCaCert := models.NewCertificate("OrgNonCA", models.NewSerialNumber(big.NewInt(2)), newMockX509Certificate(false, []string{"Non-CA Org"}, nil))
	if nonCaCert.IsCA() {
		t.Error("IsCA() should return false for non-CA certificates")
	}
}

func TestCertificate_GetSerialNumber(t *testing.T) {
	expectedSerial := big.NewInt(123)
	cert := models.NewCertificate(
		"Org123",
		models.NewSerialNumber(big.NewInt(2)),
		newMockX509Certificate(false, []string{"Test Org"}, expectedSerial),
	)
	bigIntCertSignedBy := cert.GetSerialNumber()

	if bigIntCertSignedBy.Value().Cmp(expectedSerial) != 0 {
		t.Errorf("GetSerialNumber() = %v, want %v", cert.GetSerialNumber().Value(), expectedSerial)
	}
}

func TestCertificate_GetOrganizationID(t *testing.T) {
	expectedOrgID := "Org123"
	cert := models.NewCertificate(expectedOrgID, models.NewSerialNumber(big.NewInt(1)), newMockX509Certificate(false, []string{"Test Org"}, nil))
	if cert.GetOrganizationID() != expectedOrgID {
		t.Errorf("GetOrganizationID() = %v, want %v", cert.GetOrganizationID(), expectedOrgID)
	}
}

func TestCertificate_GetOrganizationName(t *testing.T) {
	expectedOrgName := "PrimaryOrg"
	cert := models.NewCertificate("Org123", models.NewSerialNumber(big.NewInt(1)), newMockX509Certificate(false, []string{expectedOrgName}, nil))
	if cert.GetOrganizationName() != expectedOrgName {
		t.Errorf("GetOrganizationName() = %v, want %v", cert.GetOrganizationName(), expectedOrgName)
	}

	// Test with no organization names
	emptyStringCert := models.NewCertificate("EmptyOrg", models.NewSerialNumber(big.NewInt(1)), newMockX509Certificate(false, []string{""}, nil))
	if emptyStringCert.GetOrganizationName() != "" {
		t.Errorf("GetOrganizationName() should return an empty string when no organization names are present, got %v", emptyStringCert.GetOrganizationName())
	}

	// Test with empty organization array
	emptyCert := models.NewCertificate("EmptyOrg", models.NewSerialNumber(big.NewInt(1)), newMockX509Certificate(false, []string{}, nil))
	if emptyCert.GetOrganizationName() != "" {
		t.Errorf("GetOrganizationName() should return an empty string when no organization names are present, got %v", emptyCert.GetOrganizationName())
	}
}

func TestCertificate_GetOrganization(t *testing.T) {
	expectedOrgs := []string{"Org1", "Org2"}
	certData := newMockX509Certificate(false, []string{""}, nil)
	certData.Subject.Organization = expectedOrgs

	cert := models.NewCertificate("Org123", models.NewSerialNumber(big.NewInt(2)), certData)

	orgs := cert.GetOrganization()
	if len(orgs) != len(expectedOrgs) {
		t.Fatalf("GetOrganization returned %d organizations; want %d", len(orgs), len(expectedOrgs))
	}

	for i, org := range orgs {
		if org != expectedOrgs[i] {
			t.Errorf("Organization[%d] = %v, want %v", i, org, expectedOrgs[i])
		}
	}
}

func TestCertificate_GetSignedBy(t *testing.T) {
	expectedSignedBy := models.NewSerialNumber(big.NewInt(999))
	cert := models.NewCertificate("Org123", expectedSignedBy, newMockX509Certificate(false, []string{"Test Org"}, nil))

	bigIntSSerialNumber := cert.GetSignedBy()

	if bigIntSSerialNumber.Cmp(expectedSignedBy) != 0 {
		t.Errorf("GetSignedBy() = %v, want %v", cert.GetSignedBy(), expectedSignedBy)
	}
}

func TestCertificate_GetCertificate(t *testing.T) {
	expectedCert := newMockX509Certificate(true, []string{"Acme Co"}, nil)
	cert := models.NewCertificate("Acme123", models.NewSerialNumber(big.NewInt(2)), expectedCert)

	if cert.GetCertificate() != expectedCert {
		t.Error("GetCertificate did not return the expected *x509.Certificate")
	}
}
