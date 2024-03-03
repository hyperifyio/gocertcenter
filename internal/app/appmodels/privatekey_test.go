// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// mockKeyType for testing purposes
const mockKeyType appmodels.KeyType = 0

// TestNewPrivateKey verifies the NewPrivateKey function correctly initializes a PrivateKeyModel instance.
func TestNewPrivateKey(t *testing.T) {
	// Mock data for initialization
	organization := "testOrg"
	serialNumber := appmodels.NewSerialNumber(13)
	keyType := mockKeyType
	mockData := "mockPrivateKeyData" // Example mock data, could be any type

	// Call the function under test
	privateKey := appmodels.NewPrivateKey(organization, serialNumber, keyType, mockData)

	bigIntSSerialNumber := privateKey.SerialNumber()

	// Verify the PrivateKey fields are correctly assigned
	if bigIntSSerialNumber.Cmp(serialNumber) != 0 {
		t.Errorf("serialNumber = %v, want %v", privateKey.SerialNumber(), serialNumber)
	}
	if privateKey.KeyType() != keyType {
		t.Errorf("keyType = %v, want %v", privateKey.SerialNumber(), keyType)
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

	privateKey := appmodels.NewPrivateKey(organization, serialNumber, mockKeyType, unsupportedKeyType)
	publicKey := privateKey.PublicKey()

	if publicKey != nil {
		t.Fatalf("Expected PublicKey to return nil for unsupported key type, got %v", publicKey)
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
		appmodels.NewSerialNumber(1),
		appmodels.RSA_2048,
		rsaPrivKey,
	)

	// Get the public key from the PrivateKey instance
	publicKey := privateKey.PublicKey()

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
		appmodels.NewSerialNumber(1),
		appmodels.Ed25519, ed25519PrivKey,
	)

	// Get the public key from the PrivateKey instance
	publicKey := privateKey.PublicKey()

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

func TestPrivateKey_Methods(t *testing.T) {
	organization := "testOrg"
	serialNumber := appmodels.NewSerialNumber(3)
	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}
	privateKey := appmodels.NewPrivateKey(organization, serialNumber, appmodels.RSA_2048, rsaPrivKey)

	// Test SerialNumber
	if gotSerialNumber := privateKey.SerialNumber(); gotSerialNumber != serialNumber {
		t.Errorf("Expected organization ID %s, got %s", organization, gotSerialNumber)
	}

	// Test OrganizationID
	if gotOrgID := privateKey.OrganizationID(); gotOrgID != organization {
		t.Errorf("Expected organization ID %s, got %s", organization, gotOrgID)
	}

	// Test PrivateKey
	if gotPrivKey := privateKey.PrivateKey(); gotPrivKey != rsaPrivKey {
		t.Errorf("Expected private key %v, got %v", rsaPrivKey, gotPrivKey)
	}
}

func TestPrivateKey_NewPrivateKey_WithSingleSerialNumber(t *testing.T) {
	organization := "rootOrg"
	// Simulate root certificate with only one serial number
	serialNumber := appmodels.NewSerialNumber(1)
	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("Failed to generate RSA private key: %v", err)
	}
	privateKey := appmodels.NewPrivateKey(organization, serialNumber, appmodels.RSA_2048, rsaPrivKey)
	assert.NotNil(t, privateKey)
}

func TestPrivateKey_GetSerialNumber_WithZeroCertificates(t *testing.T) {
	organization := "testOrg"
	// Create a PrivateKey instance with no certificates
	privateKey := appmodels.NewPrivateKey(organization, nil, appmodels.RSA_2048, "mockData")

	// Test SerialNumber when certificates slice is empty
	serialNumber := privateKey.SerialNumber()
	if serialNumber != nil {
		t.Errorf("Expected nil serial number for PrivateKey with no certificate, got %v", serialNumber)
	}
}

func TestPrivateKey_GetPublicKey_ECDSA(t *testing.T) {
	organization := "testOrg"

	// Generate an ECDSA private key
	ecdsaPrivKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatalf("Failed to generate ECDSA private key: %v", err)
	}

	// Create a PrivateKey instance with the ECDSA private key
	privateKey := appmodels.NewPrivateKey(
		organization,
		appmodels.NewSerialNumber(1),
		appmodels.ECDSA_P256, // Assuming you have a KeyType for ECDSA_P256
		ecdsaPrivKey,
	)

	// Get the public key from the PrivateKey instance
	publicKey := privateKey.PublicKey()

	// Type assert the returned public key to *ecdsa.PublicKey
	ecdsaPubKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		t.Fatalf("Expected public key type *ecdsa.PublicKey, got %T", publicKey)
	}

	// Verify the public key corresponds to the private key
	if ecdsaPubKey.X.Cmp(ecdsaPrivKey.X) != 0 || ecdsaPubKey.Y.Cmp(ecdsaPrivKey.Y) != 0 {
		t.Errorf("Public key does not match private key")
	}
}
