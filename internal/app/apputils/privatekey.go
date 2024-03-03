// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func ToPrivateKeyDTO(certManager managers.CertificateManager, c appmodels.PrivateKey) (appdtos.PrivateKeyDTO, error) {

	bytes, err := MarshalPrivateKeyAsPEM(certManager, c.PrivateKey())
	if err != nil {
		return appdtos.PrivateKeyDTO{}, fmt.Errorf("ToPrivateKeyDTO: rsa: failed to marshal RSA private key: %v", err)
	}

	return appdtos.NewPrivateKeyDTO(
		c.SerialNumber().String(),
		c.KeyType().String(),
		string(bytes),
	), nil
}

func ToPrivateKeyDTOList(certManager managers.CertificateManager, list []appmodels.PrivateKey) ([]appdtos.PrivateKeyDTO, error) {
	result := make([]appdtos.PrivateKeyDTO, len(list))
	for i, v := range list {
		dto, err := ToPrivateKeyDTO(certManager, v)
		if err != nil {
			return nil, fmt.Errorf("ToPrivateKeyDTOList: failed: %w", err)
		}
		result[i] = dto
	}
	return result, nil
}

// GeneratePrivateKey creates a new private key of type models.KeyType
//   - organization: The organization ID
//   - certificates: The certificate signatures from root certificate to the one owning this private key
//   - keyType: The key type to generate
func GeneratePrivateKey(
	organization *big.Int,
	certificate *big.Int,
	keyType appmodels.KeyType,
) (appmodels.PrivateKey, error) {

	if organization == nil {
		return nil, fmt.Errorf("GeneratePrivateKey: organization: must not be empty")
	}

	if certificate == nil {
		return nil, fmt.Errorf("GeneratePrivateKey: certificate: must have a serial number")
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
	return appmodels.NewPrivateKey(organization, certificate, keyType, key), nil
}

// GenerateRSAPrivateKey creates a new RSA private key
//   - organization: Organization ID
//   - certificates: Certificate serial numbers from root certificate to the one owning this key
//   - keyType: Should be appmodels.RSA_1024, appmodels.RSA_2048, appmodels.RSA_3072 or appmodels.RSA_4096
func GenerateRSAPrivateKey(
	organization *big.Int,
	certificate *big.Int,
	keyType appmodels.KeyType,
) (appmodels.PrivateKey, error) {
	return GeneratePrivateKey(organization, certificate, keyType)
}

// GenerateECDSAPrivateKey creates a new private key of type models.KeyType
func GenerateECDSAPrivateKey(
	organization *big.Int,
	certificate *big.Int,
	keyType appmodels.KeyType,
) (appmodels.PrivateKey, error) {
	return GeneratePrivateKey(organization, certificate, keyType)
}

// GenerateEd25519PrivateKey creates a new private key of type models.Ed25519
func GenerateEd25519PrivateKey(
	organization *big.Int,
	certificate *big.Int,
) (appmodels.PrivateKey, error) {
	return GeneratePrivateKey(organization, certificate, appmodels.Ed25519)
}

// MarshalPrivateKeyAsPEM converts a private key to PEM data bytes
func MarshalPrivateKeyAsPEM(
	manager managers.CertificateManager,
	data any,
) ([]byte, error) {
	var pemBlock *pem.Block

	switch typedKey := data.(type) {

	case *rsa.PrivateKey:
		rsaBytes := manager.MarshalPKCS1PrivateKey(typedKey)
		if rsaBytes == nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: rsa: failed to marshal RSA private key: got nil")
		}
		pemBlock = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: rsaBytes}

	case *ecdsa.PrivateKey:
		ecBytes, err := manager.MarshalECPrivateKey(typedKey)
		if err != nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: ecdsa: failed to marshal ECDSA private key: %w", err)
		}
		if ecBytes == nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: ecdsa: failed to marshal ECDSA private key: got nil")
		}
		pemBlock = &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecBytes}

	case ed25519.PrivateKey:
		edBytes, err := manager.MarshalPKCS8PrivateKey(typedKey)
		if err != nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: ed25519: failed to marshal Ed25519 private key to PKCS#8: %w", err)
		}
		if edBytes == nil {
			return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: ed25519: failed to marshal ECDSA private key: got nil")
		}
		pemBlock = &pem.Block{Type: "PRIVATE KEY", Bytes: edBytes}

	default:
		return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: unsupported private key type")
	}

	pemData := manager.EncodePEMToMemory(pemBlock)
	if pemData == nil {
		return nil, fmt.Errorf("MarshalPrivateKeyAsPEM: could not encode to PEM")
	}
	return pemData, nil
}

