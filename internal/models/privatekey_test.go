// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"io"
	"math/big"
	"strings"
	"testing"
	"time"
)

// mockKeyType for testing purposes
const mockKeyType models.KeyType = 0

// For testing an invalid keyType, you can directly use a value outside the range of defined KeyTypes:
const InvalidKeyType = 999 // A value not represented in the KeyType enum

// TestNewPrivateKey verifies the NewPrivateKey function correctly initializes a PrivateKey instance.
func TestNewPrivateKey(t *testing.T) {
	// Mock data for initialization
	serialNumber := big.NewInt(123) // Assuming SerialNumber is of type *big.Int
	keyType := mockKeyType
	mockData := "mockPrivateKeyData" // Example mock data, could be any type

	// Call the function under test
	privateKey := models.NewPrivateKey(serialNumber, keyType, mockData)

	bigIntSSerialNumber := (*big.Int)(privateKey.GetSerialNumber())

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

func TestGeneratePrivateKey(t *testing.T) {

	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)

	keyTypes := []models.KeyType{models.RSA, models.ECDSA_P224, models.ECDSA_P256, models.ECDSA_P384, models.ECDSA_P521, models.Ed25519}
	for _, kt := range keyTypes {
		privateKey, err := models.GeneratePrivateKey(serialNumber, kt, 2048) // RSA bits size is only relevant for RSA keys
		if err != nil {
			t.Fatalf("Failed to generate private key for %v: %v", kt, err)
		}
		if privateKey == nil {
			t.Fatalf("Expected private key, got: nil")
		}

		//switch kt {
		//case models.RSA:
		//	if _, ok := privateKey.data.(*rsa.PrivateKey); !ok {
		//		t.Errorf("Expected RSA private key, got %T", privateKey.data)
		//	}
		//case models.ECDSA_P224:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		//case models.ECDSA_P256:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		//case models.ECDSA_P384:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		//case models.ECDSA_P521:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		//case models.Ed25519:
		//	if _, ok := privateKey.data.(ed25519.PrivateKey); !ok {
		//		t.Errorf("Expected Ed25519 private key, got %T", privateKey.data)
		//	}
		//}

	}
}

func TestGeneratePrivateKey_InvalidKeyType(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	_, err := models.GeneratePrivateKey(serialNumber, InvalidKeyType, 2048) // Using the invalid KeyType here
	if err == nil {
		t.Fatal("Expected GeneratePrivateKey to return an error for an invalid keyType, but it did not")
	}
}

func TestGenerateRSAPrivateKey(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	rsaBits := 2048

	privateKey, err := models.GenerateRSAPrivateKey(serialNumber, rsaBits)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}
	if privateKey == nil {
		t.Fatalf("Expected private key, got: nil")
	}

	//if _, ok := privateKey.data.(*rsa.PrivateKey); !ok {
	//	t.Errorf("Expected RSA private key, got %T", privateKey.data)
	//}

}

// TestGenerateECDSAPrivateKey checks if a new ECDSA private key is generated without error
func TestGenerateECDSAPrivateKey(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, err := models.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P384)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	if privateKey == nil {
		t.Fatal("Expected non-nil private key")
	}
	//if privateKey.data == nil {
	//	t.Fatal("Expected non-nil internal private key data")
	//}
}

func TestGenerateEd25519PrivateKey(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)

	privateKey, err := models.GenerateEd25519PrivateKey(serialNumber)
	if err != nil {
		t.Fatalf("Failed to generate Ed25519 private key: %v", err)
	}
	if privateKey == nil {
		t.Fatal("Expected non-nil private key")
	}

	//if _, ok := privateKey.data.(ed25519.PrivateKey); !ok {
	//	t.Errorf("Expected Ed25519 private key, got %T", privateKey.data)
	//}
}

// TestPrivateKey_GetSerialNumber verifies that GetSerialNumber returns the correct serial number
func TestPrivateKey_GetSerialNumber(t *testing.T) {
	randomManager := managers.NewRandomManager()
	expectedSerialNumber, _ := models.NewSerialNumber(randomManager)
	privateKey := models.NewPrivateKey(
		expectedSerialNumber,
		0,
		nil,
	)

	actualSerialNumber := privateKey.GetSerialNumber()

	bigIntSSerialNumber := (*big.Int)(actualSerialNumber)

	if bigIntSSerialNumber.Cmp((*big.Int)(expectedSerialNumber)) != 0 {
		t.Errorf("Expected serial number %v, got %v", expectedSerialNumber, actualSerialNumber)
	}
}

// TestPrivateKey_GetPublicKey checks if GetPublicKey returns a valid public key from the private key
func TestPrivateKey_GetPublicKey_ECDSA_P384(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	key, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	privateKey := models.NewPrivateKey(serialNumber, models.ECDSA_P384, key)

	publicKeyAny := privateKey.GetPublicKey()
	publicKey, ok := publicKeyAny.(*ecdsa.PublicKey) // Type assertion
	if !ok {
		t.Fatalf("Expected ECDSA public key, got different type")
	}
	if publicKey.X == nil || publicKey.Y == nil {
		t.Fatal("Expected non-nil components of the public key")
	}
}

// TestPrivateKey_GetPublicKey_Ed25519 checks if GetPublicKey returns a valid public key from the private key
func TestPrivateKey_GetPublicKey_Ed25519(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	privateKey, err := models.GenerateEd25519PrivateKey(serialNumber)
	if err != nil {
		t.Fatalf("Could not generate private key: %v", err)
	}

	publicKeyAny := privateKey.GetPublicKey()
	publicKey, ok := publicKeyAny.(ed25519.PublicKey) // Corrected type assertion
	if !ok {
		t.Fatalf("Expected Ed25519 public key, got different type")
	}
	// Check the length of the Ed25519 public key (should be 32 bytes for Ed25519)
	if len(publicKey) != ed25519.PublicKeySize {
		t.Fatalf("Expected Ed25519 public key size of %d, got %d", ed25519.PublicKeySize, len(publicKey))
	}
}

// TestPrivateKey_CreateCertificate tests the certificate creation functionality
func TestPrivateKey_CreateCertificate(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	privateKey, _ := models.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P384)

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

	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	// Set up a PrivateKey instance for testing
	privateKey, err := models.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P256)
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
	serialNumber, _ := models.NewSerialNumber(randomManager)
	privateKey, _ := models.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P384)

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
	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, _ := models.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P256)
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
	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, _ := models.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P256)
	template := &x509.Certificate{ /* Fill in required fields for the template */ }
	parent := &x509.Certificate{ /* Optionally, fill in fields for the parent certificate */ }

	_, err = privateKey.CreateCertificate(mockManager, template, parent)
	if err == nil || !strings.Contains(err.Error(), "failed to parse certificate after creating it") {
		t.Fatal("Expected a specific parsing error, got either no error or a different error")
	}
}

func TestPrivateKey_GetPublicKey_NilCase(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := models.NewSerialNumber(randomManager)
	unsupportedKeyType := "unsupported key type" // Simulate incorrect initialization

	privateKey := models.NewPrivateKey(serialNumber, mockKeyType, unsupportedKeyType)
	publicKey := privateKey.GetPublicKey()

	if publicKey != nil {
		t.Fatalf("Expected GetPublicKey to return nil for unsupported key type, got %v", publicKey)
	}
}
