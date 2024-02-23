// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"math/big"
	"testing"
	"time"
)

func TestNewCertificateManager(t *testing.T) {
	mockRandomManager := mocks.NewMockRandomManager()

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

	mockRandomManager := mocks.NewMockRandomManager()

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
