// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package fileRepository

// Collection implements a memory based repository collection
type Collection struct {
	CertificateRepository *CertificateRepository
	KeyRepository         *PrivateKeyRepository
}

func NewCollection(filePath string) *Collection {
	return &Collection{
		CertificateRepository: NewCertificateRepository(filePath),
		KeyRepository:         NewPrivateKeyRepository(filePath),
	}
}