func DetermineRSATypeFromSize(keySize int) (appmodels.KeyType, error) {
	switch keySize {
	case 1024:
		return appmodels.RSA_1024, nil
	case 2048:
		return appmodels.RSA_2048, nil
	case 3072:
		return appmodels.RSA_3072, nil
	case 4096:
		return appmodels.RSA_4096, nil
	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("DetermineRSATypeFromSize: RSA key size not supported: %d", keySize)
	}
}

func DetermineECDSACurve(curve elliptic.Curve) (appmodels.KeyType, error) {
	switch curve {
	case elliptic.P224():
		return appmodels.ECDSA_P224, nil
	case elliptic.P256():
		return appmodels.ECDSA_P256, nil
	case elliptic.P384():
		return appmodels.ECDSA_P384, nil
	case elliptic.P521():
		return appmodels.ECDSA_P521, nil
	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("DetermineECDSACurve: unsupported EC curve")
	}
}

func DetermineRSATypeFromKey(key *rsa.PrivateKey) (appmodels.KeyType, error) {
	keySize := ReadRSAKeySize(key)
	keyType, err := DetermineRSATypeFromSize(keySize)
	if err != nil {
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("DetermineRSATypeFromKey: failed: %w", err)
	}
	return keyType, nil
}

func DetermineKeyType(privateKey any) (appmodels.KeyType, error) {
	switch key := privateKey.(type) {

	case *rsa.PrivateKey:
		keyType, err := DetermineRSATypeFromKey(key)
		if err != nil {
			return appmodels.NIL_KEY_TYPE, fmt.Errorf("DetermineKeyType: could not detect RSA key type: %w", err)
		}
		return keyType, nil

	case *ecdsa.PrivateKey:
		keyType, err := DetermineECDSACurve(key.Curve)
		if err != nil {
			return appmodels.NIL_KEY_TYPE, fmt.Errorf("DetermineKeyType: could not detect ecdsa key type: %w", err)
		}
		return keyType, nil

	case ed25519.PrivateKey:
		return appmodels.Ed25519, nil

	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("DetermineKeyType: unknown or unsupported key type")
	}

}

func ReadRSAKeySize(key *rsa.PrivateKey) int {
	if key == nil {
		return 0
	}
	if key.N == nil {
		return 0
	}
	return key.N.BitLen()
}

func ParsePrivateKeyFromPEMBytes(
	certManager managers.CertificateManager,
	data []byte,
) (any, appmodels.KeyType, error) {
	block, _ := certManager.DecodePEM(data)
	if block == nil {
		return nil, appmodels.NIL_KEY_TYPE, errors.New("ParsePrivateKeyFromPEMBytes: failed to decode PEM block containing the private key")
	}
	return ParsePrivateKeyFromPEMBlock(certManager, block)
}

func ParsePrivateKeyFromPEMBlock(
	certManager managers.CertificateManager,
	block *pem.Block,
) (any, appmodels.KeyType, error) {
	if block.Type == "PRIVATE KEY" {
		return ParsePKCS8PrivateKey(certManager, block.Bytes)
	} else if block.Type == "RSA PRIVATE KEY" {
		return ParseRSAPrivateKey(certManager, block.Bytes)
	} else if block.Type == "EC PRIVATE KEY" {
		return ParseECPrivateKey(certManager, block.Bytes)
	}
	return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("ParsePrivateKeyFromPEMBlock: unsupported block type: %s", block.Type)
}

func ParsePKCS8PrivateKey(certManager managers.CertificateManager, der []byte) (any, appmodels.KeyType, error) {
	privateKey, err := certManager.ParsePKCS8PrivateKey(der)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("failed to parse private key: %w", err)
	}
	keyType, err := DetermineKeyType(privateKey)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("could not detect key type: %w", err)
	}
	return privateKey, keyType, nil
}

func ParseRSAPrivateKey(certManager managers.CertificateManager, bytes []byte) (any, appmodels.KeyType, error) {
	privateKey, err := certManager.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("ParseRSAPrivateKey: failed to parse RSA private key: %w", err)
	}
	keyType, err := DetermineRSATypeFromKey(privateKey)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("ParseRSAPrivateKey: failed to parse RSA type: %w", err)
	}
	return privateKey, keyType, nil
}

func ParseECPrivateKey(certManager managers.CertificateManager, bytes []byte) (any, appmodels.KeyType, error) {
	privateKey, err := certManager.ParseECPrivateKey(bytes)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("failed to parse EC private key: %w", err)
	}
	// Determine the curve
	keyType, err := DetermineECDSACurve(privateKey.Curve)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("could not detect key type: %w", err)
	}
	return privateKey, keyType, nil
}
