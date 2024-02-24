// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import "github.com/hyperifyio/gocertcenter/internal/models"

// MockOrganizationService is a mock implementation of models.IOrganizationService for testing purposes.
type MockOrganizationService struct {
	GetExistingOrganizationFunc func(id string) (models.IOrganization, error)
	CreateOrganizationFunc      func(certificate models.IOrganization) (models.IOrganization, error)
}

var _ models.IOrganizationService = (*MockOrganizationService)(nil)

func (m *MockOrganizationService) GetExistingOrganization(id string) (models.IOrganization, error) {
	return m.GetExistingOrganizationFunc(id)
}

func (m *MockOrganizationService) CreateOrganization(certificate models.IOrganization) (models.IOrganization, error) {
	return m.CreateOrganizationFunc(certificate)
}
