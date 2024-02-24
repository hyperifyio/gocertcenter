// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/hyperifyio/gocertcenter/internal/modelutils"
	"io"
	"math/big"
	"strings"
	"testing"
	"time"
)

// mockKeyType for testing purposes
const mockKeyType models.KeyType = 0

// TestNewPrivateKey verifies the NewPrivateKey function correctly initializes a PrivateKey instance.
func TestNewPrivateKey(t *testing.T) {
	// Mock data for initialization
	serialNumber := models.NewSerialNumber(big.NewInt(123)) // Assuming SerialNumber is of type *big.Int
	keyType := mockKeyType
	mockData := "mockPrivateKeyData" // Example mock data, could be any type

	// Call the function under test
	privateKey := models.NewPrivateKey(serialNumber, keyType, mockData)

	bigIntSSerialNumber := privateKey.GetSerialNumber()

	// Verify the PrivateKey fields are correctly assigned
	if bigIntSSerialNumber.Cmp(serialNumber) != 0 {
		t.Errorf("serialNumber = %v, want %v", privateKey.GetSerialNumber(), serialNumber)
	}
	if privateKey.GetKeyType() != keyType {
		t.Errorf("keyType = %v, want %v", privateKey.GetSerialNumber(), keyType)
	}

	//// Since data is of type any, we assert its type and value where applicable
	//if data, ok := privateKey.data.(string); !ok || data != mockData {
	//	t.Errorf("data = %v, want %v", privateKey.data, mockData)
	//}

}

// TestPrivateKey_CreateCertificate tests the certificate creation functionality
func TestPrivateKey_CreateCertificate(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	privateKey, _ := modelutils.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P384)

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName: "Test Certificate",
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24),
		KeyUsage:  x509.KeyUsageCertSign,
	}

	// Self-sign the certificate for testing
	cert, err := privateKey.CreateCertificate(managers.NewCertificateManager(managers.NewRandomManager()), template, template)
	if err != nil {
		t.Fatalf("Failed to create certificate: %v", err)
	}
	if cert == nil {
		t.Fatal("Expected non-nil certificate")
	}
}

func TestPrivateKey_CreateCertificate_Success(t *testing.T) {
	// Set up the mock certificate manager with successful responses
	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			// Simulate successful certificate creation by returning dummy bytes
			return []byte("dummy certificate bytes"), nil
		},
		ParseCertificateFunc: func(certBytes []byte) (*x509.Certificate, error) {
			// Simulate successful parsing by returning a dummy certificate
			return &x509.Certificate{}, nil
		},
	}

	randomManager := managers.NewRandomManager()

	serialNumber, err := modelutils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	// Set up a PrivateKey instance for testing
	privateKey, err := modelutils.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P256)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}

	template := &x509.Certificate{ /* Fill in required fields for the template */ }
	parent := &x509.Certificate{ /* Optionally, fill in fields for the parent certificate */ }

	// Call CreateCertificate with the mock manager
	cert, err := privateKey.CreateCertificate(mockManager, template, parent)
	if err != nil {
		t.Fatalf("CreateCertificate failed: %v", err)
	}
	if cert == nil {
		t.Fatal("Expected non-nil certificate")
	}
}

func TestPrivateKey_CreateCertificate_Errors(t *testing.T) {
	// Generate a private key for the test
	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	privateKey, _ := modelutils.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P384)

	// Create an invalid template by setting invalid values
	invalidTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(-1), // Invalid serial number
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(-24 * time.Hour), // Invalid time range
	}

	_, err := privateKey.CreateCertificate(managers.NewCertificateManager(managers.NewRandomManager()), invalidTemplate, invalidTemplate)
	if err == nil {
		t.Fatal("Expected error when creating certificate with invalid template, got nil")
	}
}

func TestPrivateKey_CreateCertificate_Failure(t *testing.T) {
	// Set up the mock certificate manager to simulate an error in certificate creation
	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			// Simulate a failure in certificate creation
			return nil, fmt.Errorf("simulated certificate creation error")
		},
	}

	randomManager := managers.NewRandomManager()
	serialNumber, err := modelutils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, _ := modelutils.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P256)
	template := &x509.Certificate{ /* Fill in required fields for the template */ }
	parent := &x509.Certificate{ /* Optionally, fill in fields for the parent certificate */ }

	_, err = privateKey.CreateCertificate(mockManager, template, parent)
	if err == nil {
		t.Fatal("Expected an error in CreateCertificate, got nil")
	}
}

