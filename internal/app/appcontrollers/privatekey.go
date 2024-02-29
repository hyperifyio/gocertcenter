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
	model  appmodels.IPrivateKey
	parent appmodels.ICertificateController

	privateKeyRepository appmodels.IPrivateKeyService
}

func (r *PrivateKeyController) GetApplicationController() appmodels.IApplicationController {
	return r.parent.GetApplicationController()
}

func (r *PrivateKeyController) GetOrganizationID() string {
	return r.parent.GetOrganizationID()
}

func (r *PrivateKeyController) GetOrganizationModel() appmodels.IOrganization {
	return r.parent.GetOrganizationModel()
}

func (r *PrivateKeyController) GetOrganizationController() appmodels.IOrganizationController {
	if r.parent == nil {
		return nil
	}
	return r.parent.GetOrganizationController()
}

func (r *PrivateKeyController) GetCertificateModel() appmodels.ICertificate {
	return r.parent.GetCertificateModel()
}

func (r *PrivateKeyController) GetCertificateController() appmodels.ICertificateController {
	return r.parent
}

func (r *PrivateKeyController) UsesPrivateKeyService(service appmodels.IPrivateKeyService) bool {
	return r.privateKeyRepository == service
}

// NewPrivateKeyController creates a new instance of PrivateKeyController
// injecting the specified appmodels.IPrivateKeyService implementation.
//
//   - model appmodels.IPrivateKey
//   - parent appmodels.ICertificateController
//   - privateKeyRepository appmodels.IPrivateKeyService
func NewPrivateKeyController(
	model appmodels.IPrivateKey,
	parent appmodels.ICertificateController,
	privateKeyRepository appmodels.IPrivateKeyService,
) *PrivateKeyController {
	return &PrivateKeyController{
		model:                model,
		parent:               parent,
		privateKeyRepository: privateKeyRepository,
	}
}

var _ appmodels.IPrivateKeyController = (*PrivateKeyController)(nil)
