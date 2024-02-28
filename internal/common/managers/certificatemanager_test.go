// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestNewCertificateManager(t *testing.T) {
	mockRandomManager := commonmocks.NewMockRandomManager()

	// Test initialization with a provided random manager
	managerWithMock := managers.NewCertificateManager(mockRandomManager)
	if managerWithMock.GetRandomManager() != mockRandomManager {
		t.Errorf("Expected random manager to be the mock instance, got different instance")
	}

	// Test initialization without providing a random manager (should default to NewRandomManager)
	managerWithDefault := managers.NewCertificateManager(nil)
	if _, ok := managerWithDefault.GetRandomManager().(*managers.RandomManager); !ok {
		t.Errorf("Expected default random manager to be of type *RandomManager, got %T", managerWithDefault.GetRandomManager())
	}
}

func TestCertificateManager_CreateAndParseCertificate(t *testing.T) {

	mockRandomManager := commonmocks.NewMockRandomManager()

	manager := managers.NewCertificateManager(mockRandomManager)

	// Create a minimal certificate template
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "test"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour),
	}

	// Use a self-signed template for simplicity; in practice, you'd have a CA sign it
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate test RSA keys: %v", err)
	}

	certBytes, err := manager.CreateCertificate(rand.Reader, template, template, &privKey.PublicKey, privKey)
	if err != nil {
		t.Fatalf("CreateCertificate failed: %v", err)
	}

	cert, err := manager.ParseCertificate(certBytes)
	if err != nil {
		t.Fatalf("ParseCertificate failed: %v", err)
	}

	if cert.Subject.CommonName != "test" {
		t.Errorf("Parsed certificate CommonName = %s, want %s", cert.Subject.CommonName, "test")
	}
}

func TestCertificateManager_MarshalPKCS1PrivateKey(t *testing.T) {
	manager := managers.NewCertificateManager(nil)

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Failed to generate RSA private key")

	pkcs1Bytes := manager.MarshalPKCS1PrivateKey(privKey)
	assert.NotEmpty(t, pkcs1Bytes, "PKCS1 marshaled data should not be empty")
}

func TestCertificateManager_MarshalECPrivateKey(t *testing.T) {
	manager := managers.NewCertificateManager(nil)

	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err, "Failed to generate ECDSA private key")

	ecBytes, err := manager.MarshalECPrivateKey(privKey)
	require.NoError(t, err, "Failed to marshal ECDSA private key")
	assert.NotEmpty(t, ecBytes, "ECDSA marshaled data should not be empty")
}

func TestCertificateManager_MarshalPKCS8PrivateKey_RSA(t *testing.T) {
	manager := managers.NewCertificateManager(nil)

	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Failed to generate RSA private key")

	pkcs8Bytes, err := manager.MarshalPKCS8PrivateKey(rsaPrivKey)
	require.NoError(t, err, "Failed to marshal RSA private key to PKCS8")
	assert.NotEmpty(t, pkcs8Bytes, "PKCS8 marshaled data for RSA key should not be empty")
}

func TestCertificateManager_MarshalPKCS8PrivateKey_ECDSA(t *testing.T) {
	manager := managers.NewCertificateManager(nil)

	ecdsaPrivKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err, "Failed to generate ECDSA private key")

	pkcs8Bytes, err := manager.MarshalPKCS8PrivateKey(ecdsaPrivKey)
	require.NoError(t, err, "Failed to marshal ECDSA private key to PKCS8")
	assert.NotEmpty(t, pkcs8Bytes, "PKCS8 marshaled data for ECDSA key should not be empty")
}

func TestCertificateManager_ParseECPrivateKey(t *testing.T) {
	// Generate an ECDSA private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err, "Failed to generate ECDSA private key")

	// Convert the private key to DER format
	der, err := x509.MarshalECPrivateKey(privateKey)
	assert.NoError(t, err, "Failed to marshal ECDSA private key to DER")

	// Assume CertificateManager is correctly instantiated
	certManager := managers.CertificateManager{}

	// Call ParseECPrivateKey with the DER bytes
	parsedKey, err := certManager.ParseECPrivateKey(der)
	assert.NoError(t, err, "Failed to parse ECDSA private key from DER")

	// Verify the parsed key
	assert.Equal(t, privateKey.D, parsedKey.D, "Private key 'D' parameters do not match")
	assert.Equal(t, privateKey.PublicKey.X, parsedKey.PublicKey.X, "Public key 'X' parameters do not match")
	assert.Equal(t, privateKey.PublicKey.Y, parsedKey.PublicKey.Y, "Public key 'Y' parameters do not match")
}

