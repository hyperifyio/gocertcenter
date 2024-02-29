// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"
	"fmt"
	"log"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MemoryCertificateRepository implements models.CertificateRepository in a memory
// @implements models.CertificateRepository
type MemoryCertificateRepository struct {
	certificates map[string]appmodels.Certificate
}

func (r *MemoryCertificateRepository) FindAllByOrganizationAndSerialNumbers(organization string, certificates []appmodels.SerialNumber) ([]appmodels.Certificate, error) {
	var result []appmodels.Certificate
	if r.certificates == nil {
		return result, nil
	}
	targetLocator := getCertificateLocator(organization, certificates)
	for _, cert := range r.certificates {
		parentLocator := getCertificateLocator(organization, cert.GetParents())
		if parentLocator == targetLocator {
			result = append(result, cert)
		}
	}
	return result, nil
}

func (r *MemoryCertificateRepository) FindAllByOrganization(organization string) ([]appmodels.Certificate, error) {
	if r.certificates == nil {
		return nil, errors.New("[CertificateModel:FindAllByOrganization]: not initialized")
	}
	var result []appmodels.Certificate
	for _, cert := range r.certificates {
		if cert.GetOrganizationID() == organization {
			result = append(result, cert)
		}
	}
	return result, nil
}

func (r *MemoryCertificateRepository) FindByOrganizationAndSerialNumbers(organization string, certificates []appmodels.SerialNumber) (appmodels.Certificate, error) {
	id := getCertificateLocator(organization, certificates)
	if certificate, exists := r.certificates[id]; exists {
		return certificate, nil
	}
	return nil, fmt.Errorf("[CertificateModel:FindByOrganizationAndSerialNumbers]: not found: %s", id)
}

func (r *MemoryCertificateRepository) Save(certificate appmodels.Certificate) (appmodels.Certificate, error) {
	id := getCertificateLocator(certificate.GetOrganizationID(), append(certificate.GetParents(), certificate.GetSerialNumber()))
	r.certificates[id] = certificate
	log.Printf("[CertificateModel:Save:%s] Saved: %v", id, certificate)
	return certificate, nil
}

// NewCertificateRepository creates a memory based repository for certificates
func NewCertificateRepository() *MemoryCertificateRepository {
	return &MemoryCertificateRepository{
		certificates: make(map[string]appmodels.Certificate),
	}
}

// Compile time assertion for implementing the interface
var _ appmodels.CertificateRepository = (*MemoryCertificateRepository)(nil)
