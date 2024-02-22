// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/x509"
	"io"
	"math/big"
)

// IRandomManager describes operations to create random values
type IRandomManager interface {
	CreateBigInt(max *big.Int) (*big.Int, error)
}

// ICertificateManager describes the operations to manage x509 certificates.
type ICertificateManager interface {
	GetRandomManager() IRandomManager
	CreateCertificate(rand io.Reader, template, parent *x509.Certificate, publicKey, privateKey any) ([]byte, error)
	ParseCertificate(certBytes []byte) (*x509.Certificate, error)
}
