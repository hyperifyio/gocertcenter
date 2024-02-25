// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"fmt"
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
	fileName := GetOrganizationJsonPath(r.filePath, id)
	dto, err := ReadOrganizationJsonFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read saved organization '%s': %w", id, err)
	}
	model := models.NewOrganization(
		dto.ID,
		dto.AllNames,
	)
	return model, nil
}

func (r *OrganizationRepository) CreateOrganization(organization models.IOrganization) (models.IOrganization, error) {
	id := organization.GetID()
	fileName := GetOrganizationJsonPath(r.filePath, id)
	err := SaveOrganizationJsonFile(fileName, organization.GetDTO())
	if err != nil {
		return nil, fmt.Errorf("organization creation failed: %w", err)
	}
	return r.GetExistingOrganization(id)
}
