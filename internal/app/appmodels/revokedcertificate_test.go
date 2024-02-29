// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmodels_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestNewRevokedCertificate(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(big.NewInt(12345))
	revocationTime := time.Now().Round(time.Second) // Use Round to normalize to seconds for comparison
	expirationTime := time.Now().Add(365 * 24 * time.Hour).Round(time.Second)

	revokedCert := appmodels.NewRevokedCertificate(serialNumber, revocationTime, expirationTime)

	if revokedCert.GetSerialNumber().Cmp(serialNumber) != 0 {
		t.Errorf("Serial number mismatch, got %v, want %v", revokedCert.GetSerialNumber(), serialNumber)
	}

	if !revokedCert.GetRevocationTime().Equal(revocationTime) {
		t.Errorf("Revocation time mismatch, got %v, want %v", revokedCert.GetRevocationTime(), revocationTime)
	}

	if !revokedCert.GetExpirationTime().Equal(expirationTime) {
		t.Errorf("Expiration time mismatch, got %v, want %v", revokedCert.GetExpirationTime(), expirationTime)
	}
}

func TestGetRevokedCertificate(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(big.NewInt(12345))
	revocationTime := time.Now().Round(time.Second) // Use Round to normalize to seconds for comparison

	revokedCert := appmodels.NewRevokedCertificate(serialNumber, revocationTime, time.Time{})

	pkixRevokedCert := revokedCert.GetRevokedCertificate()

	if pkixRevokedCert.SerialNumber.Cmp(serialNumber.Value()) != 0 {
		t.Errorf("PKIX Serial number mismatch, got %v, want %v", pkixRevokedCert.SerialNumber, serialNumber.Value())
	}

	if !pkixRevokedCert.RevocationTime.Equal(revocationTime) {
		t.Errorf("PKIX Revocation time mismatch, got %v, want %v", pkixRevokedCert.RevocationTime, revocationTime)
	}
}
