// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// PrivateKeyRepository implements models.IPrivateKeyService for a file system
type PrivateKeyRepository struct {
	filePath string
}

var _ models.IPrivateKeyService = (*PrivateKeyRepository)(nil)

// NewPrivateKeyRepository creates a file based repository for private keys
func NewPrivateKeyRepository(filePath string) *PrivateKeyRepository {
	return &PrivateKeyRepository{filePath}
}

func (r *PrivateKeyRepository) GetFilePath() string {
	return r.filePath
}

func (r *PrivateKeyRepository) GetExistingPrivateKey(
	serialNumber models.SerialNumber,
) (models.IPrivateKey, error) {
	return models.NewPrivateKey(serialNumber, models.ECDSA_P384, nil), nil
}

func (r *PrivateKeyRepository) CreatePrivateKey(
	key models.IPrivateKey,
) (models.IPrivateKey, error) {
	return key, nil
}
