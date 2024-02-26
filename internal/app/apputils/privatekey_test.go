// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// For testing an invalid keyType, you can directly use a value outside the range of defined KeyTypes:
const InvalidKeyType = 999 // A value not represented in the KeyType enum

func TestGeneratePrivateKey(t *testing.T) {
	organization := "testOrg"

	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)

	keyTypes := []appmodels.KeyType{
		appmodels.RSA_1024, appmodels.RSA_2048, appmodels.RSA_3072, appmodels.RSA_4096,
		appmodels.ECDSA_P224, appmodels.ECDSA_P256, appmodels.ECDSA_P384, appmodels.ECDSA_P521, appmodels.Ed25519}
	for _, kt := range keyTypes {
		privateKey, err := apputils.GeneratePrivateKey(
			organization, []appmodels.ISerialNumber{serialNumber}, kt) // RSA bits size is only relevant for RSA keys
		if err != nil {
			t.Fatalf("Failed to generate private key for %v: %v", kt, err)
		}
		if privateKey == nil {
			t.Fatalf("Expected private key, got: nil")
		}

		// switch kt {
		// case models.RSA:
		//	if _, ok := privateKey.data.(*rsa.PrivateKey); !ok {
		//		t.Errorf("Expected RSA private key, got %T", privateKey.data)
		//	}
		// case models.ECDSA_P224:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		// case models.ECDSA_P256:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		// case models.ECDSA_P384:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		// case models.ECDSA_P521:
		//	if _, ok := privateKey.data.(*ecdsa.PrivateKey); !ok {
		//		t.Errorf("Expected ECDSA private key, got %T", privateKey.data)
		//	}
		// case models.Ed25519:
		//	if _, ok := privateKey.data.(ed25519.PrivateKey); !ok {
		//		t.Errorf("Expected Ed25519 private key, got %T", privateKey.data)
		//	}
		// }

	}
}

func TestGeneratePrivateKey_InvalidKeyType(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)
	_, err := apputils.GeneratePrivateKey(organization, []appmodels.ISerialNumber{serialNumber}, InvalidKeyType) // Using the invalid KeyType here
	if err == nil {
		t.Fatal("Expected GeneratePrivateKey to return an error for an invalid keyType, but it did not")
	}
}

func TestGenerateRSAPrivateKey(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)

	privateKey, err := apputils.GenerateRSAPrivateKey(organization, []appmodels.ISerialNumber{serialNumber}, appmodels.RSA_2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}
	if privateKey == nil {
		t.Fatalf("Expected private key, got: nil")
	}

	// if _, ok := privateKey.data.(*rsa.PrivateKey); !ok {
	//	t.Errorf("Expected RSA private key, got %T", privateKey.data)
	// }

}

// TestGenerateECDSAPrivateKey checks if a new ECDSA private key is generated without error
func TestGenerateECDSAPrivateKey(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, err := apputils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, err := apputils.GenerateECDSAPrivateKey(organization, []appmodels.ISerialNumber{serialNumber}, appmodels.ECDSA_P384)
	if err != nil {
		t.Fatalf("Failed to generate private key: %v", err)
	}
	if privateKey == nil {
		t.Fatal("Expected non-nil private key")
	}
	// if privateKey.data == nil {
	//	t.Fatal("Expected non-nil internal private key data")
	// }
}

func TestGenerateEd25519PrivateKey(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)

	privateKey, err := apputils.GenerateEd25519PrivateKey(organization, []appmodels.ISerialNumber{serialNumber})
	if err != nil {
		t.Fatalf("Failed to generate Ed25519 private key: %v", err)
	}
	if privateKey == nil {
		t.Fatal("Expected non-nil private key")
	}

	// if _, ok := privateKey.data.(ed25519.PrivateKey); !ok {
	//	t.Errorf("Expected Ed25519 private key, got %T", privateKey.data)
	// }
}

// TestPrivateKey_GetSerialNumber verifies that GetSerialNumber returns the correct serial number
func TestPrivateKey_GetSerialNumber(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	expectedSerialNumber, _ := apputils.GenerateSerialNumber(randomManager)
	privateKey := appmodels.NewPrivateKey(
		organization, []appmodels.ISerialNumber{expectedSerialNumber},
		0,
		nil,
	)

	actualSerialNumber := privateKey.GetSerialNumber()

	if actualSerialNumber.Cmp(expectedSerialNumber) != 0 {
		t.Errorf("Expected serial number %v, got %v", expectedSerialNumber, actualSerialNumber)
	}
}

// TestPrivateKey_GetPublicKey checks if GetPublicKey returns a valid public key from the private key
func TestPrivateKey_GetPublicKey_ECDSA_P384(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)
	key, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	privateKey := appmodels.NewPrivateKey(organization, []appmodels.ISerialNumber{serialNumber}, appmodels.ECDSA_P384, key)

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
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)
	privateKey, err := apputils.GenerateEd25519PrivateKey(organization, []appmodels.ISerialNumber{serialNumber})
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

func TestMarshalPrivateKeyAsPEM_RSA(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	privateKey := &rsa.PrivateKey{}
	expectedPEMBytes := []byte("rsa private key pem")
	expectedResult := []byte("RSA PRIVATE KEY")

	mockManager.On("MarshalPKCS1PrivateKey", privateKey).Return(expectedPEMBytes)
	mockManager.On("EncodePEMToMemory", mock.AnythingOfType("*pem.Block")).Return(expectedResult)

	result, err := apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)
	assert.NoError(t, err)
	assert.Equal(t, result, expectedResult)
}

