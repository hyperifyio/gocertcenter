// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package controllers

import (
	models2 "github.com/hyperifyio/gocertcenter/internal/models"
)

// IPrivateKeyService defines the interface for key storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type IPrivateKeyService interface {

	// GetExistingPrivateKey only returns public properties of the private key
	GetExistingPrivateKey(serialNumber models2.SerialNumber) (*models2.PrivateKey, error)
	CreatePrivateKey(key *models2.PrivateKey) (*models2.PrivateKey, error)
}

// PrivateKeyController manages key operations.
//
//	It utilizes the IPrivateKeyService interface to abstract and inject the
//	underlying storage mechanism (e.g., database, memory). This design promotes
//	separation of concerns by decoupling the business logic from the specific
//	details of data persistence.
type PrivateKeyController struct {
	service IPrivateKeyService
}

// NewPrivateKeyController creates a new instance of PrivateKeyController
//
//	injecting the specified IPrivateKeyService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewPrivateKeyController(repository IPrivateKeyService) *PrivateKeyController {
	return &PrivateKeyController{
		service: repository,
	}
}
