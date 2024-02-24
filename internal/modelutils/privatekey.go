// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package modelutils

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/models"
)

// GeneratePrivateKey creates a new private key of type models.KeyType
func GeneratePrivateKey(
	serialNumber models.ISerialNumber,
	keyType models.KeyType,
	rsaBits int,
) (models.IPrivateKey, error) {
	var key any
	var err error
	switch keyType {
	case models.RSA:
		key, err = rsa.GenerateKey(rand.Reader, rsaBits)
	case models.ECDSA_P224:
		key, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case models.ECDSA_P256:
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case models.ECDSA_P384:
		key, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case models.ECDSA_P521:
		key, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	case models.Ed25519:
		_, key, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = errors.New(fmt.Sprintf("keyType is unsupported: %s", keyType.String()))
	}
	if err != nil {
		return nil, err
	}
	return models.NewPrivateKey(serialNumber, keyType, key), nil
}

// GenerateRSAPrivateKey creates a new private key of type models.RSA
func GenerateRSAPrivateKey(
	serialNumber models.ISerialNumber,
	rsaBits int,
) (models.IPrivateKey, error) {
	return GeneratePrivateKey(serialNumber, models.RSA, rsaBits)
}

// GenerateECDSAPrivateKey creates a new private key of type models.KeyType
func GenerateECDSAPrivateKey(
	serialNumber models.ISerialNumber,
	keyType models.KeyType,
) (models.IPrivateKey, error) {
	return GeneratePrivateKey(serialNumber, keyType, 2048)
}

// GenerateEd25519PrivateKey creates a new private key of type models.Ed25519
func GenerateEd25519PrivateKey(
	serialNumber models.ISerialNumber,
) (models.IPrivateKey, error) {
	return GeneratePrivateKey(serialNumber, models.Ed25519, 2048)
}