func TestMarshalPrivateKeyAsPEM_ECDSA(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	privateKey := &ecdsa.PrivateKey{}

	expectedResult := []byte("-----BEGIN EC PRIVATE KEY-----")

	expectedPEMBytes := []byte("ecdsa private key pem")

	mockManager.On("MarshalECPrivateKey", privateKey).Return(expectedPEMBytes, nil)
	mockManager.On("EncodePEMToMemory", mock.AnythingOfType("*pem.Block")).Return(expectedResult)

	result, err := apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result, expectedResult)
}

func TestMarshalPrivateKeyAsPEM_Ed25519(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	privateKey := ed25519.PrivateKey{}
	expectedResult := []byte("-----BEGIN PRIVATE KEY-----")

	expectedPEMBytes := []byte("ed25519 private key pem")

	mockManager.On("MarshalPKCS8PrivateKey", privateKey).Return(expectedPEMBytes, nil)
	mockManager.On("EncodePEMToMemory", mock.AnythingOfType("*pem.Block")).Return(expectedResult)

	result, err := apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result, expectedResult)
}

func TestMarshalPrivateKeyAsPEM_UnsupportedKeyType(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	unsupportedKeyType := "unsupported key" // Use a simple string to represent an unsupported key type

	pemBytes, err := apputils.MarshalPrivateKeyAsPEM(mockManager, unsupportedKeyType)
	assert.Error(t, err)
	assert.Nil(t, pemBytes)
}

// Additional tests for error scenarios, e.g., when the underlying manager methods return errors
func TestMarshalPrivateKeyAsPEM_ECDSAError(t *testing.T) {

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	assert.NotNil(t, privateKey)

	mockManager := new(commonmocks.MockCertificateManager)
	mockManager.On("MarshalECPrivateKey", privateKey).Return(nil, fmt.Errorf("mock error"))

	pemBytes, err := apputils.MarshalPrivateKeyAsPEM(mockManager, &privateKey)

	assert.Error(t, err)
	assert.Nil(t, pemBytes)

}
func TestGeneratePrivateKey_EmptyOrganization(t *testing.T) {
	_, err := apputils.GeneratePrivateKey(
		"", // Empty organization
		[]appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(1))},
		appmodels.RSA_2048,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "organization: must not be empty")
}

func TestGeneratePrivateKey_NoCertificates(t *testing.T) {
	_, err := apputils.GeneratePrivateKey(
		"TestOrg",
		nil, // No certificates
		appmodels.RSA_2048,
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "certificates: must have at least one serial number")
}

func TestMarshalPrivateKeyAsPEM_RSAKeyNil(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	key := (*rsa.PrivateKey)(nil)
	mockManager.On("MarshalPKCS1PrivateKey", key).Return(nil)
	_, err := apputils.MarshalPrivateKeyAsPEM(mockManager, key)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal RSA private key: got nil")
}

func TestMarshalPrivateKeyAsPEM_ECDSAKeyMarshalError(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	mockManager.On("MarshalECPrivateKey", mock.Anything).Return(nil, fmt.Errorf("marshal error"))
	_, err := apputils.MarshalPrivateKeyAsPEM(mockManager, &ecdsa.PrivateKey{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal ECDSA private key")
}

func TestMarshalPrivateKeyAsPEM_ECDSAKeyMarshalError_ReturnsBytesNil(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	mockManager.On("MarshalECPrivateKey", mock.Anything).Return(nil, nil)
	_, err := apputils.MarshalPrivateKeyAsPEM(mockManager, &ecdsa.PrivateKey{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal ECDSA private key: got nil")
}

func TestMarshalPrivateKeyAsPEM_Ed25519KeyMarshalError(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	mockManager.On("MarshalPKCS8PrivateKey", mock.Anything).Return(nil, fmt.Errorf("marshal error"))
	_, err := apputils.MarshalPrivateKeyAsPEM(mockManager, ed25519.PrivateKey{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to marshal Ed25519 private key to PKCS#8")
}

func TestMarshalPrivateKeyAsPEM_Ed25519KeyMarshalError_ReturnsBytesNil(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	mockManager.On("MarshalPKCS8PrivateKey", mock.Anything).Return(nil, nil)
	_, err := apputils.MarshalPrivateKeyAsPEM(mockManager, ed25519.PrivateKey{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "MarshalPrivateKeyAsPEM: ed25519: failed to marshal ECDSA private key: got nil")
}

func TestMarshalPrivateKeyAsPEM_Ed25519KeyMarshalError_ReturnsInvalidPEM(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	mockManager.On("MarshalPKCS8PrivateKey", mock.Anything).Return(nil, nil)
	_, err := apputils.MarshalPrivateKeyAsPEM(mockManager, ed25519.PrivateKey{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "MarshalPrivateKeyAsPEM: ed25519: failed to marshal ECDSA private key: got nil")
}

func TestMarshalPrivateKeyAsPEM_EncodePEMToMemoryFails(t *testing.T) {
	// Create a mock certificate manager
	mockManager := new(commonmocks.MockCertificateManager)

	// Generate an Ed25519 private key for testing
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	assert.NoError(t, err, "Failed to generate Ed25519 private key")

	// Configure the mock to simulate EncodePEMToMemory returning nil
	mockManager.On("MarshalPKCS8PrivateKey", privateKey).Return([]byte("dummy data"), nil) // Assume successful key marshaling
	mockManager.On("EncodePEMToMemory", mock.Anything).Return(nil)                         // Simulate failure in PEM encoding

	// Call MarshalPrivateKeyAsPEM with the mock manager and the private key
	_, err = apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)

	// Assert that an error was returned
	assert.Error(t, err, "Expected an error due to failure in encoding PEM")
	assert.Contains(t, err.Error(), "could not encode to PEM", "Error message should indicate failure to encode PEM")

	// Verify that the mock expectations were met
	mockManager.AssertExpectations(t)
}
