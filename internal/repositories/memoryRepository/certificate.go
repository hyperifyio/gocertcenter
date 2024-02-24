// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/models"
)

// CertificateRepository implements models.ICertificateService in a memory
// @implements models.ICertificateService
type CertificateRepository struct {
	certificates map[models.SerialNumber]models.ICertificate
}

// Compile time assertion for implementing the interface
var _ models.ICertificateService = (*CertificateRepository)(nil)

// NewCertificateRepository creates a memory based repository for certificates
func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		certificates: make(map[models.SerialNumber]models.ICertificate),
	}
}

func (r *CertificateRepository) GetExistingCertificate(serialNumber models.SerialNumber) (models.ICertificate, error) {
	if certificate, exists := r.certificates[serialNumber]; exists {
		return certificate, nil
	}
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	r.certificates[certificate.GetSerialNumber()] = certificate
	return certificate, nil
}
