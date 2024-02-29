// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// Collection implements collection of model controllers
type Collection struct {
	Organization appmodels.OrganizationController
	Certificate  appmodels.CertificateController
	PrivateKey   appmodels.PrivateKeyController
}

func NewCollection(
	organization appmodels.OrganizationController,
	certificate appmodels.CertificateController,
	privateKey appmodels.PrivateKeyController,
) *Collection {
	return &Collection{
		Organization: organization,
		Certificate:  certificate,
		PrivateKey:   privateKey,
	}
}
