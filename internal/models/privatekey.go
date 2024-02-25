// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/dtos"
)

// PrivateKey model implements IPrivateKey
type PrivateKey struct {

	// organization is the organization this key belongs to
	organization string

	// certificates is the path to the certificate as serial numbers from root certificate
	certificates []ISerialNumber

	// The type of the key
	keyType KeyType

	// data is the internal private key data
	data any
}

// Compile time assertion for implementing the interface
var _ IPrivateKey = (*PrivateKey)(nil)

func (k *PrivateKey) GetDTO() dtos.PrivateKeyDTO {
	return dtos.NewPrivateKeyDTO(
		k.GetSerialNumber().String(),
		k.keyType.String(),
		"",
	)
}

func (k *PrivateKey) GetSerialNumber() ISerialNumber {
	if len(k.certificates) <= 0 {
		return nil
	}
	return k.certificates[len(k.certificates)-1]
}

func (k *PrivateKey) GetParents() []ISerialNumber {
	if len(k.certificates) <= 1 {
		return []ISerialNumber{}
	}
	return k.certificates[:len(k.certificates)-1]
}

func (k *PrivateKey) GetCertificates() []ISerialNumber {
	return k.certificates
}

func (k *PrivateKey) GetOrganizationID() string {
	return k.organization
}

func (k *PrivateKey) GetKeyType() KeyType {
	return k.keyType
}

func (k *PrivateKey) GetPublicKey() any {
	switch k := k.data.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

func (k *PrivateKey) CreateCertificate(
	manager ICertificateManager,
	template *x509.Certificate,
	parent *x509.Certificate,
) (*x509.Certificate, error) {
	bytes, err := manager.CreateCertificate(rand.Reader, template, parent, k.GetPublicKey(), k.data)
	if err != nil {
		return nil, fmt.Errorf("failed to create certificate: %w", err)
	}
	cert, err := manager.ParseCertificate(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate after creating it: %w", err)
	}
	return cert, nil
}

// NewPrivateKey creates a private key model from existing data
func NewPrivateKey(
	organization string,
	certificates []ISerialNumber,
	keyType KeyType,
	data any,
) *PrivateKey {
	return &PrivateKey{
		certificates: certificates,
		data:         data,
		keyType:      keyType,
		organization: organization,
	}
}
