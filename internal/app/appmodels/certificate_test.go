// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
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
	signedBy := big.NewInt(2)
	certData := newMockX509Certificate(true, []string{"Test Org"}, nil)

	cert := appmodels.NewCertificate(orgID, []*big.Int{signedBy}, certData)
	if cert.OrganizationID() != orgID {
		t.Errorf("OrganizationID() = %v, want %v", cert.OrganizationID(), orgID)
	}
	certSignedBy := cert.SignedBy()

	if certSignedBy.Cmp(signedBy) != 0 {
		t.Errorf("SignedBy() = %v, want %v", cert.SignedBy(), signedBy)
	}
	if !cert.IsCA() {
		t.Errorf("IsCA() = false, want true")
	}
	if cert.OrganizationName() != "Test Org" {
		t.Errorf("OrganizationName() = %v, want %v", cert.OrganizationName(), "Test Org")
	}
}

func TestCertificate_IsCA(t *testing.T) {
	caCert := appmodels.NewCertificate("OrgCA", []*big.Int{big.NewInt(1)}, newMockX509Certificate(true, []string{"CA Org"}, nil))
	if !caCert.IsCA() {
		t.Error("IsCA() should return true for CA certificates")
	}

	nonCaCert := appmodels.NewCertificate("OrgNonCA", []*big.Int{big.NewInt(2)}, newMockX509Certificate(false, []string{"Non-CA Org"}, nil))
	if nonCaCert.IsCA() {
		t.Error("IsCA() should return false for non-CA certificates")
	}
}

func TestCertificate_GetSerialNumber(t *testing.T) {
	expectedSerial := big.NewInt(123)
	cert := appmodels.NewCertificate(
		"Org123",
		[]*big.Int{big.NewInt(2)},
		newMockX509Certificate(false, []string{"Test Org"}, expectedSerial),
	)
	bigIntCertSignedBy := cert.SerialNumber()

	if bigIntCertSignedBy.Cmp(expectedSerial) != 0 {
		t.Errorf("SerialNumber() = %v, want %v", cert.SerialNumber(), expectedSerial)
	}
}

func TestCertificate_GetOrganizationID(t *testing.T) {
	expectedOrgID := "Org123"
	cert := appmodels.NewCertificate(expectedOrgID, []*big.Int{big.NewInt(1)}, newMockX509Certificate(false, []string{"Test Org"}, nil))
	if cert.OrganizationID() != expectedOrgID {
		t.Errorf("OrganizationID() = %v, want %v", cert.OrganizationID(), expectedOrgID)
	}
}

func TestCertificate_GetOrganizationName(t *testing.T) {
	expectedOrgName := "PrimaryOrg"
	cert := appmodels.NewCertificate("Org123", []*big.Int{big.NewInt(1)}, newMockX509Certificate(false, []string{expectedOrgName}, nil))
	if cert.OrganizationName() != expectedOrgName {
		t.Errorf("OrganizationName() = %v, want %v", cert.OrganizationName(), expectedOrgName)
	}

	// Test with no organization names
	emptyStringCert := appmodels.NewCertificate("EmptyOrg", []*big.Int{big.NewInt(1)}, newMockX509Certificate(false, []string{""}, nil))
	if emptyStringCert.OrganizationName() != "" {
		t.Errorf("OrganizationName() should return an empty string when no organization names are present, got %v", emptyStringCert.OrganizationName())
	}

	// Test with empty organization array
	emptyCert := appmodels.NewCertificate("EmptyOrg", []*big.Int{big.NewInt(1)}, newMockX509Certificate(false, []string{}, nil))
	if emptyCert.OrganizationName() != "" {
		t.Errorf("OrganizationName() should return an empty string when no organization names are present, got %v", emptyCert.OrganizationName())
	}
}

func TestCertificate_GetOrganization(t *testing.T) {
	expectedOrgs := []string{"Org1", "Org2"}
	certData := newMockX509Certificate(false, []string{""}, nil)
	certData.Subject.Organization = expectedOrgs

	cert := appmodels.NewCertificate("Org123", []*big.Int{big.NewInt(2)}, certData)

	orgs := cert.Organization()
	if len(orgs) != len(expectedOrgs) {
		t.Fatalf("Organization returned %d organizations; want %d", len(orgs), len(expectedOrgs))
	}

	for i, org := range orgs {
		if org != expectedOrgs[i] {
			t.Errorf("Organization[%d] = %v, want %v", i, org, expectedOrgs[i])
		}
	}
}

