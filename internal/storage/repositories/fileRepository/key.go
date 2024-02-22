// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package fileRepository

import (
	"github.com/hyperifyio/gocertcenter/internal/storage/models"
)

type FileKeyRepository struct {
	filePath string
}

func NewFileKeyRepository(filePath string) *FileKeyRepository {
	return &FileKeyRepository{filePath}
}

func (r *FileKeyRepository) GetExistingKey(serialNumber string) (*models.Key, error) {
	return models.NewKey(serialNumber), nil
}

func (r *FileKeyRepository) CreateKey(certificate *models.Key) (*models.Key, error) {
	return models.NewKey(certificate.SerialNumber), nil
}
