// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"fmt"
	"log"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// PrivateKeyRepository implements models.IPrivateKeyService in a memory
// @implements models.IPrivateKeyService
type PrivateKeyRepository struct {
	keys map[string]appmodels.IPrivateKey
}

func (r *PrivateKeyRepository) FindByOrganizationAndSerialNumbers(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error) {
	id := getCertificateLocator(organization, certificates)
	if key, exists := r.keys[id]; exists {
		return key, nil
	}
	return nil, fmt.Errorf("[PrivateKey:FindById]: not found: %s", id)
}

func (r *PrivateKeyRepository) Save(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {
	id := getCertificateLocator(key.GetOrganizationID(), append(key.GetParents(), key.GetSerialNumber()))
	r.keys[id] = key
	log.Printf("[PrivateKey:Save:%s] Saved: %v", id, key)
	return key, nil
}

// NewPrivateKeyRepository is a memory based repository for private keys
func NewPrivateKeyRepository() *PrivateKeyRepository {
	return &PrivateKeyRepository{
		keys: make(map[string]appmodels.IPrivateKey),
	}
}

// Compile time assertion for implementing the interface
var _ appmodels.IPrivateKeyService = (*PrivateKeyRepository)(nil)
