// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

// Collection implements memory based repository collection
type Collection struct {
	OrganizationRepository *OrganizationRepository
	CertificateRepository  *CertificateRepository
	PrivateKeyRepository   *PrivateKeyRepository
}

func NewCollection() *Collection {
	return &Collection{
		OrganizationRepository: NewOrganizationRepository(),
		CertificateRepository:  NewCertificateRepository(),
		PrivateKeyRepository:   NewPrivateKeyRepository(),
	}
}
