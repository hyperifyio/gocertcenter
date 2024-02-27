// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"time"

	"github.com/stretchr/testify/mock"

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

func (m *MockOrganizationController) UsesApplicationController(service appmodels.IApplicationController) bool {
	args := m.Called(service)
	return args.Bool(0)
}

var _ appmodels.IOrganizationController = (*MockOrganizationController)(nil)
