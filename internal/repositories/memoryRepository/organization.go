// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// OrganizationRepository is a memory based repository for organizations
type OrganizationRepository struct {
	organizations map[string]*models.Organization
}

// NewOrganizationRepository creates a memory based repository for organizations
func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{
		organizations: make(map[string]*models.Organization),
	}
}

func (r *OrganizationRepository) GetExistingOrganization(id string) (*models.Organization, error) {
	if organization, exists := r.organizations[id]; exists {
		return organization, nil
	}
	return nil, errors.New("organization not found")
}

func (r *OrganizationRepository) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	r.organizations[organization.GetID()] = organization
	return organization, nil
}
