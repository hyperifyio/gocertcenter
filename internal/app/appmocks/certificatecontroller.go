// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"math/big"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificateController is a mock implementation of appmodels.CertificateController for testing purposes.
type MockCertificateController struct {
	mock.Mock
}

func (m *MockCertificateController) ChildCertificateCollection(certificateType string) ([]appmodels.Certificate, error) {
	args := m.Called(certificateType)
	return args.Get(0).([]appmodels.Certificate), args.Error(2)
}

func (m *MockCertificateController) ApplicationController() appmodels.ApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.ApplicationController)
}

func (m *MockCertificateController) OrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificateController) Organization() appmodels.Organization {
	args := m.Called()
	return args.Get(0).(appmodels.Organization)
}

func (m *MockCertificateController) Certificate() appmodels.Certificate {
	args := m.Called()
	return args.Get(0).(appmodels.Certificate)
}

func (m *MockCertificateController) ChildCertificate(serialNumber *big.Int) (appmodels.Certificate, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateController) ChildCertificateController(serialNumber *big.Int) (appmodels.CertificateController, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.CertificateController), args.Error(1)
}

func (m *MockCertificateController) ParentCertificate() appmodels.Certificate {
	args := m.Called()
	return args.Get(0).(appmodels.Certificate)
}

func (m *MockCertificateController) ParentCertificateController() appmodels.CertificateController {
	args := m.Called()
	return args.Get(0).(appmodels.CertificateController)
}

func (m *MockCertificateController) SetExpirationDuration(expiration time.Duration) {
	m.Called(expiration)
}

func (m *MockCertificateController) NewIntermediateCertificate(commonName string) (appmodels.Certificate, appmodels.PrivateKey, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.Certificate), args.Get(1).(appmodels.PrivateKey), args.Error(2)
}

func (m *MockCertificateController) NewServerCertificate(dnsNames ...string) (appmodels.Certificate, appmodels.PrivateKey, error) {
	args := m.Called(dnsNames)
	return args.Get(0).(appmodels.Certificate), args.Get(1).(appmodels.PrivateKey), args.Error(2)
}

func (m *MockCertificateController) NewClientCertificate(commonName string) (appmodels.Certificate, appmodels.PrivateKey, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.Certificate), args.Get(1).(appmodels.PrivateKey), args.Error(2)
}

func (m *MockCertificateController) OrganizationController() appmodels.OrganizationController {
	args := m.Called()
	return args.Get(0).(appmodels.OrganizationController)
}

func (m *MockCertificateController) PrivateKey() (appmodels.PrivateKey, error) {
	args := m.Called()
	return args.Get(0).(appmodels.PrivateKey), args.Error(1)
}

func (m *MockCertificateController) PrivateKeyController() (appmodels.PrivateKeyController, error) {
	args := m.Called()
	return args.Get(0).(appmodels.PrivateKeyController), args.Error(1)
}

var _ appmodels.CertificateController = (*MockCertificateController)(nil)
