// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	models2 "github.com/hyperifyio/gocertcenter/internal/models"
)

// PrivateKeyRepository is a file based repository for private keys
type PrivateKeyRepository struct {
	filePath string
}

// NewPrivateKeyRepository creates a file based repository for private keys
func NewPrivateKeyRepository(filePath string) *PrivateKeyRepository {
	return &PrivateKeyRepository{filePath}
}

func (r *PrivateKeyRepository) GetExistingPrivateKey(
	serialNumber models2.SerialNumber,
) (*models2.PrivateKey, error) {
	return models2.NewPrivateKey(serialNumber, nil), nil
}

func (r *PrivateKeyRepository) CreatePrivateKey(
	key *models2.PrivateKey,
) (*models2.PrivateKey, error) {
	return key, nil
}
