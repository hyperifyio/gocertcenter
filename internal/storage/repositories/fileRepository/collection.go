// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package fileRepository

type FileRepositoryCollection struct {
	CertificateRepository *FileCertificateRepository
	KeyRepository         *FileKeyRepository
}

func NewFileRepositoryCollection(filePath string) *FileRepositoryCollection {
	return &FileRepositoryCollection{
		CertificateRepository: NewFileCertificateRepository(filePath),
		KeyRepository:         NewFileKeyRepository(filePath),
	}
}