func TestCertificate_GetSignedBy(t *testing.T) {
	expectedSignedBy := big.NewInt(999)
	cert := appmodels.NewCertificate("Org123", []*big.Int{expectedSignedBy}, newMockX509Certificate(false, []string{"Test Org"}, nil))

	bigIntSSerialNumber := cert.SignedBy()

	if bigIntSSerialNumber.Cmp(expectedSignedBy) != 0 {
		t.Errorf("SignedBy() = %v, want %v", cert.SignedBy(), expectedSignedBy)
	}
}

func TestCertificate_GetCertificate(t *testing.T) {
	expectedCert := newMockX509Certificate(true, []string{"Acme Co"}, nil)
	cert := appmodels.NewCertificate("Acme123", []*big.Int{big.NewInt(2)}, expectedCert)

	if cert.Certificate() != expectedCert {
		t.Error("Certificate did not return the expected *x509.Certificate")
	}
}

// TestCertificate_IsSelfSigned tests both self-signed and not self-signed scenarios for a certificate.
func TestCertificate_IsSelfSigned(t *testing.T) {
	// Generate a public/private key pair for signing
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Template for the self-signed certificate
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "Self Signed",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	// Create a self-signed certificate
	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &privKey.PublicKey, privKey)
	if err != nil {
		t.Fatalf("Failed to create certificate: %v", err)
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		t.Fatalf("Failed to parse certificate: %v", err)
	}

	// Simulate the scenario where AuthorityKeyId and SubjectKeyId match
	cert.AuthorityKeyId = []byte{1, 2, 3}
	cert.SubjectKeyId = []byte{1, 2, 3}

	modelCert := appmodels.NewCertificate("SelfSignedOrg", []*big.Int{big.NewInt(1)}, cert)
	if !modelCert.IsSelfSigned() {
		t.Errorf("Expected IsSelfSigned to be true for a self-signed certificate")
	}

	// Simulate the scenario where AuthorityKeyId and SubjectKeyId do not match
	cert2 := *cert                       // Make a copy to modify
	cert2.SubjectKeyId = []byte{4, 5, 6} // Change SubjectKeyId to not match AuthorityKeyId

	modelCert2 := appmodels.NewCertificate("NonSelfSignedOrg", []*big.Int{big.NewInt(2)}, &cert2)
	if modelCert2.IsSelfSigned() {
		t.Errorf("Expected IsSelfSigned to be false when AuthorityKeyId and SubjectKeyId do not match")
	}

	// Simulate the scenario where AuthorityKeyId and SubjectKeyId are not set, but Issuer and Subject match
	cert3 := *cert // Make another copy
	cert3.AuthorityKeyId = nil
	cert3.SubjectKeyId = nil

	modelCert3 := appmodels.NewCertificate("SelfSignedByNamesOrg", []*big.Int{big.NewInt(3)}, &cert3)
	if !modelCert3.IsSelfSigned() {
		t.Errorf("Expected IsSelfSigned to be true when AuthorityKeyId and SubjectKeyId are not set but Issuer and Subject match")
	}
}

