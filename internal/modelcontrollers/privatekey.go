// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// PrivateKeyController manages key operations.
//
//	It utilizes the IPrivateKeyService interface to abstract and inject the
//	underlying storage mechanism (e.g., database, memory). This design promotes
//	separation of concerns by decoupling the business logic from the specific
//	details of data persistence.
type PrivateKeyController struct {
	service models.IPrivateKeyService
}

// NewPrivateKeyController creates a new instance of PrivateKeyController
//
//	injecting the specified IPrivateKeyService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewPrivateKeyController(repository models.IPrivateKeyService) *PrivateKeyController {
	return &PrivateKeyController{
		service: repository,
	}
}
