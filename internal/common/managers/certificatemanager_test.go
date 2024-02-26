// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
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
