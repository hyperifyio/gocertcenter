// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/x509"
	"crypto/x509/pkix"
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

	// GetID returns unique identifier for this organization
	GetID() string

	// GetName returns the primary organization name
	GetName() string

	// GetNames returns the full name of the organization including department
	GetNames() []string
}

// ICertificate describes an interface for Certificate model
type ICertificate interface {

	// GetCommonName
	GetCommonName() string

	GetNotBefore() time.Time
	GetNotAfter() time.Time

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
}

// IPublicKey describes an interface for PublicKey model
type IPublicKey interface {

	// GetPublicKey returns the public key
	GetPublicKey() any
}

// IPrivateKey describes an interface for PrivateKey model
type IPrivateKey interface {

	// GetPrivateKey returns the internal private key
	GetPrivateKey() any

	// GetKeyType returns the type of the internal key
	GetKeyType() KeyType

	// GetPublicKey returns the public key
	GetPublicKey() any

	// GetOrganizationID returns the organization this key belongs to
	GetOrganizationID() string

	// GetParents returns all parent certificate serial numbers
	GetParents() []ISerialNumber

	// GetCertificates returns all serial numbers from root certificate to the
	// certificate this key belongs to
	GetCertificates() []ISerialNumber

	// GetSerialNumber returns the serial number of the certificate which this
	// key belongs to
	GetSerialNumber() ISerialNumber
}

// IRevokedCertificate describes an interface for RevokedCertificate model
type IRevokedCertificate interface {

	// GetSerialNumber returns the serial number of the certificate which was revoked
	GetSerialNumber() ISerialNumber

	GetRevocationTime() time.Time
	GetExpirationTime() time.Time

	GetRevokedCertificate() pkix.RevokedCertificate
}

// IOrganizationService defines the interface for storing organization models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type IOrganizationService interface {
	FindAll() ([]IOrganization, error)
	FindById(organization string) (IOrganization, error)
	Save(certificate IOrganization) (IOrganization, error)
}

// ICertificateService defines the interface for storing certificate models,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its data layer.
type ICertificateService interface {
	FindAllByOrganization(organization string) ([]ICertificate, error)
	FindAllByOrganizationAndSerialNumbers(organization string, certificates []ISerialNumber) ([]ICertificate, error)
	FindByOrganizationAndSerialNumbers(organization string, certificates []ISerialNumber) (ICertificate, error)
	Save(certificate ICertificate) (ICertificate, error)
}

// IPrivateKeyService defines the interface for storing private keys,
// facilitating the abstraction of data access mechanisms. By declaring this
// interface it supports easy substitution of its implementation, thereby
// promoting loose coupling between the application's business logic and its
// data layer.
type IPrivateKeyService interface {

	// GetExistingPrivateKey only returns public properties of the private key
	FindByOrganizationAndSerialNumbers(organization string, certificates []ISerialNumber) (IPrivateKey, error)
	Save(key IPrivateKey) (IPrivateKey, error)
}

// IApplicationController controls an application. An application may own one
// or more organizations.
type IApplicationController interface {

	// UsesOrganizationService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesOrganizationService(service IOrganizationService) bool

	// UsesCertificateService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesCertificateService(service ICertificateService) bool

	// UsesPrivateKeyService returns true if this controller is using the
	// specified data layer service. We're intentionally not returning a
	// reference to the service because we want to keep all the control inside
	// the controller
	UsesPrivateKeyService(service IPrivateKeyService) bool

	// GetOrganizationCollection returns all organizations
	GetOrganizationCollection() ([]IOrganization, error)

	// GetOrganizationModel returns an organization model by an organization ID
	GetOrganizationModel(organization string) (IOrganization, error)

	// GetOrganizationController returns an organization controller by an organization ID
	GetOrganizationController(name string) (IOrganizationController, error)

	// NewOrganization creates a new organization
	NewOrganization(model IOrganization) (IOrganization, error)
}

