// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package fileRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/storage/models"
)

type FileCertificateRepository struct {
	filePath string
}

func NewFileCertificateRepository(filePath string) *FileCertificateRepository {
	return &FileCertificateRepository{
		filePath: filePath,
	}
}

func (r *FileCertificateRepository) GetExistingCertificate(serialNumber string) (*models.Certificate, error) {
	return nil, errors.New("certificate not found")
}

func (r *FileCertificateRepository) CreateCertificate(certificate *models.Certificate) (*models.Certificate, error) {
	return nil, errors.New("certificate creation not implemented")
}
