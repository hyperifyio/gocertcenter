// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"io"
	"math/big"
	"time"
)

type ISerialNumber interface {
	String() string
	Value() *big.Int
	Cmp(s2 ISerialNumber) int
	Sign() int
}

// IOrganization describes an interface for Organization model
type IOrganization interface {

	// GetDTO returns a data transfer object
	GetDTO() dtos.OrganizationDTO

	// GetID returns unique identifier for this organization
	GetID() string

	// GetName returns the primary organization name
	GetName() string

	// GetNames returns the full name of the organization including department
	GetNames() []string
}

// ICertificate describes an interface for Certificate model
type ICertificate interface {

	// GetDTO returns a data transfer object
	GetDTO() dtos.CertificateDTO

	GetCommonName() string

	// IsCA -
	IsCA() bool

	// IsSelfSigned -
	IsSelfSigned() bool

	// IsRootCertificate - A root certificate is a top-level CA certificate used
	// to sign other certificates and is self-signed.
	IsRootCertificate() bool

	// IsIntermediateCertificate - An intermediate certificate is typically a CA
	// certificate but not the root CA. It can sign other certificates but is
	// itself signed by another CA.
	IsIntermediateCertificate() bool

	// IsServerCertificate - A server certificate is used to identify a server
	// to a client. It's typically not a CA certificate (IsCA is false), and it
	// has specific extended key usages.
	IsServerCertificate() bool

	// IsClientCertificate - A client certificate is used to identify a client
	// to a server. Similar to server certificates, these are not CA
	// certificates and have specific extended key usages.
	IsClientCertificate() bool

	// GetParents returns all parent certificate serial numbers
	GetParents() []ISerialNumber

	GetSerialNumber() ISerialNumber
	GetOrganizationID() string
	GetOrganizationName() string
	GetOrganization() []string
	GetSignedBy() ISerialNumber
	GetCertificate() *x509.Certificate
	GetPEM() []byte
}

// IPrivateKey describes an interface for PrivateKey model
type IPrivateKey interface {

	// GetDTO returns a data transfer object
	GetDTO() dtos.PrivateKeyDTO

	// GetOrganizationID returns the organization this key belongs to
	GetOrganizationID() string

	// GetParents returns all parent certificate serial numbers
	GetParents() []ISerialNumber

	// GetCertificates returns all serial numbers from root certificate to this one
	GetCertificates() []ISerialNumber

	GetSerialNumber() ISerialNumber
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
	GetExistingOrganization(organization string) (IOrganization, error)
	CreateOrganization(certificate IOrganization) (IOrganization, error)
}

// ICertificateService defines the interface for storing certificate models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its data layer.
type ICertificateService interface {
	GetExistingCertificate(organization string, certificates []ISerialNumber) (ICertificate, error)
	CreateCertificate(certificate ICertificate) (ICertificate, error)
}

// IPrivateKeyService defines the interface for storing private keys,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type IPrivateKeyService interface {

	// GetExistingPrivateKey only returns public properties of the private key
	GetExistingPrivateKey(organization string, certificates []ISerialNumber) (IPrivateKey, error)
	CreatePrivateKey(key IPrivateKey) (IPrivateKey, error)
}

// IOrganizationController defines the interface for organization operations
type IOrganizationController interface {

	// UsesOrganizationService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesOrganizationService(service IOrganizationService) bool

	GetExistingOrganization(organization string) (IOrganization, error)

	// CreateOrganization creates a new organization
	CreateOrganization(certificate IOrganization) (IOrganization, error)

	// NewRootCertificate creates a new root CA certificate
	NewRootCertificate(
		o IOrganization,
		manager ICertificateManager,
		commonName string, // The name of the root CA
		privateKey IPrivateKey, // Private key for self signing
		expiration time.Duration,
	) (ICertificate, error)

	// NewIntermediateCertificate creates a new intermediate CA certificate
	NewIntermediateCertificate(
		o IOrganization,
		manager ICertificateManager,
		commonName string, // commonName The name of the intermediate CA
		serialNumber ISerialNumber, // serialNumber Serial Number of the intermediate certificate
		parentCertificate ICertificate, // parentCertificate The parent certificate, typically the root CA
		parentPrivateKey IPrivateKey, // parentPrivateKey Private key of the parent
		expiration time.Duration, // The expiration duration
	) (ICertificate, error)

	// NewServerCertificate creates a new server certificate
	NewServerCertificate(
		o IOrganization,
		manager ICertificateManager,
		serialNumber ISerialNumber, // Serial Number of the server certificate
		parentCertificate ICertificate, // The parent certificate, typically the intermediate or root certificate
		privateKey IPrivateKey, // Private key of the parent
		dnsNames []string, // List of domain names the certificate is valid for
		expiration time.Duration,
	) (ICertificate, error)

	// NewClientCertificate creates a new client certificate
	NewClientCertificate(
		o IOrganization,
		manager ICertificateManager,
		commonName string, // The name of the client
		serialNumber ISerialNumber, // Serial Number of the client certificate
		parentCertificate ICertificate, // The parent certificate, typically the intermediate or root certificate
		privateKey IPrivateKey, // Private key of the parent
		expiration time.Duration,
	) (ICertificate, error)
}

// ICertificateController defines the interface for certificate operations
type ICertificateController interface {

	// UsesCertificateService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesCertificateService(service ICertificateService) bool

	GetExistingCertificate(organization string, certificates []ISerialNumber) (ICertificate, error)

	CreateCertificate(certificate ICertificate) (ICertificate, error)
}

// IPrivateKeyController defines the interface for private key operations
type IPrivateKeyController interface {

	// UsesPrivateKeyService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesPrivateKeyService(service IPrivateKeyService) bool

	// GetExistingPrivateKey only returns public properties of the private key
	GetExistingPrivateKey(organization string, certificates []ISerialNumber) (IPrivateKey, error)

	CreatePrivateKey(key IPrivateKey) (IPrivateKey, error)
}
