// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
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

func TestToPrivateKeyDTO(t *testing.T) {
	mockCertManager := new(commonmocks.MockCertificateManager)
	mockPrivateKey := new(appmocks.MockPrivateKey)
	mockSerialNumber := new(appmocks.MockSerialNumber)

	// Setup expected values
	expectedSerial := "123456789"
	expectedKeyType := appmodels.RSA_2048.String()
	expectedPEM := "FAKE_PEM_DATA"

	mockSerialNumber.On("String").Return(expectedSerial)
	mockPrivateKey.On("GetSerialNumber").Return(mockSerialNumber)
	mockPrivateKey.On("GetKeyType").Return(appmodels.RSA_2048)
	mockPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})
	mockCertManager.On("MarshalPKCS1PrivateKey", mock.Anything).Return([]byte(expectedPEM))
	mockCertManager.On("EncodePEMToMemory", mock.Anything).Return([]byte(expectedPEM))

	dto, err := apputils.ToPrivateKeyDTO(mockCertManager, mockPrivateKey)
	if err != nil {
		t.Fatalf("ToPrivateKeyDTO() unexpected error = %v", err)
	}

	// Validate results
	if dto.Certificate != expectedSerial || dto.Type != expectedKeyType || dto.PrivateKey != expectedPEM {
		t.Errorf("ToPrivateKeyDTO() got = '%v', want '%v'", dto, appdtos.PrivateKeyDTO{Certificate: expectedSerial, Type: expectedKeyType, PrivateKey: expectedPEM})
	}
}

func TestToPrivateKeyDTOList(t *testing.T) {
	mockCertManager := new(commonmocks.MockCertificateManager)
	mockPrivateKey := new(appmocks.MockPrivateKey)
	mockSerialNumber := new(appmocks.MockSerialNumber)

	// Setup expected values for a list of one key for simplicity
	expectedSerial := "123456789"
	expectedKeyType := appmodels.RSA_2048.String()
	expectedPEM := "FAKE_PEM_DATA"

	mockSerialNumber.On("String").Return(expectedSerial)
	mockPrivateKey.On("GetSerialNumber").Return(mockSerialNumber)
	mockPrivateKey.On("GetKeyType").Return(appmodels.RSA_2048)
	mockPrivateKey.On("GetPrivateKey").Return(&rsa.PrivateKey{})
	mockCertManager.On("MarshalPKCS1PrivateKey", mock.Anything).Return([]byte(expectedPEM))
	mockCertManager.On("EncodePEMToMemory", mock.Anything).Return([]byte(expectedPEM))

	list, err := apputils.ToPrivateKeyDTOList(mockCertManager, []appmodels.IPrivateKey{mockPrivateKey})
	if err != nil {
		t.Fatalf("ToPrivateKeyDTOList() error = '%v', wantErr '%v'", err, false)
	}

	if len(list) != 1 || list[0].Certificate != expectedSerial || list[0].Type != expectedKeyType || list[0].PrivateKey != expectedPEM {
		t.Errorf("ToPrivateKeyDTOList() got = '%v', want '%v'", list, []appdtos.PrivateKeyDTO{{Certificate: expectedSerial, Type: expectedKeyType, PrivateKey: expectedPEM}})
	}
}

