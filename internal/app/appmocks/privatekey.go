// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appmocks

import (
	"crypto/x509"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// MockPrivateKey is a mock implementation of the PrivateKey interface
type MockPrivateKey struct {
	mock.Mock
}

func (m *MockPrivateKey) GetOrganizationID() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *MockPrivateKey) GetPrivateKey() any {
	args := m.Called()
	return args.Get(0).(any)
}

func (m *MockPrivateKey) GetSerialNumber() appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).(appmodels.SerialNumber)
}

func (m *MockPrivateKey) GetParents() []appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).([]appmodels.SerialNumber)
}

func (m *MockPrivateKey) GetCertificates() []appmodels.SerialNumber {
	args := m.Called()
	return args.Get(0).([]appmodels.SerialNumber)
}

func (m *MockPrivateKey) GetKeyType() appmodels.KeyType {
	args := m.Called()
	return args.Get(0).(appmodels.KeyType)
}

func (m *MockPrivateKey) GetPublicKey() any {
	args := m.Called()
	return args.Get(0)
}

func (m *MockPrivateKey) CreateCertificate(manager managers.CertificateManager, template, parent *x509.Certificate) (*x509.Certificate, error) {
	args := m.Called(manager, template, parent)
	return args.Get(0).(*x509.Certificate), args.Error(1)
}

var _ appmodels.PrivateKey = (*MockPrivateKey)(nil)
