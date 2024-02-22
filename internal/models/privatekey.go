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
	"github.com/hyperifyio/gocertcenter/internal/managers"
)

// PrivateKey model
type PrivateKey struct {

	// serialNumber of the certificate this key belongs to
	serialNumber SerialNumber

	// The type of the key
	keyType KeyType

	// data is the internal private key data
	data any
}

// NewPrivateKey creates a private key model from existing data
func NewPrivateKey(
	serialNumber SerialNumber,
	keyType KeyType,
	data any,
) *PrivateKey {
	return &PrivateKey{
		serialNumber: serialNumber,
		data:         data,
		keyType:      keyType,
	}
}

// GeneratePrivateKey creates a new private key of type keyType
func GeneratePrivateKey(
	serialNumber SerialNumber,
	keyType KeyType,
	rsaBits int,
) (*PrivateKey, error) {
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
	serialNumber SerialNumber,
	rsaBits int,
) (*PrivateKey, error) {
	return GeneratePrivateKey(serialNumber, RSA, rsaBits)
}

// GenerateECDSAPrivateKey creates a new private key of type keyType
func GenerateECDSAPrivateKey(
	serialNumber SerialNumber,
	keyType KeyType,
) (*PrivateKey, error) {
	return GeneratePrivateKey(serialNumber, keyType, 2048)
}

// GenerateEd25519PrivateKey creates a new private key of type keyType
func GenerateEd25519PrivateKey(
	serialNumber SerialNumber,
) (*PrivateKey, error) {
	return GeneratePrivateKey(serialNumber, Ed25519, 2048)
}

func (k *PrivateKey) GetSerialNumber() SerialNumber {
	return k.serialNumber
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
	manager managers.ICertificateManager,
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
