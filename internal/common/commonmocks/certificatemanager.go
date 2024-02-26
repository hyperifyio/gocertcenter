// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package commonmocks

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"io"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockCertificateManager is a mock implementation of models.ICertificateManager for testing.
type MockCertificateManager struct {
	mock.Mock
}

func (m *MockCertificateManager) GetRandomManager() managers.IRandomManager {
	args := m.Called()
	return args.Get(0).(managers.IRandomManager)
}

// CreateCertificate mocks a call to x509.CreateCertificate
//   - rand io.Reader
//   - template  *x509.Certificate
//   - parent *x509.Certificate
//   - publicKey *rsa.PublicKey, *ecdsa.PublicKey or ed25519.PublicKey
//   - privateKey crypto.Signer with a supported publicKey
//
// Returns a new certificate in DER format []byte or an error
func (m *MockCertificateManager) CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
	args := m.Called(rand, template, parent, publicKey, privateKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

// ParseCertificate mocks a call to x509.ParseCertificate to parse a single certification
//   - der []byte: ASN.1 DER data
//
// Returns *x509.Certificate or an error
func (m *MockCertificateManager) ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	args := m.Called(certBytes)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*x509.Certificate), args.Error(1)
}

// MarshalPKCS1PrivateKey wraps up a call to x509.MarshalPKCS1PrivateKey
//   - key *rsa.PrivateKey: RSA private key
//
// Returns PKCS #1, ASN.1 DER form []byte, e.g. "RSA PRIVATE KEY" PEM block or an error
func (m *MockCertificateManager) MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]byte)
}

func (m *MockCertificateManager) MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCertificateManager) MarshalPKCS8PrivateKey(key any) ([]byte, error) {
	args := m.Called(key)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func NewMockCertificateManager() MockCertificateManager {
	return MockCertificateManager{}
}

var _ managers.ICertificateManager = (*MockCertificateManager)(nil)