func TestCertificate_IsIntermediateCertificate(t *testing.T) {
	// Common setup for certificates
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	// Scenario 1: Certificate with BasicConstraintsValid, IsCA but is self-signed (not an intermediate)
	selfSignedTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "Self Signed CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	selfSignedCertDER, err := x509.CreateCertificate(rand.Reader, selfSignedTemplate, selfSignedTemplate, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("Failed to create self-signed certificate: %v", err)
	}
	selfSignedCert, err := x509.ParseCertificate(selfSignedCertDER)
	if err != nil {
		t.Fatalf("Failed to parse self-signed certificate: %v", err)
	}
	modelSelfSignedCert := appmodels.NewCertificate("Org1", []*big.Int{big.NewInt(1)}, selfSignedCert)
	if modelSelfSignedCert.IsIntermediateCertificate() {
		t.Errorf("Expected IsIntermediateCertificate to be false for a self-signed CA")
	}

	// Scenario 2: Certificate with BasicConstraintsValid, IsCA and not self-signed (an intermediate)
	issuerTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject: pkix.Name{
			CommonName: "Issuer CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	issuerCertDER, err := x509.CreateCertificate(rand.Reader, issuerTemplate, selfSignedTemplate, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("Failed to create issuer certificate: %v", err)
	}
	issuerCert, err := x509.ParseCertificate(issuerCertDER)
	if err != nil {
		t.Fatalf("Failed to parse issuer certificate: %v", err)
	}
	modelIssuerCert := appmodels.NewCertificate("Org2", []*big.Int{big.NewInt(2)}, issuerCert)
	if !modelIssuerCert.IsIntermediateCertificate() {
		t.Errorf("Expected IsIntermediateCertificate to be true for an intermediate CA")
	}

	// Scenario 3: Certificate without IsCA (not an intermediate, despite not being self-signed)
	nonCaTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(3),
		Subject: pkix.Name{
			CommonName: "Non CA",
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
		IsCA:                  false,
	}
	nonCaCertDER, err := x509.CreateCertificate(rand.Reader, nonCaTemplate, selfSignedTemplate, &key.PublicKey, key)
	if err != nil {
		t.Fatalf("Failed to create non-CA certificate: %v", err)
	}
	nonCaCert, err := x509.ParseCertificate(nonCaCertDER)
	if err != nil {
		t.Fatalf("Failed to parse non-CA certificate: %v", err)
	}
	modelNonCaCert := appmodels.NewCertificate("Org3", []*big.Int{big.NewInt(3)}, nonCaCert)
	if modelNonCaCert.IsIntermediateCertificate() {
		t.Errorf("Expected IsIntermediateCertificate to be false for a non-CA certificate")
	}
}

// helper function to create a mock x509.Certificate with specific extended key usages
func newMockX509CertificateWithExtUsage(extKeyUsage []x509.ExtKeyUsage, commonName string) *x509.Certificate {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour),
		ExtKeyUsage: extKeyUsage,
		PublicKey:   key.Public(),
	}

	return cert
}

func TestCertificate_IsServerCertificate(t *testing.T) {
	// Test case 1: Certificate with server authentication usage
	serverAuthCert := appmodels.NewCertificate(
		"ServerAuthOrg",
		[]*big.Int{big.NewInt(1)},
		newMockX509CertificateWithExtUsage([]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, "ServerAuth"),
	)
	if !serverAuthCert.IsServerCertificate() {
		t.Errorf("Expected IsServerCertificate to be true for a certificate with ExtKeyUsageServerAuth")
	}

	// Test case 2: Certificate without server authentication usage
	nonServerAuthCert := appmodels.NewCertificate(
		"NonServerAuthOrg",
		[]*big.Int{big.NewInt(2)},
		newMockX509CertificateWithExtUsage([]x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, "NonServerAuth"),
	)
	if nonServerAuthCert.IsServerCertificate() {
		t.Errorf("Expected IsServerCertificate to be false for a certificate without ExtKeyUsageServerAuth")
	}

	// Test case 3: Certificate with multiple extended key usages including server authentication
	multiUsageCert := appmodels.NewCertificate(
		"MultiUsageOrg",
		[]*big.Int{big.NewInt(3)},
		newMockX509CertificateWithExtUsage([]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}, "MultiUsage"),
	)
	if !multiUsageCert.IsServerCertificate() {
		t.Errorf("Expected IsServerCertificate to be true for a certificate with multiple extended key usages including ExtKeyUsageServerAuth")
	}
}

// Helper function for creating a mock x509.Certificate with specified extended key usages
func newMockX509CertificateWithClientExtUsage(extKeyUsage []x509.ExtKeyUsage, commonName string) *x509.Certificate {
	key, _ := rsa.GenerateKey(rand.Reader, 2048) // Ignoring error for brevity
	return &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().Add(365 * 24 * time.Hour), // Valid for one year
		ExtKeyUsage: extKeyUsage,
		PublicKey:   key.Public(),
	}
}

