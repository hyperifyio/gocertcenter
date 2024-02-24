// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/models"
)

// OrganizationRepository implements models.IOrganizationService in a memory
// @implements models.IOrganizationService
type OrganizationRepository struct {
	organizations map[string]models.IOrganization
}

// Compile time assertion for implementing the interface
var _ models.IOrganizationService = (*OrganizationRepository)(nil)

// NewOrganizationRepository creates a memory based repository for organizations
func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{
		organizations: make(map[string]models.IOrganization),
	}
}

func (r *OrganizationRepository) GetExistingOrganization(id string) (models.IOrganization, error) {
	if organization, exists := r.organizations[id]; exists {
		return organization, nil
	}
	return nil, errors.New("organization not found")
}

func (r *OrganizationRepository) CreateOrganization(organization models.IOrganization) (models.IOrganization, error) {
	r.organizations[organization.GetID()] = organization
	return organization, nil
}
