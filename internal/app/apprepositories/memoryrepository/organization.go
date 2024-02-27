// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"fmt"
	"log"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// OrganizationRepository implements models.IOrganizationService in a memory
// @implements models.IOrganizationService
type OrganizationRepository struct {
	organizations map[string]appmodels.IOrganization
}

func (r *OrganizationRepository) FindAll() ([]appmodels.IOrganization, error) {
	list := make([]appmodels.IOrganization, 0, len(r.organizations))
	for _, org := range r.organizations {
		list = append(list, org)
	}
	return list, nil
}

func (r *OrganizationRepository) FindById(id string) (appmodels.IOrganization, error) {
	if organization, exists := r.organizations[id]; exists {
		return organization, nil
	}
	return nil, fmt.Errorf("[Organization:FindById]: not found: %s", id)
}

func (r *OrganizationRepository) Save(organization appmodels.IOrganization) (appmodels.IOrganization, error) {
	id := organization.GetID()
	r.organizations[id] = organization
	log.Printf("[Organization:Save:%s] Saved: %v", id, organization)
	return organization, nil
}

// NewOrganizationRepository creates a memory based repository for organizations
func NewOrganizationRepository() *OrganizationRepository {
	return &OrganizationRepository{
		organizations: make(map[string]appmodels.IOrganization),
	}
}

// Compile time assertion for implementing the interface
var _ appmodels.IOrganizationService = (*OrganizationRepository)(nil)
