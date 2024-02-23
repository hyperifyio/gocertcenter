// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mocks

import "github.com/hyperifyio/gocertcenter/internal/models"

// MockOrganizationService is a mock implementation of IOrganizationService for testing purposes.
type MockOrganizationService struct {
	GetExistingOrganizationFunc func(id string) (*models.Organization, error)
	CreateOrganizationFunc      func(certificate *models.Organization) (*models.Organization, error)
}

func (m *MockOrganizationService) GetExistingOrganization(id string) (*models.Organization, error) {
	return m.GetExistingOrganizationFunc(id)
}

func (m *MockOrganizationService) CreateOrganization(certificate *models.Organization) (*models.Organization, error) {
	return m.CreateOrganizationFunc(certificate)
}
