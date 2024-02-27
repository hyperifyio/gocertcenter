// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"errors"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// PrivateKeyRepository implements models.IPrivateKeyService in a memory
// @implements models.IPrivateKeyService
type PrivateKeyRepository struct {
	keys map[string]appmodels.IPrivateKey
}

// Compile time assertion for implementing the interface
var _ appmodels.IPrivateKeyService = (*PrivateKeyRepository)(nil)

// NewPrivateKeyRepository is a memory based repository for private keys
func NewPrivateKeyRepository() *PrivateKeyRepository {
	return &PrivateKeyRepository{
		keys: make(map[string]appmodels.IPrivateKey),
	}
}

func (r *PrivateKeyRepository) FindByOrganizationAndSerialNumbers(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error) {
	if key, exists := r.keys[getCertificateLocator(organization, certificates)]; exists {
		return key, nil
	}
	return nil, errors.New("key not found")
}

func (r *PrivateKeyRepository) Save(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {
	r.keys[getCertificateLocator(key.GetOrganizationID(), append(key.GetParents(), key.GetSerialNumber()))] = key
	return key, nil
}
