// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockCertificateController is a mock implementation of appmodels.CertificateController for testing purposes.
type MockCertificateController struct {
	mock.Mock
}

func (m *MockCertificateController) GetChildCertificateCollection(certificateType string) ([]appmodels.Certificate, error) {
	args := m.Called(certificateType)
	return args.Get(0).([]appmodels.Certificate), args.Error(2)
}

func (m *MockCertificateController) GetApplicationController() appmodels.ApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.ApplicationController)
}

func (m *MockCertificateController) GetOrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockCertificateController) GetOrganizationModel() appmodels.Organization {
	args := m.Called()
	return args.Get(0).(appmodels.Organization)
}

func (m *MockCertificateController) GetCertificateModel() appmodels.Certificate {
	args := m.Called()
	return args.Get(0).(appmodels.Certificate)
}

func (m *MockCertificateController) GetChildCertificateModel(serialNumber appmodels.SerialNumber) (appmodels.Certificate, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.Certificate), args.Error(1)
}

func (m *MockCertificateController) GetChildCertificateController(serialNumber appmodels.SerialNumber) (appmodels.CertificateController, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.CertificateController), args.Error(1)
}

func (m *MockCertificateController) GetParentCertificateModel() appmodels.Certificate {
	args := m.Called()
	return args.Get(0).(appmodels.Certificate)
}

func (m *MockCertificateController) GetParentCertificateController() appmodels.CertificateController {
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

func (m *MockCertificateController) GetOrganizationController() appmodels.OrganizationController {
	args := m.Called()
	return args.Get(0).(appmodels.OrganizationController)
}

func (m *MockCertificateController) GetPrivateKeyModel() (appmodels.PrivateKey, error) {
	args := m.Called()
	return args.Get(0).(appmodels.PrivateKey), args.Error(1)
}

func (m *MockCertificateController) GetPrivateKeyController() (appmodels.PrivateKeyController, error) {
	args := m.Called()
	return args.Get(0).(appmodels.PrivateKeyController), args.Error(1)
}

var _ appmodels.CertificateController = (*MockCertificateController)(nil)
