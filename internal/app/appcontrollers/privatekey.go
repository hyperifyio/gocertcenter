// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// CertPrivateKeyController implements appmodels.PrivateKeyController to control
// operations for private key models.
//
// It utilizes appmodels.PrivateKeyRepository interface to abstract and
// inject the underlying storage mechanism (e.g., database, memory). This design
// promotes separation of concerns by decoupling the business logic from the
// specific details of data persistence.
type CertPrivateKeyController struct {
	model  appmodels.PrivateKey
	parent appmodels.CertificateController

	privateKeyRepository appmodels.PrivateKeyRepository
}

func (r *CertPrivateKeyController) GetApplicationController() appmodels.ApplicationController {
	return r.parent.GetApplicationController()
}

func (r *CertPrivateKeyController) GetOrganizationID() string {
	return r.parent.GetOrganizationID()
}

func (r *CertPrivateKeyController) GetOrganizationModel() appmodels.Organization {
	return r.parent.GetOrganizationModel()
}

func (r *CertPrivateKeyController) GetOrganizationController() appmodels.OrganizationController {
	if r.parent == nil {
		return nil
	}
	return r.parent.GetOrganizationController()
}

func (r *CertPrivateKeyController) GetCertificateModel() appmodels.Certificate {
	return r.parent.GetCertificateModel()
}

func (r *CertPrivateKeyController) GetCertificateController() appmodels.CertificateController {
	return r.parent
}

func (r *CertPrivateKeyController) UsesPrivateKeyService(service appmodels.PrivateKeyRepository) bool {
	return r.privateKeyRepository == service
}

// NewPrivateKeyController creates a new instance of CertPrivateKeyController
// injecting the specified appmodels.PrivateKeyRepository implementation.
//
//   - model appmodels.PrivateKey
//   - parent appmodels.CertificateController
//   - privateKeyRepository appmodels.PrivateKeyRepository
func NewPrivateKeyController(
	model appmodels.PrivateKey,
	parent appmodels.CertificateController,
	privateKeyRepository appmodels.PrivateKeyRepository,
) *CertPrivateKeyController {
	return &CertPrivateKeyController{
		model:                model,
		parent:               parent,
		privateKeyRepository: privateKeyRepository,
	}
}

var _ appmodels.PrivateKeyController = (*CertPrivateKeyController)(nil)
