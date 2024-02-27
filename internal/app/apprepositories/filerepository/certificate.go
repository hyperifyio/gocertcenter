// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// CertificateRepository implements models.ICertificateService for a file system
type CertificateRepository struct {
	filePath    string
	certManager managers.ICertificateManager
	fileManager managers.IFileManager
}

func (r *CertificateRepository) FindAllByOrganizationAndSerialNumbers(organization string, certificates []appmodels.ISerialNumber) ([]appmodels.ICertificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateRepository) FindAllByOrganization(organization string) ([]appmodels.ICertificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *CertificateRepository) FindByOrganizationAndSerialNumbers(
	organization string,
	certificates []appmodels.ISerialNumber,
) (appmodels.ICertificate, error) {

	if len(certificates) <= 0 {
		return nil, errors.New("no certificate serial numbers provided")
	}

	fileName := GetCertificatePemPath(r.filePath, organization, certificates)
	cert, err := ReadCertificateFile(r.fileManager, r.certManager, fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %w", err)
	}

	return appmodels.NewCertificate(organization, certificates[:len(certificates)-1], cert), nil
}

func (r *CertificateRepository) Save(certificate appmodels.ICertificate) (appmodels.ICertificate, error) {
	organization := certificate.GetOrganizationID()
	parents := certificate.GetParents()
	serialNumber := certificate.GetSerialNumber()
	fullPath := append(parents, serialNumber)
	fileName := GetCertificatePemPath(
		r.filePath,
		organization,
		fullPath,
	)
	err := SaveCertificateFile(r.fileManager, r.certManager, fileName, certificate.GetCertificate())
	if err != nil {
		return nil, fmt.Errorf("failed to save certificate: %w", err)
	}
	return r.FindByOrganizationAndSerialNumbers(organization, fullPath)
}

// NewCertificateRepository creates a file based repository
func NewCertificateRepository(
	certManager managers.ICertificateManager,
	fileManager managers.IFileManager,
	filePath string,
) *CertificateRepository {
	return &CertificateRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.ICertificateService = (*CertificateRepository)(nil)
