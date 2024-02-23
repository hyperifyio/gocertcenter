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

// IOrganizationService defines the interface for organization storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type IOrganizationService interface {
	GetExistingOrganization(id string) (*Organization, error)
	CreateOrganization(certificate *Organization) (*Organization, error)
}

// ICertificateService defines the interface for certificate storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type ICertificateService interface {
	GetExistingCertificate(serialNumber SerialNumber) (*Certificate, error)
	CreateCertificate(certificate *Certificate) (*Certificate, error)
}

// IPrivateKeyService defines the interface for key storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type IPrivateKeyService interface {

	// GetExistingPrivateKey only returns public properties of the private key
	GetExistingPrivateKey(serialNumber SerialNumber) (*PrivateKey, error)
	CreatePrivateKey(key *PrivateKey) (*PrivateKey, error)
}