func TestCertificateManager_EncodePEMToMemory(t *testing.T) {
	// Create a sample PEM block
	block := &pem.Block{
		Type:  "TEST CERTIFICATE",
		Bytes: []byte("test content"),
	}

	// Create an instance of CertificateManager
	certManager := managers.CertificateManager{}

	// Encode the PEM block to memory
	encoded := certManager.EncodePEMToMemory(block)

	// Convert the encoded bytes back to a string for easy assertion
	encodedStr := string(encoded)

	// Check that the encoded data is in PEM format
	require.Contains(t, encodedStr, "-----BEGIN TEST CERTIFICATE-----", "PEM header not found")
	require.Contains(t, encodedStr, "-----END TEST CERTIFICATE-----", "PEM footer not found")
	require.Contains(t, encodedStr, "dGVzdCBjb250ZW50", "PEM content does not match expected")

	// Optionally, decode the PEM to verify its integrity
	decodedBlock, _ := pem.Decode(encoded)
	require.NotNil(t, decodedBlock, "Failed to decode encoded PEM")
	require.Equal(t, block.Type, decodedBlock.Type, "Decoded PEM type does not match")
	require.Equal(t, block.Bytes, decodedBlock.Bytes, "Decoded PEM content does not match")
}

func TestCertificateManager_DecodePEM(t *testing.T) {
	// Create a sample PEM block
	block := &pem.Block{
		Type:  "TEST CERTIFICATE",
		Bytes: []byte("test content"),
	}

	// Encode the block to PEM format
	encodedPEM := pem.EncodeToMemory(block)

	// Create an instance of CertificateManager
	certManager := managers.CertificateManager{}

	// Decode the PEM encoded block
	decodedBlock, rest := certManager.DecodePEM(encodedPEM)

	// Assert the decoded block is not nil and matches the original
	require.NotNil(t, decodedBlock, "Decoded block should not be nil")
	require.Equal(t, block.Type, decodedBlock.Type, "Decoded block type mismatch")
	require.Equal(t, block.Bytes, decodedBlock.Bytes, "Decoded block content mismatch")

	// Assert that there is no remaining data after the decoded block
	require.Empty(t, rest, "There should be no remaining data after the decoded block")
}

func TestCertificateManager_ParsePKCS8PrivateKey(t *testing.T) {
	// Generate an ECDSA private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err, "Failed to generate ECDSA private key")

	// Convert the private key to PKCS#8 DER format
	der, err := x509.MarshalPKCS8PrivateKey(privateKey)
	require.NoError(t, err, "Failed to convert private key to PKCS#8 DER format")

	certManager := managers.CertificateManager{}

	// Parse the PKCS#8 DER-encoded private key
	parsedKey, err := certManager.ParsePKCS8PrivateKey(der)
	require.NoError(t, err, "Failed to parse PKCS#8 private key")

	// Assert the parsed key is not nil and is of *ecdsa.PrivateKey type
	assert.NotNil(t, parsedKey, "Parsed key should not be nil")
	_, ok := parsedKey.(*ecdsa.PrivateKey)
	assert.True(t, ok, "Parsed key is not of expected *ecdsa.PrivateKey type")
}

func TestCertificateManager_ParsePKCS1PrivateKey(t *testing.T) {
	// Generate an RSA private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err, "Failed to generate RSA private key")

	// Convert the private key to PKCS#1 DER format
	der := x509.MarshalPKCS1PrivateKey(privateKey)

	certManager := managers.CertificateManager{}

	// Parse the PKCS#1 DER-encoded private key
	parsedKey, err := certManager.ParsePKCS1PrivateKey(der)
	require.NoError(t, err, "Failed to parse PKCS#1 private key")

	// Assert the parsed key is not nil and matches the original
	assert.NotNil(t, parsedKey, "Parsed key should not be nil")
	assert.Equal(t, privateKey, parsedKey, "Parsed key does not match the original")
}
