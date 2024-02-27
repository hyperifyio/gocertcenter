// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"
	"fmt"
	"log"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// CertificateRepository implements models.ICertificateService in a memory
// @implements models.ICertificateService
type CertificateRepository struct {
	certificates map[string]appmodels.ICertificate
}

func (r *CertificateRepository) FindAllByOrganizationAndSerialNumbers(organization string, certificates []appmodels.ISerialNumber) ([]appmodels.ICertificate, error) {
	targetLocator := getCertificateLocator(organization, certificates)
	var result []appmodels.ICertificate
	if r.certificates == nil {
		return result, nil
	}
	for _, cert := range r.certificates {
		parentLocator := getCertificateLocator(organization, cert.GetParents())
		if parentLocator == targetLocator {
			result = append(result, cert)
		}
	}
	return result, nil
}

func (r *CertificateRepository) FindAllByOrganization(organization string) ([]appmodels.ICertificate, error) {
	if r.certificates == nil {
		return nil, errors.New("[Certificate:FindAllByOrganization]: not initialized")
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
	id := getCertificateLocator(organization, certificates)
	if certificate, exists := r.certificates[id]; exists {
		return certificate, nil
	}
	return nil, fmt.Errorf("[Certificate:FindByOrganizationAndSerialNumbers]: not found: %s", id)
}

func (r *CertificateRepository) Save(certificate appmodels.ICertificate) (appmodels.ICertificate, error) {
	id := getCertificateLocator(certificate.GetOrganizationID(), append(certificate.GetParents(), certificate.GetSerialNumber()))
	r.certificates[id] = certificate
	log.Printf("[Certificate:Save:%s] Saved: %v", id, certificate)
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
