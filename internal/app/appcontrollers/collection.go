// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// Collection implements collection of model controllers
type Collection struct {
	Organization appmodels.IOrganizationController
	Certificate  appmodels.ICertificateController
	PrivateKey   appmodels.IPrivateKeyController
}

func NewCollection(
	organization appmodels.IOrganizationController,
	certificate appmodels.ICertificateController,
	privateKey appmodels.IPrivateKeyController,
) *Collection {
	return &Collection{
		Organization: organization,
		Certificate:  certificate,
		PrivateKey:   privateKey,
	}
}
