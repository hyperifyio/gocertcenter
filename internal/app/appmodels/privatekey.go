// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
)

// PrivateKeyModel model implements PrivateKey
type PrivateKeyModel struct {

	// organization is the organization this key belongs to
	organization string

	// certificates is the path to the certificate as serial numbers from root certificate
	certificates []SerialNumber

	// The type of the key
	keyType KeyType

	// data is the internal private key data
	data any
}

func (k *PrivateKeyModel) GetSerialNumber() SerialNumber {
	if len(k.certificates) <= 0 {
		return nil
	}
	return k.certificates[len(k.certificates)-1]
}

func (k *PrivateKeyModel) GetParents() []SerialNumber {
	if len(k.certificates) <= 1 {
		return []SerialNumber{}
	}
	originalSlice := k.certificates[:len(k.certificates)-1]
	sliceCopy := make([]SerialNumber, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (k *PrivateKeyModel) GetCertificates() []SerialNumber {
	originalSlice := k.certificates
	sliceCopy := make([]SerialNumber, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

func (k *PrivateKeyModel) GetOrganizationID() string {
	return k.organization
}

func (k *PrivateKeyModel) GetPrivateKey() any {
	return k.data
}

func (k *PrivateKeyModel) GetKeyType() KeyType {
	return k.keyType
}

func (k *PrivateKeyModel) GetPublicKey() any {
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
	certificates []SerialNumber,
	keyType KeyType,
	data any,
) *PrivateKeyModel {
	return &PrivateKeyModel{
		certificates: certificates,
		data:         data,
		keyType:      keyType,
		organization: organization,
	}
}

// Compile time assertion for implementing the interface
var _ PrivateKey = (*PrivateKeyModel)(nil)
