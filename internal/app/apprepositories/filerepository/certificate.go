// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// CertificateRepository implements models.ICertificateService for a file system
type CertificateRepository struct {
	filePath string
}

// NewCertificateRepository creates a file based repository
func NewCertificateRepository(filePath string) *CertificateRepository {
	return &CertificateRepository{
		filePath: filePath,
	}
}

func (r *CertificateRepository) GetExistingCertificate(
	organization string,
	certificates []appmodels.ISerialNumber,
) (appmodels.ICertificate, error) {

	if len(certificates) <= 0 {
		return nil, errors.New("no certificate serial numbers provided")
	}

	fileName := GetCertificatePemPath(r.filePath, organization, certificates)
	cert, err := ReadCertificateFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %w", err)
	}

	return appmodels.NewCertificate(organization, certificates[:len(certificates)-1], cert), nil
}

func (r *CertificateRepository) CreateCertificate(certificate appmodels.ICertificate) (appmodels.ICertificate, error) {
	organization := certificate.GetOrganizationID()
	parents := certificate.GetParents()
	serialNumber := certificate.GetSerialNumber()
	fullPath := append(parents, serialNumber)
	fileName := GetCertificatePemPath(
		r.filePath,
		organization,
		fullPath,
	)
	err := SaveCertificateFile(fileName, certificate.GetCertificate())
	if err != nil {
		return nil, fmt.Errorf("failed to save certificate: %w", err)
	}
	return r.GetExistingCertificate(organization, fullPath)
}

var _ appmodels.ICertificateService = (*CertificateRepository)(nil)
