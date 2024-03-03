// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appmocks

import (
	"crypto/x509"
	"math/big"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockPrivateKey is a mock implementation of the PrivateKey interface
type MockPrivateKey struct {
	mock.Mock
}

func (m *MockPrivateKey) OrganizationID() *big.Int {
	args := m.Called()
	return args.Get(0).(*big.Int)
}

func (m *MockPrivateKey) PrivateKey() any {
	args := m.Called()
	return args.Get(0).(any)
}

func (m *MockPrivateKey) SerialNumber() *big.Int {
	args := m.Called()
	return args.Get(0).(*big.Int)
}

func (m *MockPrivateKey) Certificates() []*big.Int {
	args := m.Called()
	return args.Get(0).([]*big.Int)
}

func (m *MockPrivateKey) KeyType() appmodels.KeyType {
	args := m.Called()
	return args.Get(0).(appmodels.KeyType)
}

func (m *MockPrivateKey) PublicKey() any {
	args := m.Called()
	return args.Get(0)
}

func (m *MockPrivateKey) CreateCertificate(manager managers.CertificateManager, template, parent *x509.Certificate) (*x509.Certificate, error) {
	args := m.Called(manager, template, parent)
	return args.Get(0).(*x509.Certificate), args.Error(1)
}

var _ appmodels.PrivateKey = (*MockPrivateKey)(nil)
