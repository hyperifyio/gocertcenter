// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"errors"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// CertificateRepository implements models.ICertificateService for a file system
type CertificateRepository struct {
	filePath string
}

var _ models.ICertificateService = (*CertificateRepository)(nil)

// NewCertificateRepository creates a file based repository
func NewCertificateRepository(filePath string) *CertificateRepository {
	return &CertificateRepository{
		filePath: filePath,
	}
}

func (r *CertificateRepository) GetExistingCertificate(
	organization string,
	certificates []models.ISerialNumber,
) (models.ICertificate, error) {

	if len(certificates) <= 0 {
		return nil, errors.New("no certificate serial numbers provided")
	}

	fileName := GetCertificatePemPath(r.filePath, organization, certificates)
	cert, err := ReadCertificateFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %w", err)
	}

	return models.NewCertificate(organization, certificates[:len(certificates)-1], cert), nil
}

func (r *CertificateRepository) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	fileName := GetCertificatePemPath(
		r.filePath,
		certificate.GetOrganizationID(),
		append(certificate.GetParents(), certificate.GetSerialNumber()),
	)
	err := SaveCertificateFile(fileName, certificate.GetCertificate())
	if err != nil {
		return nil, fmt.Errorf("failed to save certificate: %w", err)
	}
	return nil, errors.New("certificate creation not implemented")
}
