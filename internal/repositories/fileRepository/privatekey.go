// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// PrivateKeyRepository is a file based repository for private keys
type PrivateKeyRepository struct {
	filePath string
}

// NewPrivateKeyRepository creates a file based repository for private keys
func NewPrivateKeyRepository(filePath string) *PrivateKeyRepository {
	return &PrivateKeyRepository{filePath}
}

func (r *PrivateKeyRepository) GetFilePath() string {
	return r.filePath
}

func (r *PrivateKeyRepository) GetExistingPrivateKey(
	serialNumber models.SerialNumber,
) (*models.PrivateKey, error) {
	return models.NewPrivateKey(serialNumber, models.ECDSA_P384, nil), nil
}

func (r *PrivateKeyRepository) CreatePrivateKey(
	key *models.PrivateKey,
) (*models.PrivateKey, error) {
	return key, nil
}
