// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

import (
	"errors"
	models2 "github.com/hyperifyio/gocertcenter/internal/models"
)

// PrivateKeyRepository is a memory based repository for private keys
type PrivateKeyRepository struct {
	keys map[models2.SerialNumber]*models2.PrivateKey
}

// NewPrivateKeyRepository is a memory based repository for private keys
func NewPrivateKeyRepository() *PrivateKeyRepository {
	return &PrivateKeyRepository{
		keys: make(map[models2.SerialNumber]*models2.PrivateKey),
	}
}

func (r *PrivateKeyRepository) GetExistingPrivateKey(serialNumber models2.SerialNumber) (*models2.PrivateKey, error) {
	if key, exists := r.keys[serialNumber]; exists {
		return key, nil
	}
	return nil, errors.New("key not found")
}

func (r *PrivateKeyRepository) CreatePrivateKey(key *models2.PrivateKey) (*models2.PrivateKey, error) {
	r.keys[key.GetSerialNumber()] = key
	return key, nil
}
