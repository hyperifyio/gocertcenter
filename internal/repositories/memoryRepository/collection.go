// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryRepository

// Collection implements memory based repository collection
type Collection struct {
	CertificateRepository *CertificateRepository
	PrivateKeyRepository  *PrivateKeyRepository
}

func NewCollection() *Collection {
	return &Collection{
		CertificateRepository: NewCertificateRepository(),
		PrivateKeyRepository:  NewPrivateKeyRepository(),
	}
}
