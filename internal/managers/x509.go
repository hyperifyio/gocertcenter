// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"io"
)

type CertificateManager struct {
	randomManager models.IRandomManager
}

func NewCertificateManager(randomManager models.IRandomManager) CertificateManager {
	if randomManager == nil {
		return CertificateManager{randomManager: NewRandomManager()}
	}
	return CertificateManager{randomManager}
}

func (m CertificateManager) GetRandomManager() models.IRandomManager {
	return m.randomManager
}

func (m CertificateManager) CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error) {
	return x509.CreateCertificate(rand, template, parent, publicKey, privateKey)
}

func (m CertificateManager) ParseCertificate(certBytes []byte) (*x509.Certificate, error) {
	return x509.ParseCertificate(certBytes)
}
