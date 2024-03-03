// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

// Organization describes an interface for OrganizationModel model
type Organization interface {

	// GetID returns unique identifier for this organization
	ID() string

	// GetName returns the primary organization name
	Name() string

	// GetNames returns the full name of the organization including department
	Names() []string
}

// Certificate describes an interface for CertificateModel model
type Certificate interface {
	CommonName() string

	NotBefore() time.Time
	NotAfter() time.Time

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

	// SignedBy returns the parent certificate serial number
	SignedBy() *big.Int

	SerialNumber() *big.Int
	OrganizationID() string
	OrganizationName() string
	Organization() []string
	Certificate() *x509.Certificate
}

// PublicKey describes an interface for PublicKey model
type PublicKey interface {

	// GetPublicKey returns the public key
	PublicKey() any
}

// PrivateKey describes an interface for PrivateKeyModel model
type PrivateKey interface {

	// PrivateKey returns the internal private key
	PrivateKey() any

	// KeyType returns the type of the internal key
	KeyType() KeyType

	// PublicKey returns the public key
	PublicKey() any

	// OrganizationID returns the organization this key belongs to
	OrganizationID() string

	// SerialNumber returns the serial number of the certificate which this
	// key belongs to
	SerialNumber() *big.Int
}

// RevokedCertificate describes an interface for RevokedCertificateModel model
type RevokedCertificate interface {

	// SerialNumber returns the serial number of the certificate which was revoked
	SerialNumber() *big.Int

	RevocationTime() time.Time
	ExpirationTime() time.Time

	RevokedCertificate() pkix.RevokedCertificate
}

// OrganizationRepository defines the interface for storing organization models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type OrganizationRepository interface {
	FindAll() ([]Organization, error)
	FindById(organization string) (Organization, error)
	Save(certificate Organization) (Organization, error)
}

// CertificateRepository defines the interface for storing certificate models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its data layer.
type CertificateRepository interface {
	FindAllByOrganization(organization string) ([]Certificate, error)

	// FindAllByOrganizationAndSignedBy returns all certificates signed by this certificate
	FindAllByOrganizationAndSignedBy(organization string, certificate *big.Int) ([]Certificate, error)

	FindByOrganizationAndSerialNumber(organization string, certificate *big.Int) (Certificate, error)
	Save(certificate Certificate) (Certificate, error)
}

// PrivateKeyRepository defines the interface for storing private keys,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type PrivateKeyRepository interface {

	// FindByOrganizationAndSerialNumber only returns public properties of the private key
	FindByOrganizationAndSerialNumber(organization string, certificate *big.Int) (PrivateKey, error)
	Save(key PrivateKey) (PrivateKey, error)
}

// ApplicationController controls an application. An application may own one
// or more organizations.
type ApplicationController interface {

	// UsesOrganizationService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesOrganizationService(service OrganizationRepository) bool

	// UsesCertificateService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesCertificateService(service CertificateRepository) bool

	// UsesPrivateKeyService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesPrivateKeyService(service PrivateKeyRepository) bool

	// OrganizationCollection returns all organizations
	OrganizationCollection() ([]Organization, error)

	// Organization returns an organization model by an organization ID
	Organization(organization string) (Organization, error)

	// OrganizationController returns an organization controller by an organization ID
	OrganizationController(name string) (OrganizationController, error)

	// NewOrganization creates a new organization
	NewOrganization(model Organization) (Organization, error)
}

