// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// FileCertificateRepository implements models.CertificateRepository for a file system
type FileCertificateRepository struct {
	filePath    string
	certManager managers.CertificateManager
	fileManager managers.FileManager
}

func (r *FileCertificateRepository) FindAllByOrganizationAndSerialNumbers(organization string, certificates []appmodels.SerialNumber) ([]appmodels.Certificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FileCertificateRepository) FindAllByOrganization(organization string) ([]appmodels.Certificate, error) {
	// TODO implement me
	panic("implement me")
}

func (r *FileCertificateRepository) FindByOrganizationAndSerialNumbers(
	organization string,
	certificates []appmodels.SerialNumber,
) (appmodels.Certificate, error) {

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

func (r *FileCertificateRepository) Save(certificate appmodels.Certificate) (appmodels.Certificate, error) {
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
	certManager managers.CertificateManager,
	fileManager managers.FileManager,
	filePath string,
) *FileCertificateRepository {
	return &FileCertificateRepository{
		fileManager: fileManager,
		certManager: certManager,
		filePath:    filePath,
	}
}

var _ appmodels.CertificateRepository = (*FileCertificateRepository)(nil)
