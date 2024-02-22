// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package controllers

import (
	"github.com/hyperifyio/gocertcenter/internal/storage/repositories/memoryRepository"
)

// ControllerCollection contains all the controller instances
type ControllerCollection struct {
	Certificate *CertificateController
	Key         *KeyController
}

// NewControllerCollection returns a new ControllerCollection instance
func NewControllerCollection(repositories *memoryRepository.MemoryRepositoryCollection) *ControllerCollection {
	return &ControllerCollection{
		Certificate: NewCertificateController(repositories.CertificateRepository),
		Key:         NewKeyController(repositories.KeyRepository),
	}
}
