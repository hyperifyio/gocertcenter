// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"math/big"
)

// PrivateKeyModel model implements PrivateKey
type PrivateKeyModel struct {

	// organization is the organization this key belongs to
	organization string

	// certificate is the serial number of the certificate this private key is for
	certificate *big.Int

	// The type of the key
	keyType KeyType

	// data is the internal private key data
	data any
}

func (k *PrivateKeyModel) SerialNumber() *big.Int {
	return k.certificate
}

func (k *PrivateKeyModel) OrganizationID() string {
	return k.organization
}

func (k *PrivateKeyModel) PrivateKey() any {
	return k.data
}

func (k *PrivateKeyModel) KeyType() KeyType {
	return k.keyType
}

func (k *PrivateKeyModel) PublicKey() any {
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
	certificate *big.Int,
	keyType KeyType,
	data any,
) *PrivateKeyModel {
	return &PrivateKeyModel{
		certificate:  certificate,
		data:         data,
		keyType:      keyType,
		organization: organization,
	}
}

// Compile time assertion for implementing the interface
var _ PrivateKey = (*PrivateKeyModel)(nil)