func TestPrivateKey_CreateCertificate_ParseFailure(t *testing.T) {
	// Simulate successful creation but failure in parsing
	mockManager := &mocks.MockCertificateManager{
		CreateCertificateFunc: func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
			return []byte("dummy certificate bytes"), nil
		},
		ParseCertificateFunc: func(certBytes []byte) (*x509.Certificate, error) {
			return nil, fmt.Errorf("simulated parsing error")
		},
	}

	randomManager := managers.NewRandomManager()
	serialNumber, err := modelutils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, _ := modelutils.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P256)
	template := &x509.Certificate{ /* Fill in required fields for the template */ }
	parent := &x509.Certificate{ /* Optionally, fill in fields for the parent certificate */ }

	_, err = privateKey.CreateCertificate(mockManager, template, parent)
	if err == nil || !strings.Contains(err.Error(), "failed to parse certificate after creating it") {
		t.Fatal("Expected a specific parsing error, got either no error or a different error")
	}
}

func TestPrivateKey_GetPublicKey_NilCase(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	unsupportedKeyType := "unsupported key type" // Simulate incorrect initialization

	privateKey := models.NewPrivateKey(serialNumber, mockKeyType, unsupportedKeyType)
	publicKey := privateKey.GetPublicKey()

	if publicKey != nil {
		t.Fatalf("Expected GetPublicKey to return nil for unsupported key type, got %v", publicKey)
	}
}

func TestPrivateKey_GetDTO(t *testing.T) {
	// Setup
	serialNumber := models.NewSerialNumber(big.NewInt(123))
	keyType := models.RSA            // Using RSA as an example
	mockData := "mockPrivateKeyData" // Placeholder for private key data, could be any type

	privateKey := models.NewPrivateKey(serialNumber, keyType, mockData)

	// Act
	dto := privateKey.GetDTO()

	// Assert
	expectedSerialNumberString := "123" // Assuming the String method of serialNumber just returns its int value as string
	expectedKeyTypeString := "RSA"      // Based on the KeyType.String() implementation

	if dto.Certificate != expectedSerialNumberString {
		t.Errorf("GetDTO().SerialNumber = %s, want %s", dto.Certificate, expectedSerialNumberString)
	}
	if dto.Type != expectedKeyTypeString {
		t.Errorf("GetDTO().KeyType = %s, want %s", dto.Type, expectedKeyTypeString)
	}

	// Since we don't have a meaningful assertion for the private key content in the DTO (it's empty ""),
	// we'll just check that it's indeed empty as expected.
	if dto.PrivateKey != "" {
		t.Errorf("GetDTO().PrivateKeyContent = %s, want %s", dto.PrivateKey, "")
	}
}

func TestPrivateKey_GetPublicKey_RSA(t *testing.T) {
	// Generate an RSA private key
	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}

	// Create a PrivateKey instance with the RSA private key
	privateKey := models.NewPrivateKey(models.NewSerialNumber(big.NewInt(1)), models.RSA, rsaPrivKey)

	// Get the public key from the PrivateKey instance
	publicKey := privateKey.GetPublicKey()

	// Type assert the returned public key to *rsa.PublicKey
	rsaPubKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		t.Fatalf("Expected public key type *rsa.PublicKey, got %T", publicKey)
	}

	// Verify the public key corresponds to the private key
	if rsaPubKey.N.Cmp(rsaPrivKey.N) != 0 || rsaPubKey.E != rsaPrivKey.E {
		t.Errorf("Public key does not match private key")
	}
}

func TestPrivateKey_GetPublicKey_Ed25519(t *testing.T) {
	// Generate an Ed25519 private key
	_, ed25519PrivKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate Ed25519 private key: %v", err)
	}

	// Create a PrivateKey instance with the Ed25519 private key
	privateKey := models.NewPrivateKey(models.NewSerialNumber(big.NewInt(1)), models.Ed25519, ed25519PrivKey)

	// Get the public key from the PrivateKey instance
	publicKey := privateKey.GetPublicKey()

	// Type assert the returned public key to ed25519.PublicKey
	ed25519PubKey, ok := publicKey.(ed25519.PublicKey)
	if !ok {
		t.Fatalf("Expected public key type ed25519.PublicKey, got %T", publicKey)
	}

	// Verify the public key corresponds to the private key
	if !ed25519.Verify(ed25519PubKey, []byte("test message"), ed25519.Sign(ed25519PrivKey, []byte("test message"))) {
		t.Errorf("Public key does not verify signature made with private key")
	}
}
