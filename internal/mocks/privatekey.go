// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package mocks

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"github.com/stretchr/testify/mock"
)

// MockPrivateKey is a mock implementation of the IPrivateKey interface
type MockPrivateKey struct {
	mock.Mock
}

func (m *MockPrivateKey) GetSerialNumber() models.SerialNumber {
	args := m.Called()
	return args.Get(0).(models.SerialNumber)
}

func (m *MockPrivateKey) GetKeyType() models.KeyType {
	args := m.Called()
	return args.Get(0).(models.KeyType)
}

func (m *MockPrivateKey) GetPublicKey() any {
	args := m.Called()
	return args.Get(0)
}

func (m *MockPrivateKey) CreateCertificate(manager models.ICertificateManager, template, parent *x509.Certificate) (*x509.Certificate, error) {
	args := m.Called(manager, template, parent)
	return args.Get(0).(*x509.Certificate), args.Error(1)
}
