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

func (r *CertPrivateKeyController) ApplicationController() appmodels.ApplicationController {
	return r.parent.ApplicationController()
}

func (r *CertPrivateKeyController) OrganizationID() string {
	return r.parent.OrganizationID()
}

func (r *CertPrivateKeyController) Organization() appmodels.Organization {
	return r.parent.Organization()
}

func (r *CertPrivateKeyController) OrganizationController() appmodels.OrganizationController {
	if r.parent == nil {
		return nil
	}
	return r.parent.OrganizationController()
}

func (r *CertPrivateKeyController) Certificate() appmodels.Certificate {
	return r.parent.Certificate()
}

func (r *CertPrivateKeyController) CertificateController() appmodels.CertificateController {
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
