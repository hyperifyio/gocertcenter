// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// PrivateKeyController implements appmodels.IPrivateKeyController to control
// operations for private key models.
//
// It utilizes appmodels.IPrivateKeyService interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type PrivateKeyController struct {
	repository appmodels.IPrivateKeyService
}

func (r *PrivateKeyController) GetApplicationController() appmodels.IApplicationController {
	// TODO implement me
	panic("implement me")
}

func (r *PrivateKeyController) GetOrganizationID() string {
	// TODO implement me
	panic("implement me")
}

func (r *PrivateKeyController) GetOrganizationModel() appmodels.IOrganization {
	// TODO implement me
	panic("implement me")
}

func (r *PrivateKeyController) GetOrganizationController() appmodels.IOrganizationController {
	// TODO implement me
	panic("implement me")
}

func (r *PrivateKeyController) GetCertificateModel() appmodels.ICertificate {
	// TODO implement me
	panic("implement me")
}

func (r *PrivateKeyController) GetCertificateController() appmodels.ICertificateController {
	// TODO implement me
	panic("implement me")
}

func (r *PrivateKeyController) UsesPrivateKeyService(service appmodels.IPrivateKeyService) bool {
	return r.repository == service
}

func (r *PrivateKeyController) GetExistingPrivateKey(organization string, certificates []appmodels.ISerialNumber) (appmodels.IPrivateKey, error) {
	return r.repository.GetExistingPrivateKey(organization, certificates)
}

func (r *PrivateKeyController) CreatePrivateKey(key appmodels.IPrivateKey) (appmodels.IPrivateKey, error) {
	return r.repository.CreatePrivateKey(key)
}

// NewPrivateKeyController creates a new instance of PrivateKeyController
//
//	injecting the specified IPrivateKeyService implementation. This setup
//	facilitates the separation of business logic from data access layers,
//	aligning with the principles of dependency injection.
func NewPrivateKeyController(repository appmodels.IPrivateKeyService) *PrivateKeyController {
	return &PrivateKeyController{
		repository: repository,
	}
}

var _ appmodels.IPrivateKeyController = (*PrivateKeyController)(nil)
