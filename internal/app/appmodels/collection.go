// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

// Collection implements collection of model services
type Collection struct {
	Organization OrganizationRepository
	Certificate  CertificateRepository
	PrivateKey   PrivateKeyRepository
}

func NewCollection(
	organization OrganizationRepository,
	certificate CertificateRepository,
	privateKey PrivateKeyRepository,
) *Collection {
	return &Collection{
		Organization: organization,
		Certificate:  certificate,
		PrivateKey:   privateKey,
	}
}
