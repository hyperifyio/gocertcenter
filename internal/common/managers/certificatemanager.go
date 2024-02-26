// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"io"
)

// CertificateManager implements operations to manage x509 certificates by
// implementing models.ICertificateManager. This is intended to wrap low level
// external library operations for easier testing by using mocks. Any higher
// level operations shouldn't be implemented inside it.
type CertificateManager struct {
	randomManager IRandomManager
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

func (m CertificateManager) MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}

func (m CertificateManager) MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	return x509.MarshalECPrivateKey(key)
}

func (m CertificateManager) MarshalPKCS8PrivateKey(key any) ([]byte, error) {
	return x509.MarshalPKCS8PrivateKey(key)
}

func NewCertificateManager(randomManager IRandomManager) CertificateManager {
	if randomManager == nil {
		return CertificateManager{randomManager: NewRandomManager()}
	}
	return CertificateManager{randomManager}
}

var _ ICertificateManager = (*CertificateManager)(nil)
