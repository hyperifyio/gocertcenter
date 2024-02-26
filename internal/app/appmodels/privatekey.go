// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
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
	originalSlice := k.certificates[:len(k.certificates)-1]
	sliceCopy := make([]ISerialNumber, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (k *PrivateKey) GetCertificates() []ISerialNumber {
	originalSlice := k.certificates
	sliceCopy := make([]ISerialNumber, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (k *PrivateKey) GetOrganizationID() string {
	return k.organization
}

func (k *PrivateKey) GetPrivateKey() any {
	return k.data
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

// NewPrivateKey creates a private key model from existing data
//   - organization is the organization
//   - certificates is the parent certificate serial numbers owning this key
//   - keyType is the private key type
//   - data is the private key data
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

// Compile time assertion for implementing the interface
var _ IPrivateKey = (*PrivateKey)(nil)
