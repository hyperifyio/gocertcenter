// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/models"
)

// OrganizationRepository implements models.IOrganizationService for a file system
type OrganizationRepository struct {
	filePath string
}

var _ models.IOrganizationService = (*OrganizationRepository)(nil)

// NewOrganizationRepository creates a file based repository
func NewOrganizationRepository(filePath string) *OrganizationRepository {
	return &OrganizationRepository{
		filePath: filePath,
	}
}

func (r *OrganizationRepository) GetExistingOrganization(id string) (models.IOrganization, error) {
	return nil, errors.New("organization not found")
}

func (r *OrganizationRepository) CreateOrganization(organization models.IOrganization) (models.IOrganization, error) {
	return nil, errors.New("organization creation not implemented")
}
