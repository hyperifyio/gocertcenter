// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

// Collection implements collection of model services
type Collection struct {
	Organization IOrganizationService
	Certificate  ICertificateService
	PrivateKey   IPrivateKeyService
}

func NewCollection(
	organization IOrganizationService,
	certificate ICertificateService,
	privateKey IPrivateKeyService,
) *Collection {
	return &Collection{
		Organization: organization,
		Certificate:  certificate,
		PrivateKey:   privateKey,
	}
}
