// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"fmt"
	"log"
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MemoryPrivateKeyRepository implements models.PrivateKeyRepository in a memory
// @implements models.PrivateKeyRepository
type MemoryPrivateKeyRepository struct {
	keys map[string]appmodels.PrivateKey
}

func (r *MemoryPrivateKeyRepository) FindByOrganizationAndSerialNumbers(organization string, certificates []*big.Int) (appmodels.PrivateKey, error) {
	id := getCertificateLocator(organization, certificates)
	if key, exists := r.keys[id]; exists {
		return key, nil
	}
	return nil, fmt.Errorf("[PrivateKey:FindById]: not found: %s", id)
}

func (r *MemoryPrivateKeyRepository) Save(key appmodels.PrivateKey) (appmodels.PrivateKey, error) {
	id := getCertificateLocator(key.OrganizationID(), append(key.Parents(), key.SerialNumber()))
	r.keys[id] = key
	log.Printf("[PrivateKey:Save:%s] Saved: %v", id, key)
	return key, nil
}

// NewPrivateKeyRepository is a memory based repository for private keys
func NewPrivateKeyRepository() *MemoryPrivateKeyRepository {
	return &MemoryPrivateKeyRepository{
		keys: make(map[string]appmodels.PrivateKey),
	}
}

// Compile time assertion for implementing the interface
var _ appmodels.PrivateKeyRepository = (*MemoryPrivateKeyRepository)(nil)
