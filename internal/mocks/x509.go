// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"io"
)

// MockCertificateManager is a mock implementation of ICertificateManager for testing.
type MockCertificateManager struct {
	RandomManager         models.IRandomManager
	CreateCertificateFunc func(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error)
	ParseCertificateFunc  func(certBytes []byte) (*x509.Certificate, error)
}

func NewMockCertificateManager() MockCertificateManager {
	return MockCertificateManager{
		RandomManager: NewMockRandomManager(),
	}
}

func (m MockCertificateManager) GetRandomManager() models.IRandomManager {
	return m.RandomManager
}

// CreateCertificate calls the mocked function.
func (m *MockCertificateManager) CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
	if m.CreateCertificateFunc != nil {
		return m.CreateCertificateFunc(rand, template, parent, publicKey, privateKey)
	}
	// Return nil or some default value if not specifically mocked
	return nil, nil
}

// ParseCertificate calls the mocked function.
func (m *MockCertificateManager) ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	if m.ParseCertificateFunc != nil {
		return m.ParseCertificateFunc(certBytes)
	}
	// Return nil or some default value if not specifically mocked
	return nil, nil
}
