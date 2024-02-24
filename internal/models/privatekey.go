// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
)

// PrivateKey model implements IPrivateKey
type PrivateKey struct {

	// serialNumber of the certificate this key belongs to
	serialNumber ISerialNumber

	// The type of the key
	keyType KeyType

	// data is the internal private key data
	data any
}

// Compile time assertion for implementing the interface
var _ IPrivateKey = (*PrivateKey)(nil)

// NewPrivateKey creates a private key model from existing data
func NewPrivateKey(
	serialNumber ISerialNumber,
	keyType KeyType,
	data any,
) *PrivateKey {
	return &PrivateKey{
		serialNumber: serialNumber,
		data:         data,
		keyType:      keyType,
	}
}

func (k *PrivateKey) GetDTO() dtos.PrivateKeyDTO {
	return dtos.NewPrivateKeyDTO(
		k.serialNumber.String(),
		k.keyType.String(),
		"",
	)
}

func (k *PrivateKey) GetSerialNumber() ISerialNumber {
	return k.serialNumber
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
		// FIXME: Add tests for this scope
		return nil, fmt.Errorf("failed to parse certificate after creating it: %w", err)
	}
	return cert, nil
}

// GeneratePrivateKey creates a new private key of type keyType
func GeneratePrivateKey(
	serialNumber ISerialNumber,
	keyType KeyType,
	rsaBits int,
) (IPrivateKey, error) {
	var key any
	var err error
	switch keyType {
	case RSA:
		key, err = rsa.GenerateKey(rand.Reader, rsaBits)
	case ECDSA_P224:
		key, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case ECDSA_P256:
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case ECDSA_P384:
		key, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case ECDSA_P521:
		key, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	case Ed25519:
		_, key, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = errors.New(fmt.Sprintf("keyType is unsupported: %s", keyType.String()))
	}
	if err != nil {
		return nil, err
	}
	return NewPrivateKey(serialNumber, keyType, key), nil
}

// GenerateRSAPrivateKey creates a new private key of type keyType
func GenerateRSAPrivateKey(
	serialNumber ISerialNumber,
	rsaBits int,
) (IPrivateKey, error) {
	return GeneratePrivateKey(serialNumber, RSA, rsaBits)
}

// GenerateECDSAPrivateKey creates a new private key of type keyType
func GenerateECDSAPrivateKey(
	serialNumber ISerialNumber,
	keyType KeyType,
) (IPrivateKey, error) {
	return GeneratePrivateKey(serialNumber, keyType, 2048)
}

// GenerateEd25519PrivateKey creates a new private key of type keyType
func GenerateEd25519PrivateKey(
	serialNumber ISerialNumber,
) (IPrivateKey, error) {
	return GeneratePrivateKey(serialNumber, Ed25519, 2048)
}