// IOrganizationController controls an organization owned by the application. An
// organization may own one or more root certificates.
type IOrganizationController interface {

	// GetOrganizationID returns the organization ID which this controller controls
	GetOrganizationID() string

	// GetOrganizationModel returns the model of the organization this controller controls
	GetOrganizationModel() IOrganization

	// GetApplicationController returns the parent controller who owns this organization controller
	GetApplicationController() IApplicationController

	// GetCertificateCollection returns all the root level certificates for the organization
	GetCertificateCollection() ([]ICertificate, error)

	// GetCertificateController returns a controller for a root certificate specified by its serial number
	//  * serialNumber - The serial number of the root certificate
	GetCertificateController(serialNumber ISerialNumber) (ICertificateController, error)

	// GetCertificateModel returns a model for a root certificate specified by its serial number
	//  * serialNumber - The serial number of the root certificate
	GetCertificateModel(serialNumber ISerialNumber) (ICertificate, error)

	// SetExpirationDuration sets the expiration duration used in NewRootCertificate
	//  * expiration - the expiration duration
	SetExpirationDuration(expiration time.Duration)

	// NewRootCertificate creates a new root certificate for the organization
	//  * commonName - The name of the root CA
	NewRootCertificate(commonName string) (ICertificate, error)

	UsesOrganizationService(service IOrganizationService) bool
	UsesApplicationController(service IApplicationController) bool

	RevokeCertificate(certificate ICertificate) (IRevokedCertificate, error)
}

// ICertificateController controls a certificate owned by the organization. It
// can be directly owned by the organization (when it's a root certificate),
// or it may be owned by another root or intermediate certificate. It also owns
// one private key.
type ICertificateController interface {

	// GetApplicationController returns the parent controller who owns this
	// organization controller
	GetApplicationController() IApplicationController

	// GetOrganizationID returns the organization ID who owns the certificate
	// this controller controls
	GetOrganizationID() string

	// GetOrganizationModel returns the model of the organization who owns the
	// certificate this controller controls
	GetOrganizationModel() IOrganization

	// GetOrganizationController returns the organization controller who owns
	// the certificate this controller controls
	GetOrganizationController() IOrganizationController

	// GetCertificateModel returns the model of the certificate this controller
	// controls
	GetCertificateModel() ICertificate

	// GetChildCertificateCollection returns all child certificates
	GetChildCertificateCollection(certificateType string) ([]ICertificate, error)

	// GetChildCertificateModel returns a child certificate model
	//  * serialNumber - The serial number of the child certificate
	GetChildCertificateModel(serialNumber ISerialNumber) (ICertificate, error)

	// GetChildCertificateController returns a child certificate controller
	//  * serialNumber - The serial number of the child certificate
	GetChildCertificateController(serialNumber ISerialNumber) (ICertificateController, error)

	// GetParentCertificateModel returns the parent certificate model if this
	// certificate is not a root certificate
	GetParentCertificateModel() ICertificate

	// GetParentCertificateController returns the parent certificate controller
	// if this certificate is not a root certificate
	GetParentCertificateController() ICertificateController

	// GetPrivateKeyModel returns the private key model of this certificate
	GetPrivateKeyModel() (IPrivateKey, error)

	// GetPrivateKeyController returns the private key controller of this
	// certificate
	GetPrivateKeyController() (IPrivateKeyController, error)

	// SetExpirationDuration sets the expiration duration used in
	// NewIntermediateCertificate, NewServerCertificate, or NewClientCertificate
	//  * expiration - The expiration duration
	SetExpirationDuration(expiration time.Duration)

	// NewIntermediateCertificate creates a new child certificate as an
	// intermediate CA certificate
	//  * commonName - The name of the intermediate CA
	NewIntermediateCertificate(commonName string) (ICertificate, IPrivateKey, error)

	// NewServerCertificate creates a new server certificate.
	//   - dnsNames: List of domain names the new certificate. The first one is
	//     used as a common name as well.
	NewServerCertificate(dnsNames ...string) (ICertificate, IPrivateKey, error)

	// NewClientCertificate creates a new client certificate
	//  * commonName - The name of the client
	NewClientCertificate(commonName string) (ICertificate, IPrivateKey, error)
}

// IPrivateKeyController controls a private key owned by the certificate
type IPrivateKeyController interface {

	// GetApplicationController returns the parent controller who owns this
	// organization controller
	GetApplicationController() IApplicationController

	// GetOrganizationID returns the organization ID who owns the certificate
	// this controller controls
	GetOrganizationID() string

	// GetOrganizationModel returns the model of the organization who owns the
	// certificate this controller controls
	GetOrganizationModel() IOrganization

	// GetOrganizationController returns the model of the organization who owns the
	// certificate this controller controls
	GetOrganizationController() IOrganizationController

	// GetCertificateModel returns the model of the certificate this controller
	// controls
	GetCertificateModel() ICertificate

	// GetCertificateController returns the controller of the certificate
	// controls
	GetCertificateController() ICertificateController
}
