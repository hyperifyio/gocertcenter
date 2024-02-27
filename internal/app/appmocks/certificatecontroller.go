// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
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

func (m *MockCertificateController) SetExpirationDuration(expiration time.Duration) {
	m.Called(expiration)
}

func (m *MockCertificateController) NewIntermediateCertificate(commonName string) (appmodels.ICertificate, appmodels.IPrivateKey, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.ICertificate), args.Get(1).(appmodels.IPrivateKey), args.Error(2)
}

func (m *MockCertificateController) NewServerCertificate(dnsNames ...string) (appmodels.ICertificate, appmodels.IPrivateKey, error) {
	args := m.Called(dnsNames)
	return args.Get(0).(appmodels.ICertificate), args.Get(1).(appmodels.IPrivateKey), args.Error(2)
}

func (m *MockCertificateController) NewClientCertificate(commonName string) (appmodels.ICertificate, appmodels.IPrivateKey, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.ICertificate), args.Get(1).(appmodels.IPrivateKey), args.Error(2)
}

func (m *MockCertificateController) GetOrganizationController() appmodels.IOrganizationController {
	args := m.Called()
	return args.Get(0).(appmodels.IOrganizationController)
}

func (m *MockCertificateController) GetPrivateKeyModel() (appmodels.IPrivateKey, error) {
	args := m.Called()
	return args.Get(0).(appmodels.IPrivateKey), args.Error(1)
}

func (m *MockCertificateController) GetPrivateKeyController() (appmodels.IPrivateKeyController, error) {
	args := m.Called()
	return args.Get(0).(appmodels.IPrivateKeyController), args.Error(1)
}

var _ appmodels.ICertificateController = (*MockCertificateController)(nil)
