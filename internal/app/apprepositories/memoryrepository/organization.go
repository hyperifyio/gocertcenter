// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"fmt"
	"log"
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MemoryOrganizationRepository implements models.OrganizationRepository in a memory
// @implements models.OrganizationRepository
type MemoryOrganizationRepository struct {
	organizations map[*big.Int]appmodels.Organization
}

func (r *MemoryOrganizationRepository) FindAll() ([]appmodels.Organization, error) {
	list := make([]appmodels.Organization, 0, len(r.organizations))
	for _, org := range r.organizations {
		list = append(list, org)
	}
	return list, nil
}

func (r *MemoryOrganizationRepository) FindById(id *big.Int) (appmodels.Organization, error) {
	if organization, exists := r.organizations[id]; exists {
		return organization, nil
	}
	return nil, fmt.Errorf("[Organization:FindById]: not found: %s", id)
}

func (r *MemoryOrganizationRepository) Save(organization appmodels.Organization) (appmodels.Organization, error) {
	id := organization.ID()
	r.organizations[id] = organization
	log.Printf("[Organization:Save:%s] Saved: %v", id, organization)
	return organization, nil
}

// NewOrganizationRepository creates a memory based repository for organizations
func NewOrganizationRepository() *MemoryOrganizationRepository {
	return &MemoryOrganizationRepository{
		organizations: make(map[*big.Int]appmodels.Organization),
	}
}

// Compile time assertion for implementing the interface
var _ appmodels.OrganizationRepository = (*MemoryOrganizationRepository)(nil)
