// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package modelutils_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/hyperifyio/gocertcenter/internal/modelutils"
)

// For testing an invalid keyType, you can directly use a value outside the range of defined KeyTypes:
const InvalidKeyType = 999 // A value not represented in the KeyType enum

func TestGeneratePrivateKey(t *testing.T) {

	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)

	keyTypes := []models.KeyType{models.RSA, models.ECDSA_P224, models.ECDSA_P256, models.ECDSA_P384, models.ECDSA_P521, models.Ed25519}
	for _, kt := range keyTypes {
		privateKey, err := modelutils.GeneratePrivateKey(serialNumber, kt, 2048) // RSA bits size is only relevant for RSA keys
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
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	_, err := modelutils.GeneratePrivateKey(serialNumber, InvalidKeyType, 2048) // Using the invalid KeyType here
	if err == nil {
		t.Fatal("Expected GeneratePrivateKey to return an error for an invalid keyType, but it did not")
	}
}

func TestGenerateRSAPrivateKey(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	rsaBits := 2048

	privateKey, err := modelutils.GenerateRSAPrivateKey(serialNumber, rsaBits)
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
	serialNumber, err := modelutils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Failed to generate serial number: %v", err)
	}

	privateKey, err := modelutils.GenerateECDSAPrivateKey(serialNumber, models.ECDSA_P384)
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
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)

	privateKey, err := modelutils.GenerateEd25519PrivateKey(serialNumber)
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
	expectedSerialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	privateKey := models.NewPrivateKey(
		expectedSerialNumber,
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
	randomManager := managers.NewRandomManager()
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
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
	serialNumber, _ := modelutils.GenerateSerialNumber(randomManager)
	privateKey, err := modelutils.GenerateEd25519PrivateKey(serialNumber)
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