func TestDetermineRSATypeFromSize(t *testing.T) {
	tests := []struct {
		name     string
		keySize  int
		expected appmodels.KeyType
		wantErr  bool
	}{
		{"RSA_1024", 1024, appmodels.RSA_1024, false},
		{"RSA_2048", 2048, appmodels.RSA_2048, false},
		{"RSA_3072", 3072, appmodels.RSA_3072, false},
		{"RSA_4096", 4096, appmodels.RSA_4096, false},
		{"Invalid", 512, appmodels.NIL_KEY_TYPE, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := apputils.DetermineRSATypeFromSize(tt.keySize)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetermineRSATypeFromSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("DetermineRSATypeFromSize() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDetermineECDSACurve(t *testing.T) {
	tests := []struct {
		name     string
		curve    elliptic.Curve
		expected appmodels.KeyType
		wantErr  bool
	}{
		{"P224", elliptic.P224(), appmodels.ECDSA_P224, false},
		{"P256", elliptic.P256(), appmodels.ECDSA_P256, false},
		{"P384", elliptic.P384(), appmodels.ECDSA_P384, false},
		{"P521", elliptic.P521(), appmodels.ECDSA_P521, false},
		{"Invalid", nil, appmodels.NIL_KEY_TYPE, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := apputils.DetermineECDSACurve(tt.curve)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetermineECDSACurve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("DetermineECDSACurve() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestReadRSAKeySize(t *testing.T) {
	tests := []struct {
		name    string
		keySize int // RSA key sizes to test
	}{
		{"RSA_1024", 1024},
		{"RSA_2048", 2048},
		{"RSA_3072", 3072},
		{"RSA_4096", 4096},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Generate RSA key of specified size
			key, err := rsa.GenerateKey(rand.Reader, tt.keySize)
			if err != nil {
				t.Fatalf("Failed to generate RSA key of size %d: %v", tt.keySize, err)
			}

			// Call ReadRSAKeySize
			size := apputils.ReadRSAKeySize(key)

			// Assert key size
			if size != tt.keySize {
				t.Errorf("ReadRSAKeySize() returned %d, want %d", size, tt.keySize)
			}
		})
	}
}

func TestDetermineRSATypeFromKey(t *testing.T) {
	tests := []struct {
		name    string
		keySize int
		want    appmodels.KeyType
	}{
		{"RSA_1024", 1024, appmodels.RSA_1024},
		{"RSA_2048", 2048, appmodels.RSA_2048},
		{"RSA_3072", 3072, appmodels.RSA_3072},
		{"RSA_4096", 4096, appmodels.RSA_4096},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsaKey, err := rsa.GenerateKey(rand.Reader, tt.keySize)
			if err != nil {
				t.Fatalf("Failed to generate RSA key of size %d: %v", tt.keySize, err)
			}

			got, err := apputils.DetermineRSATypeFromKey(rsaKey)
			if err != nil {
				t.Fatalf("DetermineRSATypeFromKey() error = %v, wantErr false", err)
			}
			if got != tt.want {
				t.Errorf("DetermineRSATypeFromKey() got = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("NonRSAType", func(t *testing.T) {
		ecdsaKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Fatalf("Failed to generate ECDSA key: %v", err)
		}

		_, err = apputils.DetermineRSATypeFromKey(ecdsaKey)
		if err == nil {
			t.Errorf("DetermineRSATypeFromKey() expected error, got nil")
		}
	})
}

func TestDetermineKeyType(t *testing.T) {
	// Test RSA keys
	rsaKeySizes := map[int]appmodels.KeyType{
		1024: appmodels.RSA_1024,
		2048: appmodels.RSA_2048,
		3072: appmodels.RSA_3072,
		4096: appmodels.RSA_4096,
	}
	for size, expected := range rsaKeySizes {
		rsaKey, err := rsa.GenerateKey(rand.Reader, size)
		if err != nil {
			t.Fatalf("Failed to generate RSA key of size %d: %v", size, err)
		}
		got, err := apputils.DetermineKeyType(rsaKey)
		if err != nil {
			t.Errorf("DetermineKeyType() with RSA key size %d error = %v", size, err)
			continue
		}
		if got != expected {
			t.Errorf("DetermineKeyType() with RSA key size %d got = %v, want %v", size, got, expected)
		}
	}

	// Test ECDSA keys
	ecdsaCurves := map[elliptic.Curve]appmodels.KeyType{
		elliptic.P224(): appmodels.ECDSA_P224,
		elliptic.P256(): appmodels.ECDSA_P256,
		elliptic.P384(): appmodels.ECDSA_P384,
		elliptic.P521(): appmodels.ECDSA_P521,
	}
	for curve, expected := range ecdsaCurves {
		ecdsaKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			t.Fatalf("Failed to generate ECDSA key for curve %v: %v", curve, err)
		}
		got, err := apputils.DetermineKeyType(ecdsaKey)
		if err != nil {
			t.Errorf("DetermineKeyType() with ECDSA curve %v error = %v", curve, err)
			continue
		}
		if got != expected {
			t.Errorf("DetermineKeyType() with ECDSA curve %v got = %v, want %v", curve, got, expected)
		}
	}

	// Test Ed25519 key
	_, ed25519Key, _ := ed25519.GenerateKey(rand.Reader)
	if got, err := apputils.DetermineKeyType(ed25519Key); err != nil || got != appmodels.Ed25519 {
		t.Errorf("DetermineKeyType() with Ed25519 key got = %v, err = %v; want %v", got, err, appmodels.Ed25519)
	}

	// Test unsupported key type
	unsupportedKey := "this is not a key"
	if _, err := apputils.DetermineKeyType(unsupportedKey); err == nil {
		t.Errorf("DetermineKeyType() with unsupported key type did not return an error")
	}
}

func TestParsePrivateKeyFromPEMBytes(t *testing.T) {
	tests := []struct {
		name    string
		keyType appmodels.KeyType
		setup   func(certManager *commonmocks.MockCertificateManager) ([]byte, error)
	}{
		{
			name:    "RSA",
			keyType: appmodels.RSA_2048,
			setup: func(certManager *commonmocks.MockCertificateManager) ([]byte, error) {

				key, err := rsa.GenerateKey(rand.Reader, 2048)
				if err != nil {
					return nil, err
				}
				der := x509.MarshalPKCS1PrivateKey(key)
				if der == nil {
					return nil, fmt.Errorf("failed to marshal RSA to DER")
				}
				pemBlock := &pem.Block{
					Type:  "RSA PRIVATE KEY",
					Bytes: der,
				}

				pemBytes := pem.EncodeToMemory(pemBlock)

				certManager.On("DecodePEM", pemBytes).Return(pemBlock, []byte(nil))
				certManager.On("ParsePKCS1PrivateKey", der).Return(key, nil)

				return pemBytes, nil

			},
		},
		{
			name:    "ECDSA",
			keyType: appmodels.ECDSA_P256,
			setup: func(certManager *commonmocks.MockCertificateManager) ([]byte, error) {
				key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
				if err != nil {
					return nil, err
				}
				der, err := x509.MarshalECPrivateKey(key)
				if err != nil {
					return nil, err
				}
				if der == nil {
					return nil, fmt.Errorf("failed to marshal ECDSA to DER")
				}
				pemBlock := &pem.Block{
					Type:  "EC PRIVATE KEY",
					Bytes: der,
				}

				pemBytes := pem.EncodeToMemory(pemBlock)
				certManager.On("DecodePEM", pemBytes).Return(pemBlock, []byte(nil))
				certManager.On("ParseECPrivateKey", der).Return(key, nil)

				return pemBytes, nil
			},
		},
		{
			name:    "Ed25519",
			keyType: appmodels.Ed25519,
			setup: func(certManager *commonmocks.MockCertificateManager) ([]byte, error) {
				_, key, err := ed25519.GenerateKey(rand.Reader)
				if err != nil {
					return nil, err
				}
				assert.NoError(t, err)
				der, err := x509.MarshalPKCS8PrivateKey(key)
				if err != nil {
					return nil, err
				}
				if der == nil {
					return nil, fmt.Errorf("failed to marshal Ed25519 to DER")
				}
				pemBlock := &pem.Block{
					Type:  "PRIVATE KEY",
					Bytes: der,
				}
				pemBytes := pem.EncodeToMemory(pemBlock)

				certManager.On("DecodePEM", pemBytes).Return(pemBlock, []byte(nil))
				certManager.On("ParsePKCS8PrivateKey", der).Return(key, nil)

				return pemBytes, nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCertManager := &commonmocks.MockCertificateManager{}
			pemBytes, err := tt.setup(mockCertManager)
			assert.NoError(t, err)

			// Call the function under test
			gotKey, gotKeyType, gotErr := apputils.ParsePrivateKeyFromPEMBytes(mockCertManager, pemBytes)

			assert.NoError(t, gotErr)
			assert.NotNil(t, gotKey)
			assert.Equal(t, tt.keyType, gotKeyType)
		})
	}

	// Test with invalid PEM data
	t.Run("Invalid PEM", func(t *testing.T) {
		mockCertManager := &commonmocks.MockCertificateManager{}
		mockCertManager.On("DecodePEM", mock.Anything).Return(nil, []byte(nil))

		_, _, err := apputils.ParsePrivateKeyFromPEMBytes(mockCertManager, []byte("invalid pem data"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to decode PEM block containing the private key")
	})
}
