// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	"errors"

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

func (r *CertificateRepository) GetExistingCertificate(serialNumber models.SerialNumber) (models.ICertificate, error) {
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	return nil, errors.New("certificate creation not implemented")
}
