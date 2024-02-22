// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package memoryRepository

type MemoryRepositoryCollection struct {
	CertificateRepository *MemoryCertificateRepository
	KeyRepository         *MemoryKeyRepository
}

func NewMemoryRepositoryCollection() *MemoryRepositoryCollection {
	return &MemoryRepositoryCollection{
		CertificateRepository: NewMemoryCertificateRepository(),
		KeyRepository:         NewMemoryKeyRepository(),
	}
}
