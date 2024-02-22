// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import "testing"

func TestKeyType_String(t *testing.T) {
	tests := []struct {
		keyType KeyType
		want    string
	}{
		{RSA, "RSA"},
		{ECDSA_P224, "ECDSA_P224"},
		{ECDSA_P256, "ECDSA_P256"},
		{ECDSA_P384, "ECDSA_P384"},
		{ECDSA_P521, "ECDSA_P521"},
		{Ed25519, "Ed25519"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.keyType.String(); got != tt.want {
				t.Errorf("KeyType.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyType_IsRSA(t *testing.T) {
	if !RSA.IsRSA() {
		t.Errorf("RSA.IsRSA() = false, want true")
	}

	if ECDSA_P256.IsRSA() {
		t.Errorf("ECDSA_P256.IsRSA() = true, want false")
	}
}

func TestKeyType_IsECDSA(t *testing.T) {
	tests := []struct {
		keyType KeyType
		want    bool
	}{
		{RSA, false},
		{ECDSA_P224, true},
		{ECDSA_P256, true},
		{ECDSA_P384, true},
		{ECDSA_P521, true},
		{Ed25519, false},
	}

	for _, tt := range tests {
		if got := tt.keyType.IsECDSA(); got != tt.want {
			t.Errorf("%v.IsECDSA() = %v, want %v", tt.keyType, got, tt.want)
		}
	}
}

func TestKeyType_IsEd25519(t *testing.T) {
	if !Ed25519.IsEd25519() {
		t.Errorf("Ed25519.IsEd25519() = false, want true")
	}

	if RSA.IsEd25519() {
		t.Errorf("RSA.IsEd25519() = true, want false")
	}
}
