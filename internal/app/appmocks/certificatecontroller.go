// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"crypto/x509"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificateController is a mock implementation of appmodels.ICertificateController for testing purposes.
type MockCertificateController struct {
	mock.Mock
}

func (m *MockCertificateController) GetApplicationController() appmodels.IApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.IApplicationController)
}

func (m *MockCertificateController) GetOrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificateController) GetOrganizationModel() appmodels.IOrganization {
	args := m.Called()
	return args.Get(0).(appmodels.IOrganization)
}

func (m *MockCertificateController) GetCertificateModel() appmodels.ICertificate {
	args := m.Called()
	return args.Get(0).(appmodels.ICertificate)
}

func (m *MockCertificateController) GetChildCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockCertificateController) GetChildCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.ICertificateController), args.Error(1)
}

func (m *MockCertificateController) GetParentCertificateModel() appmodels.ICertificate {
	args := m.Called()
	return args.Get(0).(appmodels.ICertificate)
}

func (m *MockCertificateController) GetParentCertificateController() appmodels.ICertificateController {
	args := m.Called()
	return args.Get(0).(appmodels.ICertificateController)
}

func (m *MockCertificateController) GetPrivateKeyModel() appmodels.IPrivateKey {
	args := m.Called()
	return args.Get(0).(appmodels.IPrivateKey)
}

func (m *MockCertificateController) GetPrivateKeyController() appmodels.IPrivateKeyController {
	args := m.Called()
	return args.Get(0).(appmodels.IPrivateKeyController)
}

func (m *MockCertificateController) NewCertificate(template *x509.Certificate) (appmodels.ICertificate, error) {
	args := m.Called(template)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockCertificateController) SetExpirationDuration(expiration time.Duration) {
	m.Called(expiration)
}

func (m *MockCertificateController) NewIntermediateCertificate(commonName string) (appmodels.ICertificate, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockCertificateController) NewServerCertificate(dnsNames ...string) (appmodels.ICertificate, error) {
	args := m.Called(dnsNames)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockCertificateController) NewClientCertificate(commonName string) (appmodels.ICertificate, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

var _ appmodels.ICertificateController = (*MockCertificateController)(nil)
