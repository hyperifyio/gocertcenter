// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockApplicationController is a mock type for the IApplicationController interface
type MockApplicationController struct {
	mock.Mock
}

func (m *MockApplicationController) GetOrganizationCollection() ([]appmodels.IOrganization, error) {
	// TODO implement me
	panic("implement me")
}

// UsesOrganizationService mocks the UsesOrganizationService method
func (m *MockApplicationController) UsesOrganizationService(service appmodels.IOrganizationService) bool {
	args := m.Called(service)
	return args.Bool(0)
}

// UsesCertificateService mocks the UsesCertificateService method
func (m *MockApplicationController) UsesCertificateService(service appmodels.ICertificateService) bool {
	args := m.Called(service)
	return args.Bool(0)
}

// UsesPrivateKeyService mocks the UsesPrivateKeyService method
func (m *MockApplicationController) UsesPrivateKeyService(service appmodels.IPrivateKeyService) bool {
	args := m.Called(service)
	return args.Bool(0)
}

// GetOrganizationModel mocks the GetOrganizationModel method
func (m *MockApplicationController) GetOrganizationModel(organization string) (appmodels.IOrganization, error) {
	args := m.Called(organization)
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

// GetOrganizationController mocks the GetOrganizationController method
func (m *MockApplicationController) GetOrganizationController(name string) (appmodels.IOrganizationController, error) {
	args := m.Called(name)
	return args.Get(0).(appmodels.IOrganizationController), args.Error(1)
}

// NewOrganization mocks the NewOrganization method
func (m *MockApplicationController) NewOrganization(certificate appmodels.IOrganization) (appmodels.IOrganization, error) {
	args := m.Called(certificate)
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

var _ appmodels.IApplicationController = (*MockApplicationController)(nil)
