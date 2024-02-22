// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package memoryRepository

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/storage/models"
)

type MemoryKeyRepository struct {
	keys map[string]*models.Key
}

func NewMemoryKeyRepository() *MemoryKeyRepository {
	return &MemoryKeyRepository{
		keys: make(map[string]*models.Key),
	}
}

func (r *MemoryKeyRepository) GetExistingKey(serialNumber string) (*models.Key, error) {
	if cert, exists := r.keys[serialNumber]; exists {
		return cert, nil
	}
	return nil, errors.New("key not found")
}

func (r *MemoryKeyRepository) CreateKey(key *models.Key) (*models.Key, error) {
	r.keys[key.SerialNumber] = key
	return key, nil
}
