// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

import (
	"errors"
	models2 "github.com/hyperifyio/gocertcenter/internal/models"
)

// CertificateRepository is a memory based repository for certificates
type CertificateRepository struct {
	certificates map[models2.SerialNumber]*models2.Certificate
}

// NewCertificateRepository creates a memory based repository for certificates
func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		certificates: make(map[models2.SerialNumber]*models2.Certificate),
	}
}

func (r *CertificateRepository) GetExistingCertificate(serialNumber models2.SerialNumber) (*models2.Certificate, error) {
	if certificate, exists := r.certificates[serialNumber]; exists {
		return certificate, nil
	}
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) CreateCertificate(certificate *models2.Certificate) (*models2.Certificate, error) {
	r.certificates[certificate.GetSerialNumber()] = certificate
	return certificate, nil
}
