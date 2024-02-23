// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// CertificateRepository is a file based repository
type CertificateRepository struct {
	filePath string
}

// NewCertificateRepository creates a file based repository
func NewCertificateRepository(filePath string) *CertificateRepository {
	return &CertificateRepository{
		filePath: filePath,
	}
}

func (r *CertificateRepository) GetExistingCertificate(serialNumber models.SerialNumber) (*models.Certificate, error) {
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) CreateCertificate(certificate *models.Certificate) (*models.Certificate, error) {
	return nil, errors.New("certificate creation not implemented")
}
