// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// OrganizationRepository is a file based repository
type OrganizationRepository struct {
	filePath string
}

// NewOrganizationRepository creates a file based repository
func NewOrganizationRepository(filePath string) *OrganizationRepository {
	return &OrganizationRepository{
		filePath: filePath,
	}
}

func (r *OrganizationRepository) GetExistingOrganization(id string) (*models.Organization, error) {
	return nil, errors.New("organization not found")
}

func (r *OrganizationRepository) CreateOrganization(organization *models.Organization) (*models.Organization, error) {
	return nil, errors.New("organization creation not implemented")
}
