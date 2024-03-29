// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import "fmt"

// KeyType represents the type of private key.
type KeyType int

const (
	// RSA represents an RSA private key.
	NIL_KEY_TYPE KeyType = iota

	// RSA_1024 represents an RSA private key with 1024 bits size.
	RSA_1024

	// RSA_2048 represents an RSA private key with 2048 bits size.
	RSA_2048

	// RSA_3072 represents an RSA private key with 3072 bits size.
	RSA_3072

	// RSA_4096 represents an RSA private key with 4096 bits size.
	RSA_4096

	// ECDSA_P224 represents an ECDSA private key with P224
	ECDSA_P224

	// ECDSA_P256 represents an ECDSA private key with P256
	ECDSA_P256

	// ECDSA_P384 represents an ECDSA private key with P384
	ECDSA_P384

	// ECDSA_P521 represents an ECDSA private key with P521
	ECDSA_P521

	// Ed25519 represents an Ed25519 private key.
	Ed25519
)

func (kt KeyType) String() string {
	switch kt {
	case RSA_1024:
		return "RSA_1024"
	case RSA_2048:
		return "RSA_2048"
	case RSA_3072:
		return "RSA_3072"
	case RSA_4096:
		return "RSA_4096"
	case ECDSA_P224:
		return "ECDSA_P224"
	case ECDSA_P256:
		return "ECDSA_P256"
	case ECDSA_P384:
		return "ECDSA_P384"
	case ECDSA_P521:
		return "ECDSA_P521"
	case Ed25519:
		return "Ed25519"
	default:
		return fmt.Sprintf("KeyType(%d)", kt)
	}
}

func (kt KeyType) IsRSA() bool {
	return kt == RSA_1024 || kt == RSA_2048 || kt == RSA_3072 || kt == RSA_4096
}

func (kt KeyType) IsECDSA() bool {
	return kt == ECDSA_P224 || kt == ECDSA_P256 || kt == ECDSA_P384 || kt == ECDSA_P521
}

func (kt KeyType) IsEd25519() bool {
	return kt == Ed25519
}
