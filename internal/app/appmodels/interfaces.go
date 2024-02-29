// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

type SerialNumber interface {
	String() string
	Value() *big.Int
	Cmp(s2 SerialNumber) int
	Sign() int
}

// Organization describes an interface for OrganizationModel model
type Organization interface {

	// GetID returns unique identifier for this organization
	GetID() string

	// GetName returns the primary organization name
	GetName() string

	// GetNames returns the full name of the organization including department
	GetNames() []string
}

// Certificate describes an interface for CertificateModel model
type Certificate interface {

	// GetCommonName
	GetCommonName() string

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

	// GetParents returns all parent certificate serial numbers
	GetParents() []SerialNumber

	GetSerialNumber() SerialNumber
	GetOrganizationID() string
	GetOrganizationName() string
	GetOrganization() []string
	GetSignedBy() SerialNumber
	GetCertificate() *x509.Certificate
}

// PublicKey describes an interface for PublicKey model
type PublicKey interface {

	// GetPublicKey returns the public key
	GetPublicKey() any
}

// PrivateKey describes an interface for PrivateKeyModel model
type PrivateKey interface {

	// GetPrivateKey returns the internal private key
	GetPrivateKey() any

	// GetKeyType returns the type of the internal key
	GetKeyType() KeyType

	// GetPublicKey returns the public key
	GetPublicKey() any

	// GetOrganizationID returns the organization this key belongs to
	GetOrganizationID() string

	// GetParents returns all parent certificate serial numbers
	GetParents() []SerialNumber

	// GetCertificates returns all serial numbers from root certificate to the
	// certificate this key belongs to
	GetCertificates() []SerialNumber

	// GetSerialNumber returns the serial number of the certificate which this
	// key belongs to
	GetSerialNumber() SerialNumber
}

// RevokedCertificate describes an interface for RevokedCertificateModel model
type RevokedCertificate interface {

	// GetSerialNumber returns the serial number of the certificate which was revoked
	GetSerialNumber() SerialNumber

	GetRevocationTime() time.Time
	GetExpirationTime() time.Time

	GetRevokedCertificate() pkix.RevokedCertificate
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
	FindAllByOrganizationAndSerialNumbers(organization string, certificates []SerialNumber) ([]Certificate, error)
	FindByOrganizationAndSerialNumbers(organization string, certificates []SerialNumber) (Certificate, error)
	Save(certificate Certificate) (Certificate, error)
}

// PrivateKeyRepository defines the interface for storing private keys,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type PrivateKeyRepository interface {

	// FindByOrganizationAndSerialNumbers only returns public properties of the private key
	FindByOrganizationAndSerialNumbers(organization string, certificates []SerialNumber) (PrivateKey, error)
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

	// GetOrganizationCollection returns all organizations
	GetOrganizationCollection() ([]Organization, error)

	// GetOrganizationModel returns an organization model by an organization ID
	GetOrganizationModel(organization string) (Organization, error)

	// GetOrganizationController returns an organization controller by an organization ID
	GetOrganizationController(name string) (OrganizationController, error)

	// NewOrganization creates a new organization
	NewOrganization(model Organization) (Organization, error)
}

// OrganizationController controls an organization owned by the application. An
// organization may own one or more root certificates.
type OrganizationController interface {

	// GetOrganizationID returns the organization ID which this controller controls
	GetOrganizationID() string

	// GetOrganizationModel returns the model of the organization this controller controls
	GetOrganizationModel() Organization

	// GetApplicationController returns the parent controller who owns this organization controller
	GetApplicationController() ApplicationController

	// GetCertificateCollection returns all the root level certificates for the organization
	GetCertificateCollection() ([]Certificate, error)

	// GetCertificateController returns a controller for a root certificate specified by its serial number
	//  * serialNumber - The serial number of the root certificate
	GetCertificateController(serialNumber SerialNumber) (CertificateController, error)

	// GetCertificateModel returns a model for a root certificate specified by its serial number
	//  * serialNumber - The serial number of the root certificate
	GetCertificateModel(serialNumber SerialNumber) (Certificate, error)

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

	// GetApplicationController returns the parent controller who owns this
	// organization controller
	GetApplicationController() ApplicationController

	// GetOrganizationID returns the organization ID who owns the certificate
	// this controller controls
	GetOrganizationID() string

	// GetOrganizationModel returns the model of the organization who owns the
	// certificate this controller controls
	GetOrganizationModel() Organization

	// GetOrganizationController returns the organization controller who owns
	// the certificate this controller controls
	GetOrganizationController() OrganizationController

	// GetCertificateModel returns the model of the certificate this controller
	// controls
	GetCertificateModel() Certificate

	// GetChildCertificateCollection returns all child certificates
	GetChildCertificateCollection(certificateType string) ([]Certificate, error)

	// GetChildCertificateModel returns a child certificate model
	//  * serialNumber - The serial number of the child certificate
	GetChildCertificateModel(serialNumber SerialNumber) (Certificate, error)

	// GetChildCertificateController returns a child certificate controller
	//  * serialNumber - The serial number of the child certificate
	GetChildCertificateController(serialNumber SerialNumber) (CertificateController, error)

	// GetParentCertificateModel returns the parent certificate model if this
	// certificate is not a root certificate
	GetParentCertificateModel() Certificate

	// GetParentCertificateController returns the parent certificate controller
	// if this certificate is not a root certificate
	GetParentCertificateController() CertificateController

	// GetPrivateKeyModel returns the private key model of this certificate
	GetPrivateKeyModel() (PrivateKey, error)

	// GetPrivateKeyController returns the private key controller of this
	// certificate
	GetPrivateKeyController() (PrivateKeyController, error)

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
	GetApplicationController() ApplicationController

	// GetOrganizationID returns the organization ID who owns the certificate
	// this controller controls
	GetOrganizationID() string

	// GetOrganizationModel returns the model of the organization who owns the
	// certificate this controller controls
	GetOrganizationModel() Organization

	// GetOrganizationController returns the model of the organization who owns the
	// certificate this controller controls
	GetOrganizationController() OrganizationController

	// GetCertificateModel returns the model of the certificate this controller
	// controls
	GetCertificateModel() Certificate

	// GetCertificateController returns the controller of the certificate
	// controls
	GetCertificateController() CertificateController
}
