// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/models"
)

// CertificateRepository implements models.ICertificateService in a memory
// @implements models.ICertificateService
type CertificateRepository struct {
	certificates map[models.ISerialNumber]models.ICertificate
}

// Compile time assertion for implementing the interface
var _ models.ICertificateService = (*CertificateRepository)(nil)

// NewCertificateRepository creates a memory based repository for certificates
func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		certificates: make(map[models.ISerialNumber]models.ICertificate),
	}
}

func (r *CertificateRepository) GetExistingCertificate(
	orgId string,
	signedBy models.ISerialNumber,
	serialNumber models.ISerialNumber,
) (models.ICertificate, error) {
	if certificate, exists := r.certificates[serialNumber]; exists {
		return certificate, nil
	}
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) CreateCertificate(certificate models.ICertificate) (models.ICertificate, error) {
	r.certificates[certificate.GetSerialNumber()] = certificate
	return certificate, nil
}
