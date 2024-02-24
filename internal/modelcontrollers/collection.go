// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// Collection implements collection of model controllers
type Collection struct {
	Organization models.IOrganizationController
	Certificate  models.ICertificateController
	PrivateKey   models.IPrivateKeyController
}

func NewCollection(
	organization models.IOrganizationController,
	certificate models.ICertificateController,
	privateKey models.IPrivateKeyController,
) *Collection {
	return &Collection{
		Organization: organization,
		Certificate:  certificate,
		PrivateKey:   privateKey,
	}
}
