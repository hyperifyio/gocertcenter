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
	orgId string,
	signedBy models.ISerialNumber,
	serialNumber models.ISerialNumber,
) (models.ICertificate, error) {
	fileName := r.getCertificateLocation(orgId, signedBy, serialNumber)
	cert, err := readCertificateFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate: %w", err)
	}
	return models.NewCertificate(orgId, signedBy, cert), nil
}

func (r *CertificateRepository) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	fileName := r.getCertificateLocation(certificate.GetOrganizationID(), certificate.GetSignedBy(), certificate.GetSerialNumber())
	err := saveCertificateFile(fileName, certificate.GetCertificate())
	if err != nil {
		return nil, fmt.Errorf("failed to save certificate: %w", err)
	}
	return nil, errors.New("certificate creation not implemented")
}

func (r *CertificateRepository) getCertificateLocation(
	orgId string,
	signedBy models.ISerialNumber,
	serialNumber models.ISerialNumber,
) string {
	return fmt.Sprintf("%s/organizations/%s/certificates/%s/certificates/%s/cert.pem", r.filePath, orgId, signedBy.String(), serialNumber.String())
}
