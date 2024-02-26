// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockOrganizationService is a mock implementation of appmodels.IOrganizationService
type MockOrganizationService struct {
	mock.Mock
}

// GetExistingOrganization mocks the GetExistingOrganization method
func (m *MockOrganizationService) GetExistingOrganization(organization string) (appmodels.IOrganization, error) {
	args := m.Called(organization)
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

// CreateOrganization mocks the CreateOrganization method
func (m *MockOrganizationService) CreateOrganization(organization appmodels.IOrganization) (appmodels.IOrganization, error) {
	args := m.Called(organization)
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

// Ensure that MockOrganizationService implements IOrganizationService
var _ appmodels.IOrganizationService = (*MockOrganizationService)(nil)
