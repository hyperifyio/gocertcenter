// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/models"
)

// PrivateKeyRepository implements models.IPrivateKeyService in a memory
// @implements models.IPrivateKeyService
type PrivateKeyRepository struct {
	keys map[string]models.IPrivateKey
}

// Compile time assertion for implementing the interface
var _ models.IPrivateKeyService = (*PrivateKeyRepository)(nil)

// NewPrivateKeyRepository is a memory based repository for private keys
func NewPrivateKeyRepository() *PrivateKeyRepository {
	return &PrivateKeyRepository{
		keys: make(map[string]models.IPrivateKey),
	}
}

func (r *PrivateKeyRepository) GetExistingPrivateKey(organization string, certificates []models.ISerialNumber) (models.IPrivateKey, error) {
	if key, exists := r.keys[getCertificateLocator(organization, certificates)]; exists {
		return key, nil
	}
	return nil, errors.New("key not found")
}

func (r *PrivateKeyRepository) CreatePrivateKey(key models.IPrivateKey) (models.IPrivateKey, error) {
	r.keys[getCertificateLocator(key.GetOrganizationID(), append(key.GetParents(), key.GetSerialNumber()))] = key
	return key, nil
}
