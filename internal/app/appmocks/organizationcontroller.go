// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockOrganizationController is a mock implementation of models.IOrganizationController for testing purposes.
type MockOrganizationController struct {
	mock.Mock
}

func (m *MockOrganizationController) GetOrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockOrganizationController) GetOrganizationModel() appmodels.IOrganization {
	args := m.Called()
	return args.Get(0).(appmodels.IOrganization)
}

func (m *MockOrganizationController) GetApplicationController() appmodels.IApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.IApplicationController)
}

func (m *MockOrganizationController) GetCertificateController(serialNumber appmodels.ISerialNumber) (appmodels.ICertificateController, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.ICertificateController), args.Error(1)
}

func (m *MockOrganizationController) GetCertificateModel(serialNumber appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	args := m.Called(serialNumber)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) SetExpirationDuration(expiration time.Duration) {
	m.Called(expiration)
}

func (m *MockOrganizationController) NewRootCertificate(commonName string) (appmodels.ICertificate, error) {
	args := m.Called(commonName)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) UsesOrganizationService(service appmodels.IOrganizationService) bool {
	args := m.Called(service)
	return args.Bool(0)
}

func (m *MockOrganizationController) GetExistingOrganization(id string) (appmodels.IOrganization, error) {
	args := m.Called(id)
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

func (m *MockOrganizationController) CreateOrganization(organization appmodels.IOrganization) (appmodels.IOrganization, error) {
	args := m.Called(organization)
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

func (m *MockOrganizationController) NewIntermediateCertificate(o appmodels.IOrganization, manager managers.ICertificateManager, commonName string, serialNumber appmodels.ISerialNumber, parentCertificate appmodels.ICertificate, parentPrivateKey appmodels.IPrivateKey, expiration time.Duration) (appmodels.ICertificate, error) {
	args := m.Called(o, manager, commonName, serialNumber, parentCertificate, parentPrivateKey, expiration)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) NewServerCertificate(o appmodels.IOrganization, manager managers.ICertificateManager, serialNumber appmodels.ISerialNumber, parentCertificate appmodels.ICertificate, privateKey appmodels.IPrivateKey, dnsNames []string, expiration time.Duration) (appmodels.ICertificate, error) {
	args := m.Called(o, manager, serialNumber, parentCertificate, privateKey, dnsNames, expiration)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) NewClientCertificate(o appmodels.IOrganization, manager managers.ICertificateManager, commonName string, serialNumber appmodels.ISerialNumber, parentCertificate appmodels.ICertificate, privateKey appmodels.IPrivateKey, expiration time.Duration) (appmodels.ICertificate, error) {
	args := m.Called(o, manager, commonName, serialNumber, parentCertificate, privateKey, expiration)
	return args.Get(0).(appmodels.ICertificate), args.Error(1)
}

func (m *MockOrganizationController) UsesApplicationController(service appmodels.IApplicationController) bool {
	args := m.Called(service)
	return args.Bool(0)
}

var _ appmodels.IOrganizationController = (*MockOrganizationController)(nil)
