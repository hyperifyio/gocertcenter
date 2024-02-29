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

func (m *MockOrganizationService) FindAll() ([]appmodels.IOrganization, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]appmodels.IOrganization), args.Error(1)
}

// GetExistingOrganization mocks the GetExistingOrganization method
func (m *MockOrganizationService) FindById(organization string) (appmodels.IOrganization, error) {
	args := m.Called(organization)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

// CreateOrganization mocks the CreateOrganization method
func (m *MockOrganizationService) Save(organization appmodels.IOrganization) (appmodels.IOrganization, error) {
	args := m.Called(organization)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(appmodels.IOrganization), args.Error(1)
}

// Ensure that MockOrganizationService implements IOrganizationService
var _ appmodels.IOrganizationService = (*MockOrganizationService)(nil)
