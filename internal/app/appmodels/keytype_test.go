// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestKeyType_String(t *testing.T) {
	tests := []struct {
		keyType appmodels.KeyType
		want    string
	}{
		{appmodels.RSA_1024, "RSA_1024"},
		{appmodels.RSA_2048, "RSA_2048"},
		{appmodels.RSA_3072, "RSA_3072"},
		{appmodels.RSA_4096, "RSA_4096"},
		{appmodels.ECDSA_P224, "ECDSA_P224"},
		{appmodels.ECDSA_P256, "ECDSA_P256"},
		{appmodels.ECDSA_P384, "ECDSA_P384"},
		{appmodels.ECDSA_P521, "ECDSA_P521"},
		{appmodels.Ed25519, "Ed25519"},
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
	if !appmodels.RSA_1024.IsRSA() {
		t.Errorf("RSA_1024.IsRSA() = false, want true")
	}

	if !appmodels.RSA_2048.IsRSA() {
		t.Errorf("RSA_2048.IsRSA() = false, want true")
	}

	if !appmodels.RSA_3072.IsRSA() {
		t.Errorf("RSA_3072.IsRSA() = false, want true")
	}

	if !appmodels.RSA_4096.IsRSA() {
		t.Errorf("RSA_4096.IsRSA() = false, want true")
	}

	if appmodels.ECDSA_P256.IsRSA() {
		t.Errorf("ECDSA_P256.IsRSA() = true, want false")
	}
}

func TestKeyType_IsECDSA(t *testing.T) {
	tests := []struct {
		keyType appmodels.KeyType
		want    bool
	}{
		{appmodels.RSA_1024, false},
		{appmodels.RSA_2048, false},
		{appmodels.RSA_3072, false},
		{appmodels.RSA_4096, false},
		{appmodels.ECDSA_P224, true},
		{appmodels.ECDSA_P256, true},
		{appmodels.ECDSA_P384, true},
		{appmodels.ECDSA_P521, true},
		{appmodels.Ed25519, false},
	}

	for _, tt := range tests {
		if got := tt.keyType.IsECDSA(); got != tt.want {
			t.Errorf("%v.IsECDSA() = %v, want %v", tt.keyType, got, tt.want)
		}
	}
}

func TestKeyType_IsEd25519(t *testing.T) {
	if !appmodels.Ed25519.IsEd25519() {
		t.Errorf("Ed25519.IsEd25519() = false, want true")
	}
	if appmodels.RSA_1024.IsEd25519() {
		t.Errorf("RSA_1024.IsEd25519() = true, want false")
	}
	if appmodels.RSA_2048.IsEd25519() {
		t.Errorf("RSA_2048.IsEd25519() = true, want false")
	}
	if appmodels.RSA_3072.IsEd25519() {
		t.Errorf("RSA_3072.IsEd25519() = true, want false")
	}
	if appmodels.RSA_4096.IsEd25519() {
		t.Errorf("RSA_4096.IsEd25519() = true, want false")
	}
}

func TestKeyType_String_DefaultCase(t *testing.T) {
	// Define a KeyType value that is outside the range of defined constants
	undefinedKeyType := appmodels.KeyType(999)
	expected := "KeyType(999)"

	// Call the String method on the undefined KeyType
	result := undefinedKeyType.String()

	// Assert that the result matches the expected output for the default case
	if result != expected {
		t.Errorf("Unexpected default string for undefined KeyType. Got: %v, Want: %v", result, expected)
	}
}
