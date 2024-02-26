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

func (m *MockCertificateManager) CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
	args := m.Called(rand, template, parent, publicKey, privateKey)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCertificateManager) ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	args := m.Called(certBytes)
	return args.Get(0).(*x509.Certificate), args.Error(1)
}

func (m *MockCertificateManager) MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	args := m.Called(key)
	return args.Get(0).([]byte)
}

func (m *MockCertificateManager) MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCertificateManager) MarshalPKCS8PrivateKey(key any) ([]byte, error) {
	args := m.Called(key)
	return args.Get(0).([]byte), args.Error(1)
}

func NewMockCertificateManager() MockCertificateManager {
	return MockCertificateManager{}
}

var _ managers.ICertificateManager = (*MockCertificateManager)(nil)
