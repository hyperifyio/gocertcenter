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
