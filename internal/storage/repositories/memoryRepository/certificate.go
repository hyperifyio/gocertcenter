// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package memoryRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/storage/models"
)

type MemoryCertificateRepository struct {
	certificates map[string]*models.Certificate
}

func NewMemoryCertificateRepository() *MemoryCertificateRepository {
	return &MemoryCertificateRepository{
		certificates: make(map[string]*models.Certificate),
	}
}

func (r *MemoryCertificateRepository) GetExistingCertificate(serialNumber string) (*models.Certificate, error) {
	if cert, exists := r.certificates[serialNumber]; exists {
		return cert, nil
	}
	return nil, errors.New("certificate not found")
}

func (r *MemoryCertificateRepository) CreateCertificate(certificate *models.Certificate) (*models.Certificate, error) {
	r.certificates[certificate.SerialNumber] = certificate
	return certificate, nil
}
