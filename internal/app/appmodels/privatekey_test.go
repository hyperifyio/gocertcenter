// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"math/big"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// mockKeyType for testing purposes
const mockKeyType appmodels.KeyType = 0

// TestNewPrivateKey verifies the NewPrivateKey function correctly initializes a PrivateKey instance.
func TestNewPrivateKey(t *testing.T) {
	// Mock data for initialization
	organization := "testOrg"
	serialNumber := appmodels.NewSerialNumber(big.NewInt(123)) // Assuming SerialNumber is of type *big.Int
	keyType := mockKeyType
	mockData := "mockPrivateKeyData" // Example mock data, could be any type

	// Call the function under test
	privateKey := appmodels.NewPrivateKey(organization, []appmodels.ISerialNumber{serialNumber}, keyType, mockData)

	bigIntSSerialNumber := privateKey.GetSerialNumber()

	// Verify the PrivateKey fields are correctly assigned
	if bigIntSSerialNumber.Cmp(serialNumber) != 0 {
		t.Errorf("serialNumber = %v, want %v", privateKey.GetSerialNumber(), serialNumber)
	}
	if privateKey.GetKeyType() != keyType {
		t.Errorf("keyType = %v, want %v", privateKey.GetSerialNumber(), keyType)
	}

	// // Since data is of type any, we assert its type and value where applicable
	// if data, ok := privateKey.data.(string); !ok || data != mockData {
	//	t.Errorf("data = %v, want %v", privateKey.data, mockData)
	// }

}

func TestPrivateKey_GetPublicKey_NilCase(t *testing.T) {
	organization := "testOrg"
	randomManager := managers.NewRandomManager()
	serialNumber, _ := apputils.GenerateSerialNumber(randomManager)
	unsupportedKeyType := "unsupported key type" // Simulate incorrect initialization

	privateKey := appmodels.NewPrivateKey(organization, []appmodels.ISerialNumber{serialNumber}, mockKeyType, unsupportedKeyType)
	publicKey := privateKey.GetPublicKey()

	if publicKey != nil {
		t.Fatalf("Expected GetPublicKey to return nil for unsupported key type, got %v", publicKey)
	}
}

func TestPrivateKey_GetPublicKey_RSA(t *testing.T) {
	organization := "testOrg"

	// Generate an RSA private key
	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}

	// Create a PrivateKey instance with the RSA private key
	privateKey := appmodels.NewPrivateKey(
		organization,
		[]appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(1))},
		appmodels.RSA_2048,
		rsaPrivKey,
	)

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
	organization := "testOrg"

	_, ed25519PrivKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate Ed25519 private key: %v", err)
	}

	// Create a PrivateKey instance with the Ed25519 private key
	privateKey := appmodels.NewPrivateKey(
		organization,
		[]appmodels.ISerialNumber{appmodels.NewSerialNumber(big.NewInt(1))},
		appmodels.Ed25519, ed25519PrivKey,
	)

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