func TestCertificate_IsClientCertificate(t *testing.T) {
	// Case 1: Certificate specifically marked for client authentication
	clientAuthCert := appmodels.NewCertificate(
		"ClientAuthOrg",
		[]*big.Int{big.NewInt(1)},
		newMockX509CertificateWithClientExtUsage([]x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}, "ClientAuth"),
	)
	if !clientAuthCert.IsClientCertificate() {
		t.Errorf("Expected IsClientCertificate to be true for a certificate with ExtKeyUsageClientAuth")
	}

	// Case 2: Certificate without client authentication usage
	nonClientAuthCert := appmodels.NewCertificate(
		"NonClientAuthOrg",
		[]*big.Int{big.NewInt(2)},
		newMockX509CertificateWithClientExtUsage([]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}, "NonClientAuth"),
	)
	if nonClientAuthCert.IsClientCertificate() {
		t.Errorf("Expected IsClientCertificate to be false for a certificate without ExtKeyUsageClientAuth")
	}

	// Case 3: Certificate with multiple extended key usages including client authentication
	multiUsageCert := appmodels.NewCertificate(
		"MultiUsageOrg",
		[]*big.Int{big.NewInt(3)},
		newMockX509CertificateWithClientExtUsage([]x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}, "MultiUsageClient"),
	)
	if !multiUsageCert.IsClientCertificate() {
		t.Errorf("Expected IsClientCertificate to be true for a certificate with multiple extended key usages including ExtKeyUsageClientAuth")
	}
}

// Helper function to create a self-signed x509.Certificate
func newMockSelfSignedX509Certificate(isCA bool, commonName string, serialNumber *big.Int) *x509.Certificate {
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048) // Ignoring error for simplicity

	template := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: commonName,
		},
		Issuer: pkix.Name{
			CommonName: commonName, // Same as Subject for self-signed
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  isCA,
		PublicKey:             &privKey.PublicKey,
	}

	certDER, _ := x509.CreateCertificate(rand.Reader, template, template, &privKey.PublicKey, privKey)
	cert, _ := x509.ParseCertificate(certDER)

	return cert
}

func TestCertificate_IsRootCertificate(t *testing.T) {
	// Root CA Certificate
	rootCASerialNumber := big.NewInt(1)
	rootCACert := appmodels.NewCertificate(
		"RootCAOrg",
		[]*big.Int{rootCASerialNumber},
		newMockSelfSignedX509Certificate(true, "Root CA", rootCASerialNumber),
	)
	if !rootCACert.IsRootCertificate() {
		t.Errorf("Expected IsRootCertificate to be true for a root CA certificate")
	}

	// Non-CA Certificate (should not be considered a root certificate)
	nonCASerialNumber := big.NewInt(2)
	nonCACert := appmodels.NewCertificate(
		"NonCAOrg",
		[]*big.Int{nonCASerialNumber},
		newMockSelfSignedX509Certificate(false, "Non CA", nonCASerialNumber),
	)
	if nonCACert.IsRootCertificate() {
		t.Errorf("Expected IsRootCertificate to be false for a non-CA certificate")
	}

	// Intermediate CA Certificate (not self-signed, hence not a root certificate)
	intermediateCASerialNumber := big.NewInt(3)
	intermediateCACert := appmodels.NewCertificate(
		"IntermediateCAOrg",
		[]*big.Int{intermediateCASerialNumber},
		newMockSelfSignedX509Certificate(true, "Intermediate CA", intermediateCASerialNumber),
	)
	// Manually break the self-signed condition
	intermediateCACert.Certificate().Issuer = pkix.Name{CommonName: "Different Issuer"}
	if intermediateCACert.IsRootCertificate() {
		t.Errorf("Expected IsRootCertificate to be false for an intermediate CA certificate")
	}
}

// Helper function to create a mock x509.Certificate with a specific CommonName
func newMockX509CertificateWithCommonName(commonName string) *x509.Certificate {
	return &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: commonName,
		},
	}
}

func TestCertificate_GetCommonName(t *testing.T) {
	tests := []struct {
		name       string
		commonName string
	}{
		{name: "Test CommonName 1", commonName: "example.com"},
		{name: "Test CommonName 2", commonName: "www.example.com"},
		{name: "Test CommonName 3", commonName: "mail.example.com"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			certData := newMockX509CertificateWithCommonName(test.commonName)
			cert := appmodels.NewCertificate("Org123", []*big.Int{big.NewInt(1)}, certData)

			if got := cert.CommonName(); got != test.commonName {
				t.Errorf("CommonName() = %v, want %v", got, test.commonName)
			}
		})
	}
}

