// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// FileOrganizationRepository implements models.OrganizationRepository for a file system
type FileOrganizationRepository struct {
	filePath    string
	certManager managers.CertificateManager
	fileManager managers.FileManager
}

func (r *FileOrganizationRepository) FindAll() ([]appmodels.Organization, error) {
	// FIXME: Not implemented yet
	var list []appmodels.Organization
	return list, nil
}

func (r *FileOrganizationRepository) FindById(id string) (appmodels.Organization, error) {
	fileName := OrganizationJsonPath(r.filePath, id)
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

func (r *FileOrganizationRepository) Save(organization appmodels.Organization) (appmodels.Organization, error) {
	id := organization.ID()
	fileName := OrganizationJsonPath(r.filePath, id)
	err := SaveOrganizationJsonFile(r.fileManager, fileName, apputils.ToOrganizationDTO(organization))
	if err != nil {
		return nil, fmt.Errorf("organization creation failed: %w", err)
	}
	return r.FindById(id)
}

// NewOrganizationRepository creates a file based repository
func NewOrganizationRepository(
	certManager managers.CertificateManager,
	fileManager managers.FileManager,
	filePath string,
) *FileOrganizationRepository {
	return &FileOrganizationRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.OrganizationRepository = (*FileOrganizationRepository)(nil)