// OrganizationController controls an organization owned by the application. An
// organization may own one or more root certificates.
type OrganizationController interface {

	// OrganizationID returns the organization ID which this controller controls
	OrganizationID() string

	// Organization returns the model of the organization this controller controls
	Organization() Organization

	// ApplicationController returns the parent controller who owns this organization controller
	ApplicationController() ApplicationController

	// CertificateCollection returns all the root level certificates for the organization
	CertificateCollection() ([]Certificate, error)

	// CertificateController returns a controller for a root certificate specified by its serial number
	//  * serialNumber - The serial number of the root certificate
	CertificateController(serialNumber *big.Int) (CertificateController, error)

	// Certificate returns a model for a root certificate specified by its serial number
	//  * serialNumber - The serial number of the root certificate
	Certificate(serialNumber *big.Int) (Certificate, error)

	// SetExpirationDuration sets the expiration duration used in NewRootCertificate
	//  * expiration - the expiration duration
	SetExpirationDuration(expiration time.Duration)

	// NewRootCertificate creates a new root certificate for the organization
	//  * commonName - The name of the root CA
	NewRootCertificate(commonName string) (Certificate, error)

	UsesOrganizationService(service OrganizationRepository) bool
	UsesApplicationController(service ApplicationController) bool

	RevokeCertificate(certificate Certificate) (RevokedCertificate, error)
}

// CertificateController controls a certificate owned by the organization. It
// can be directly owned by the organization (when it's a root certificate),
// or it may be owned by another root or intermediate certificate. It also owns
// one private key.
type CertificateController interface {

	// ApplicationController returns the parent controller who owns this
	// organization controller
	ApplicationController() ApplicationController

	// OrganizationID returns the organization ID who owns the certificate
	// this controller controls
	OrganizationID() string

	// Organization returns the model of the organization who owns the
	// certificate this controller controls
	Organization() Organization

	// OrganizationController returns the organization controller who owns
	// the certificate this controller controls
	OrganizationController() OrganizationController

	// Certificate returns the model of the certificate this controller
	// controls
	Certificate() Certificate

	// ChildCertificateCollection returns all child certificates
	ChildCertificateCollection(certificateType string) ([]Certificate, error)

	// Certificate returns a child certificate model
	//  * serialNumber - The serial number of the child certificate
	ChildCertificate(serialNumber *big.Int) (Certificate, error)

	// GetChildCertificateController returns a child certificate controller
	//  * serialNumber - The serial number of the child certificate
	ChildCertificateController(serialNumber *big.Int) (CertificateController, error)

	// ParentCertificate returns the parent certificate model if this
	// certificate is not a root certificate
	ParentCertificate() Certificate

	// GetParentCertificateController returns the parent certificate controller
	// if this certificate is not a root certificate
	ParentCertificateController() CertificateController

	// PrivateKey returns the private key model of this certificate
	PrivateKey() (PrivateKey, error)

	// GetPrivateKeyController returns the private key controller of this
	// certificate
	PrivateKeyController() (PrivateKeyController, error)

	// SetExpirationDuration sets the expiration duration used in
	// NewIntermediateCertificate, NewServerCertificate, or NewClientCertificate
	//  * expiration - The expiration duration
	SetExpirationDuration(expiration time.Duration)

	// NewIntermediateCertificate creates a new child certificate as an
	// intermediate CA certificate
	//  * commonName - The name of the intermediate CA
	NewIntermediateCertificate(commonName string) (Certificate, PrivateKey, error)

	// NewServerCertificate creates a new server certificate.
	//   - dnsNames: List of domain names the new certificate. The first one is
	//     used as a common name as well.
	NewServerCertificate(dnsNames ...string) (Certificate, PrivateKey, error)

	// NewClientCertificate creates a new client certificate
	//  * commonName - The name of the client
	NewClientCertificate(commonName string) (Certificate, PrivateKey, error)
}

// PrivateKeyController controls a private key owned by the certificate
type PrivateKeyController interface {

	// GetApplicationController returns the parent controller who owns this
	// organization controller
	ApplicationController() ApplicationController

	// GetOrganizationID returns the organization ID who owns the certificate
	// this controller controls
	OrganizationID() string

	// Organization returns the model of the organization who owns the
	// certificate this controller controls
	Organization() Organization

	// GetOrganizationController returns the model of the organization who owns the
	// certificate this controller controls
	OrganizationController() OrganizationController

	// Certificate returns the model of the certificate this controller
	// controls
	Certificate() Certificate

	// GetCertificateController returns the controller of the certificate
	// controls
	CertificateController() CertificateController
}
