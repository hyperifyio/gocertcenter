// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// OrganizationRepository implements models.IOrganizationService for a file system
type OrganizationRepository struct {
	filePath string
}

// NewOrganizationRepository creates a file based repository
func NewOrganizationRepository(filePath string) *OrganizationRepository {
	return &OrganizationRepository{
		filePath: filePath,
	}
}

func (r *OrganizationRepository) GetExistingOrganization(id string) (appmodels.IOrganization, error) {
	fileName := GetOrganizationJsonPath(r.filePath, id)
	dto, err := ReadOrganizationJsonFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read saved organization '%s': %w", id, err)
	}
	model := appmodels.NewOrganization(
		dto.ID,
		dto.AllNames,
	)
	return model, nil
}

func (r *OrganizationRepository) CreateOrganization(organization appmodels.IOrganization) (appmodels.IOrganization, error) {
	id := organization.GetID()
	fileName := GetOrganizationJsonPath(r.filePath, id)
	err := SaveOrganizationJsonFile(fileName, apputils.GetOrganizationDTO(organization))
	if err != nil {
		return nil, fmt.Errorf("organization creation failed: %w", err)
	}
	return r.GetExistingOrganization(id)
}

var _ appmodels.IOrganizationService = (*OrganizationRepository)(nil)
