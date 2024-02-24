// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/x509"
	"io"
	"math/big"
	"time"
)

// IOrganization describes an interface for Organization model
type IOrganization interface {

	// GetID returns unique identifier for this organization
	GetID() string

	// GetName returns the primary organization name
	GetName() string

	// GetNames returns the full name of the organization including department
	GetNames() []string

	// NewRootCertificate creates a new root CA certificate
	NewRootCertificate(
		manager ICertificateManager,
		commonName string, // The name of the root CA
		privateKey IPrivateKey, // Private key for self signing
		expiration time.Duration,
	) (ICertificate, error)

	// NewIntermediateCertificate creates a new intermediate CA certificate
	NewIntermediateCertificate(
		manager ICertificateManager,
		commonName string, // commonName The name of the intermediate CA
		serialNumber SerialNumber, // serialNumber Serial Number of the intermediate certificate
		parentCertificate ICertificate, // parentCertificate The parent certificate, typically the root CA
		parentPrivateKey IPrivateKey, // parentPrivateKey Private key of the parent
		expiration time.Duration, // The expiration duration
	) (ICertificate, error)

	// NewServerCertificate creates a new server certificate
	NewServerCertificate(
		manager ICertificateManager,
		serialNumber SerialNumber, // Serial Number of the server certificate
		parentCertificate ICertificate, // The parent certificate, typically the intermediate or root certificate
		privateKey IPrivateKey, // Private key of the parent
		dnsNames []string, // List of domain names the certificate is valid for
		expiration time.Duration,
	) (ICertificate, error)

	// NewClientCertificate creates a new client certificate
	NewClientCertificate(
		manager ICertificateManager,
		commonName string, // The name of the client
		serialNumber SerialNumber, // Serial Number of the client certificate
		parentCertificate ICertificate, // The parent certificate, typically the intermediate or root certificate
		privateKey IPrivateKey, // Private key of the parent
		expiration time.Duration,
	) (ICertificate, error)
}

// ICertificate describes an interface for Certificate model
type ICertificate interface {
	IsCA() bool
	GetSerialNumber() SerialNumber
	GetOrganizationID() string
	GetOrganizationName() string
	GetOrganization() []string
	GetSignedBy() SerialNumber
	GetCertificate() *x509.Certificate
}

// IPrivateKey describes an interface for PrivateKey model
type IPrivateKey interface {
	GetSerialNumber() SerialNumber
	GetKeyType() KeyType
	GetPublicKey() any
	CreateCertificate(
		manager ICertificateManager,
		template *x509.Certificate,
		parent *x509.Certificate,
	) (*x509.Certificate, error)
}

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

// IOrganizationService defines the interface for storing organization models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type IOrganizationService interface {
	GetExistingOrganization(id string) (IOrganization, error)
	CreateOrganization(certificate IOrganization) (IOrganization, error)
}

// ICertificateService defines the interface for storing certificate models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its data layer.
type ICertificateService interface {
	GetExistingCertificate(serialNumber SerialNumber) (ICertificate, error)
	CreateCertificate(certificate ICertificate) (ICertificate, error)
}

// IPrivateKeyService defines the interface for storing private keys,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type IPrivateKeyService interface {

	// GetExistingPrivateKey only returns public properties of the private key
	GetExistingPrivateKey(serialNumber SerialNumber) (IPrivateKey, error)
	CreatePrivateKey(key IPrivateKey) (IPrivateKey, error)
}

// IOrganizationController defines the interface for organization storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type IOrganizationController interface {

	// UsesOrganizationService returns true if this controller is using the
	// specified service. We're intentionally not returning a reference to the
	// service because we want to keep all the control inside the controller
	UsesOrganizationService(service IOrganizationService) bool

	GetExistingOrganization(id string) (IOrganization, error)
	CreateOrganization(certificate IOrganization) (IOrganization, error)
}

// ICertificateController defines the interface for certificate storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type ICertificateController interface {

	// UsesCertificateService returns true if this controller is using the
	// specified service. We're intentionally not returning a reference to the
	// service because we want to keep all the control inside the controller
	UsesCertificateService(service ICertificateService) bool

	GetExistingCertificate(serialNumber SerialNumber) (ICertificate, error)
	CreateCertificate(certificate ICertificate) (ICertificate, error)
}

// IPrivateKeyController defines the interface for key storage operations,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface within the controller package, it supports easy substitution of its
// implementation, thereby promoting loose coupling between the application's
// business logic and its data layer.
type IPrivateKeyController interface {

	// UsesPrivateKeyService returns true if this controller is using the
	// specified service. We're intentionally not returning a reference to the
	// service because we want to keep all the control inside the controller
	UsesPrivateKeyService(service IPrivateKeyService) bool

	// GetExistingPrivateKey only returns public properties of the private key
	GetExistingPrivateKey(serialNumber SerialNumber) (IPrivateKey, error)
	CreatePrivateKey(key IPrivateKey) (IPrivateKey, error)
}
