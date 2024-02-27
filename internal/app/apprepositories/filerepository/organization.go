// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// OrganizationRepository implements models.IOrganizationService for a file system
type OrganizationRepository struct {
	filePath    string
	certManager managers.ICertificateManager
	fileManager managers.IFileManager
}

func (r *OrganizationRepository) FindAll() ([]appmodels.IOrganization, error) {
	// FIXME: Not implemented yet
	var list []appmodels.IOrganization
	return list, nil
}

func (r *OrganizationRepository) FindById(id string) (appmodels.IOrganization, error) {
	fileName := GetOrganizationJsonPath(r.filePath, id)
	dto, err := ReadOrganizationJsonFile(r.fileManager, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read saved organization '%s': %w", id, err)
	}
	model := appmodels.NewOrganization(
		dto.ID,
		dto.AllNames,
	)
	return model, nil
}

func (r *OrganizationRepository) Save(organization appmodels.IOrganization) (appmodels.IOrganization, error) {
	id := organization.GetID()
	fileName := GetOrganizationJsonPath(r.filePath, id)
	err := SaveOrganizationJsonFile(r.fileManager, fileName, apputils.GetOrganizationDTO(organization))
	if err != nil {
		return nil, fmt.Errorf("organization creation failed: %w", err)
	}
	return r.FindById(id)
}

// NewOrganizationRepository creates a file based repository
func NewOrganizationRepository(
	certManager managers.ICertificateManager,
	fileManager managers.IFileManager,
	filePath string,
) *OrganizationRepository {
	return &OrganizationRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.IOrganizationService = (*OrganizationRepository)(nil)
