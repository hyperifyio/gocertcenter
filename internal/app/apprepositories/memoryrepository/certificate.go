// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// CertificateRepository implements models.ICertificateService in a memory
// @implements models.ICertificateService
type CertificateRepository struct {
	certificates map[string]appmodels.ICertificate
}

func (r *CertificateRepository) FindAllByOrganization(organization string) ([]appmodels.ICertificate, error) {
	if r.certificates == nil {
		return nil, errors.New("certificate repository is not initialized")
	}

	var result []appmodels.ICertificate
	for _, cert := range r.certificates {
		if cert.GetOrganizationID() == organization {
			result = append(result, cert)
		}
	}
	return result, nil
}

func (r *CertificateRepository) FindByOrganizationAndSerialNumbers(organization string, certificates []appmodels.ISerialNumber) (appmodels.ICertificate, error) {
	if certificate, exists := r.certificates[getCertificateLocator(organization, certificates)]; exists {
		return certificate, nil
	}
	return nil, errors.New("certificate not found")
}

func (r *CertificateRepository) Save(certificate appmodels.ICertificate) (appmodels.ICertificate, error) {
	r.certificates[getCertificateLocator(certificate.GetOrganizationID(), append(certificate.GetParents(), certificate.GetSerialNumber()))] = certificate
	return certificate, nil
}

// NewCertificateRepository creates a memory based repository for certificates
func NewCertificateRepository() *CertificateRepository {
	return &CertificateRepository{
		certificates: make(map[string]appmodels.ICertificate),
	}
}

// Compile time assertion for implementing the interface
var _ appmodels.ICertificateService = (*CertificateRepository)(nil)
