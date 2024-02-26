// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	expectedPEMBytes := []byte("rsa private key pem")
	mockManager.On("MarshalPKCS1PrivateKey", privateKey).Return(expectedPEMBytes)

	pemBytes, err := apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)
	assert.NoError(t, err)
	assert.Contains(t, string(pemBytes), "RSA PRIVATE KEY")
}

func TestMarshalPrivateKeyAsPEM_Ed25519(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	_, privateKey, err := ed25519.GenerateKey(rand.Reader)
	assert.NoError(t, err)

	expectedPEMBytes := []byte("ed25519 private key pem")
	mockManager.On("MarshalPKCS8PrivateKey", privateKey).Return(expectedPEMBytes, nil)

	pemBytes, err := apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)
	assert.NoError(t, err)
	assert.NotNil(t, pemBytes)

	// Further validate the output is in proper PEM format. This checks if the output
	// starts with the expected PEM header for a PRIVATE KEY.
	expectedPEMHeader := "-----BEGIN PRIVATE KEY-----"
	if !strings.HasPrefix(string(pemBytes), expectedPEMHeader) {
		t.Errorf("PEM does not start with expected header %q", expectedPEMHeader)
	}

}

func TestMarshalPrivateKeyAsPEM_UnsupportedKeyType(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	unsupportedKeyType := "unsupported key" // Use a simple string to represent an unsupported key type

	pemBytes, err := apputils.MarshalPrivateKeyAsPEM(mockManager, unsupportedKeyType)
	assert.Error(t, err)
	assert.Nil(t, pemBytes)
}

func TestMarshalPrivateKeyAsPEM_ECDSA(t *testing.T) {
	mockManager := new(commonmocks.MockCertificateManager)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	expectedPEMBytes := []byte("ecdsa private key pem")
	mockManager.On("MarshalECPrivateKey", privateKey).Return(expectedPEMBytes, nil)

	pemBytes, err := apputils.MarshalPrivateKeyAsPEM(mockManager, privateKey)
	assert.NoError(t, err)
	assert.NotNil(t, pemBytes)

	// Further validate the output is in proper PEM format. This checks if the output
	// starts with the expected PEM header for a PRIVATE KEY.
	expectedPEMHeader := "-----BEGIN EC PRIVATE KEY-----"
	if !strings.HasPrefix(string(pemBytes), expectedPEMHeader) {
		t.Errorf("PEM does not start with expected header %q\nreceived: %q", expectedPEMHeader, string(pemBytes))
	}
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
