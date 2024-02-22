// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package managers

import (
	"crypto/x509"
	"io"
)

// ICertificateManager describes the operations to manage x509 certificates.
type ICertificateManager interface {
	GetRandomManager() IRandomManager
	CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error)
	ParseCertificate(certBytes []byte) (*x509.Certificate, error)
}

type CertificateManager struct {
	randomManager IRandomManager
}

func NewCertificateManager(randomManager IRandomManager) CertificateManager {
	if randomManager == nil {
		return CertificateManager{randomManager: NewRandomManager()}
	}
	return CertificateManager{randomManager}
}

func (m CertificateManager) GetRandomManager() IRandomManager {
	return m.randomManager
}

func (m CertificateManager) CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
	return x509.CreateCertificate(rand, template, parent, publicKey, privateKey)
}

func (m CertificateManager) ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	return x509.ParseCertificate(certBytes)
}
