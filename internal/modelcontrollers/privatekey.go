// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// PrivateKeyController implements models.IPrivateKeyController to control
// operations for private key models.
//
// It utilizes another models.IPrivateKeyService interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type PrivateKeyController struct {
	repository models.IPrivateKeyService
}

var _ models.IPrivateKeyController = (*PrivateKeyController)(nil)

func (r *PrivateKeyController) UsesPrivateKeyService(service models.IPrivateKeyService) bool {
	return r.repository == service
}

func (r *PrivateKeyController) GetExistingPrivateKey(serialNumber models.ISerialNumber) (models.IPrivateKey, error) {
	return r.repository.GetExistingPrivateKey(serialNumber)
}

func (r *PrivateKeyController) CreatePrivateKey(key models.IPrivateKey) (models.IPrivateKey, error) {
	return r.repository.CreatePrivateKey(key)
}

// NewPrivateKeyController creates a new instance of PrivateKeyController
//
//	injecting the specified IPrivateKeyService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewPrivateKeyController(repository models.IPrivateKeyService) *PrivateKeyController {
	return &PrivateKeyController{
		repository: repository,
	}
}