func TestCertificate_GetParents(t *testing.T) {
	// Setup
	parentSerialNumbers := []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
	}
	cert := appmodels.NewCertificate("Org123", parentSerialNumbers, newMockX509Certificate(false, []string{"Test Org"}, big.NewInt(123)))

	// Act
	returnedParents := cert.Parents()

	// Assert that returnedParents is a correct copy of parentSerialNumbers
	if len(returnedParents) != len(parentSerialNumbers) {
		t.Fatalf("Parents returned %d parents; want %d", len(returnedParents), len(parentSerialNumbers))
	}

	for i, serialNumber := range returnedParents {
		if serialNumber.String() != parentSerialNumbers[i].String() {
			t.Errorf("Parent serial number at index %d = %v, want %v", i, serialNumber.String(), parentSerialNumbers[i].String())
		}
	}

	// Modify returned slice to test if original slice is unaffected
	if len(returnedParents) > 0 {
		returnedParents[0] = big.NewInt(999) // Change first element
		if cert.Parents()[0].String() == "999" {
			t.Error("Modifying the returned slice from Parents should not affect the original parents slice")
		}
	}
}

func TestCertificate_GetParents_Incomplete(t *testing.T) {

	// Setup
	cert := &appmodels.CertificateModel{}

	// Act
	returnedParents := cert.Parents()

	// Assert that returnedParents is a correct copy of parentSerialNumbers
	if len(returnedParents) != 0 {
		t.Fatalf("Parents returned %d parents; want 0", len(returnedParents))
	}

}

func TestCertificate_GetSignedBy_WithParents(t *testing.T) {

	// Setup a certificate with parents
	parentSerialNumbers := []*big.Int{
		big.NewInt(1),
		big.NewInt(2), // This should be returned by SignedBy
	}
	cert := appmodels.NewCertificate("Org123", parentSerialNumbers, newMockX509Certificate(false, []string{"Test Org"}, big.NewInt(3)))

	// Act and assert
	signedBy := cert.SignedBy()
	if signedBy.String() != "2" {
		t.Errorf("SignedBy() returned %v; want %v", signedBy.String(), "2")
	}
}

func TestCertificate_GetSignedBy_WithoutParents(t *testing.T) {

	// Setup a certificate without parents
	ownSerialNumber := big.NewInt(3)

	cert := appmodels.NewCertificate(
		"Org123",
		[]*big.Int{},
		newMockX509Certificate(
			false,
			[]string{"Test Org"},
			ownSerialNumber,
		),
	)

	// Act and assert
	signedBy := cert.SignedBy()
	if signedBy.String() != ownSerialNumber.String() {
		t.Errorf("SignedBy() returned %v; want %v", signedBy.String(), ownSerialNumber.String())
	}
}

func newMockCertificateWithValidity(notBefore, notAfter time.Time) *x509.Certificate {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err) // For simplicity in a test setup
	}

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "www.example.com",
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,
		PublicKey: key.Public(),
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	if err != nil {
		panic(err) // For simplicity in a test setup
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		panic(err) // For simplicity in a test setup
	}

	return cert
}

func TestCertificate_GetNotBeforeAndAfter(t *testing.T) {
	// Set up specific notBefore and notAfter times for the test
	notBefore := time.Now().Add(-time.Hour).UTC().Truncate(time.Second)
	notAfter := time.Now().Add(time.Hour).UTC().Truncate(time.Second)

	// Create a mock certificate with the specified validity
	mockCert := newMockCertificateWithValidity(notBefore, notAfter)

	// Wrap the mock certificate in your Certificate model
	certModel := appmodels.NewCertificate("org123", nil, mockCert)

	assert.Equal(t, notBefore.String(), certModel.NotBefore().String())
	assert.Equal(t, notAfter.String(), certModel.NotAfter().String())

	// Test NotBefore method
	assert.True(t, certModel.NotBefore().Equal(mockCert.NotBefore), "NotBefore should return the correct notBefore time")
	assert.True(t, certModel.NotBefore().Equal(notBefore), "NotBefore should return the correct notBefore time")

	// Test NotAfter method
	assert.True(t, certModel.NotAfter().Equal(mockCert.NotAfter), "NotAfter should return the correct notAfter time")
	assert.True(t, certModel.NotAfter().Equal(notAfter), "NotAfter should return the correct notAfter time")
}
