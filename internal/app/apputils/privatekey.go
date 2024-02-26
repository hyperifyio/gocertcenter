// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// GeneratePrivateKey creates a new private key of type models.KeyType
//   - organization: The organization ID
//   - certificates: The certificate signatures from root certificate to the one owning this private key
//   - keyType: The key type to generate
func GeneratePrivateKey(
	organization string,
	certificates []appmodels.ISerialNumber,
	keyType appmodels.KeyType,
) (appmodels.IPrivateKey, error) {

	if organization == "" {
		return nil, fmt.Errorf("GeneratePrivateKey: organization: must not be empty")
	}

	if len(certificates) <= 0 {
		return nil, fmt.Errorf("GeneratePrivateKey: certificates: must have at least one serial number")
	}

	var key any
	var err error
	switch keyType {
	case appmodels.RSA_1024:
		key, err = rsa.GenerateKey(rand.Reader, 1024)
	case appmodels.RSA_2048:
		key, err = rsa.GenerateKey(rand.Reader, 2048)
	case appmodels.RSA_3072:
		key, err = rsa.GenerateKey(rand.Reader, 3072)
	case appmodels.RSA_4096:
		key, err = rsa.GenerateKey(rand.Reader, 4096)
	case appmodels.ECDSA_P224:
		key, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case appmodels.ECDSA_P256:
		key, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case appmodels.ECDSA_P384:
		key, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case appmodels.ECDSA_P521:
		key, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	case appmodels.Ed25519:
		_, key, err = ed25519.GenerateKey(rand.Reader)
	default:
		err = fmt.Errorf("GeneratePrivateKey: keyType: iunsupported: %s", keyType.String())
	}
	if err != nil {
		return nil, fmt.Errorf("GeneratePrivateKey: failed to generate key: %w", err)
	}
	return appmodels.NewPrivateKey(organization, certificates, keyType, key), nil
}

// GenerateRSAPrivateKey creates a new RSA private key
//   - organization: Organization ID
//   - certificates: Certificate serial numbers from root certificate to the one owning this key
//   - keyType: Should be appmodels.RSA_1024, appmodels.RSA_2048, appmodels.RSA_3072 or appmodels.RSA_4096
func GenerateRSAPrivateKey(
	organization string,
	certificates []appmodels.ISerialNumber,
	keyType appmodels.KeyType,
) (appmodels.IPrivateKey, error) {
	return GeneratePrivateKey(organization, certificates, keyType)
}

// GenerateECDSAPrivateKey creates a new private key of type models.KeyType
func GenerateECDSAPrivateKey(
	organization string,
	certificates []appmodels.ISerialNumber,
	keyType appmodels.KeyType,
) (appmodels.IPrivateKey, error) {
	return GeneratePrivateKey(organization, certificates, keyType)
}

// GenerateEd25519PrivateKey creates a new private key of type models.Ed25519
func GenerateEd25519PrivateKey(
	organization string,
	certificates []appmodels.ISerialNumber,
) (appmodels.IPrivateKey, error) {
	return GeneratePrivateKey(organization, certificates, appmodels.Ed25519)
}

// MarshalPrivateKeyAsPEM converts a private key to PEM data bytes
func MarshalPrivateKeyAsPEM(
	manager managers.ICertificateManager,
	data any,
) ([]byte, error) {
	var pemBlock *pem.Block

	switch typedKey := data.(type) {

	case *rsa.PrivateKey:
		pemBlock = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: manager.MarshalPKCS1PrivateKey(typedKey)}

	case *ecdsa.PrivateKey:
		ecBytes, err := manager.MarshalECPrivateKey(typedKey)
		if err != nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: failed to marshal ECDSA private key: %w", err)
		}
		pemBlock = &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecBytes}

	case ed25519.PrivateKey:
		edBytes, err := manager.MarshalPKCS8PrivateKey(typedKey)
		if err != nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: failed to marshal Ed25519 private key to PKCS#8: %w", err)
		}
		pemBlock = &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: edBytes,
		}

	default:
		return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: unsupported private key type")
	}

	pemData := pem.EncodeToMemory(pemBlock)

	return pemData, nil
}
