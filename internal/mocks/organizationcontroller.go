// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import "github.com/hyperifyio/gocertcenter/internal/models"

// MockOrganizationController is a mock implementation of models.IOrganizationController for testing purposes.
type MockOrganizationController struct {
	UsesOrganizationServiceFunc func(service models.IOrganizationService) bool
	GetExistingOrganizationFunc func(id string) (models.IOrganization, error)
	CreateOrganizationFunc      func(certificate models.IOrganization) (models.IOrganization, error)
}

var _ models.IOrganizationController = (*MockOrganizationController)(nil)

func (m *MockOrganizationController) UsesOrganizationService(service models.IOrganizationService) bool {
	return m.UsesOrganizationServiceFunc(service)
}

func (m *MockOrganizationController) GetExistingOrganization(id string) (models.IOrganization, error) {
	return m.GetExistingOrganizationFunc(id)
}

func (m *MockOrganizationController) CreateOrganization(certificate models.IOrganization) (models.IOrganization, error) {
	return m.CreateOrganizationFunc(certificate)
}
