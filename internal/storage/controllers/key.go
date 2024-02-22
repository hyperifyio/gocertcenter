// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package controllers

import (
	"github.com/hyperifyio/gocertcenter/internal/storage/models"
	"github.com/hyperifyio/gocertcenter/internal/storage/repositories/memoryRepository"
)

// IKeyService defines the interface for key storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type IKeyService interface {
	GetExistingKey(serialNumber string) (*models.Key, error)
	CreateKey(key *models.Key) (*models.Key, error)
}

// KeyController manages key operations.
//
//	It utilizes the IKeyService interface to abstract and inject the
//	underlying storage mechanism (e.g., database, memory). This design promotes
//	separation of concerns by decoupling the business logic from the specific
//	details of data persistence.
type KeyController struct {
	service IKeyService
}

// NewKeyController creates a new instance of KeyController
//
//	injecting the specified IKeyService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewKeyController(repository *memoryRepository.MemoryKeyRepository) *KeyController {
	return &KeyController{
		service: repository,
	}
}
