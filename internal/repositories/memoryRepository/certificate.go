// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// CertificateRepository is a memory based repository for certificates
type CertificateRepository struct {
	certificates map[models.SerialNumber]*models.Certificate
}

// NewCertificateRepository creates a memory based repository for certificates
func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		certificates: make(map[models.SerialNumber]*models.Certificate),
	}
}

func (r *CertificateRepository) GetExistingCertificate(serialNumber models.SerialNumber) (*models.Certificate, error) {
	if certificate, exists := r.certificates[serialNumber]; exists {
		return certificate, nil
	}
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) CreateCertificate(certificate *models.Certificate) (*models.Certificate, error) {
	r.certificates[certificate.GetSerialNumber()] = certificate
	return certificate, nil
}
