// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io"
)

// SystemCertificateManager implements operations to manage x509 certificates by
// implementing models.CertificateManager. This is intended to wrap low level
// external library operations for easier testing by using mocks. Any higher
// level operations shouldn't be implemented inside it.
type SystemCertificateManager struct {
	randomManager RandomManager
}

func (m SystemCertificateManager) RandomManager() RandomManager {
	return m.randomManager
}

func (m SystemCertificateManager) CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
	return x509.CreateCertificate(rand, template, parent, publicKey, privateKey)
}

func (m SystemCertificateManager) ParseCertificate(der []byte) (*x509.Certificate, error) {
	return x509.ParseCertificate(der)
}

func (m SystemCertificateManager) ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error) {
	return x509.ParseECPrivateKey(der)
}

func (m SystemCertificateManager) MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)
}

func (m SystemCertificateManager) MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error) {
	return x509.MarshalECPrivateKey(key)
}

func (m SystemCertificateManager) MarshalPKCS8PrivateKey(key any) ([]byte, error) {
	return x509.MarshalPKCS8PrivateKey(key)
}

func (m SystemCertificateManager) EncodePEMToMemory(b *pem.Block) []byte {
	return pem.EncodeToMemory(b)
}

func (m SystemCertificateManager) DecodePEM(data []byte) (p *pem.Block, rest []byte) {
	return pem.Decode(data)
}

func (m SystemCertificateManager) ParsePKCS8PrivateKey(der []byte) (any, error) {
	return x509.ParsePKCS8PrivateKey(der)
}

func (m SystemCertificateManager) ParsePKCS1PrivateKey(der []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(der)
}

func NewCertificateManager(randomManager RandomManager) SystemCertificateManager {
	if randomManager == nil {
		return SystemCertificateManager{randomManager: NewRandomManager()}
	}
	return SystemCertificateManager{randomManager}
}

var _ CertificateManager = (*SystemCertificateManager)(nil)
