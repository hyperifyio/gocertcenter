// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MemoryCertificateRepository implements models.CertificateRepository in a memory
// @implements models.CertificateRepository
type MemoryCertificateRepository struct {
	certificates map[string]appmodels.Certificate
}

func (r *MemoryCertificateRepository) FindAllByOrganizationAndSignedBy(organization string, certificate *big.Int) ([]appmodels.Certificate, error) {
	var result []appmodels.Certificate
	if r.certificates == nil {
		return result, nil
	}
	for _, cert := range r.certificates {
		if cert.OrganizationID() == organization && certificate == cert.SignedBy() {
			result = append(result, cert)
		}
	}
	return result, nil
}

func (r *MemoryCertificateRepository) FindAllByOrganization(organization string) ([]appmodels.Certificate, error) {
	if r.certificates == nil {
		return nil, errors.New("[Certificate:FindAllByOrganization]: not initialized")
	}
	var result []appmodels.Certificate
	for _, cert := range r.certificates {
		if cert.OrganizationID() == organization {
			result = append(result, cert)
		}
	}
	return result, nil
}

func (r *MemoryCertificateRepository) FindByOrganizationAndSerialNumber(organization string, certificate *big.Int) (appmodels.Certificate, error) {
	id := getCertificateLocator(organization, certificate)
	if certificate, exists := r.certificates[id]; exists {
		return certificate, nil
	}
	return nil, fmt.Errorf("[Certificate:FindByOrganizationAndSerialNumber]: not found: %s", id)
}

func (r *MemoryCertificateRepository) Save(certificate appmodels.Certificate) (appmodels.Certificate, error) {
	id := getCertificateLocator(certificate.OrganizationID(), certificate.SerialNumber())
	r.certificates[id] = certificate
	log.Printf("[Certificate:Save:%s] Saved: %v", id, certificate)
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
