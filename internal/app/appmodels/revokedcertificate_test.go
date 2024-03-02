// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package appmodels_test

import (
	"testing"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestNewRevokedCertificate(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(12345)
	revocationTime := time.Now().Round(time.Second) // Use Round to normalize to seconds for comparison
	expirationTime := time.Now().Add(365 * 24 * time.Hour).Round(time.Second)

	revokedCert := appmodels.NewRevokedCertificate(serialNumber, revocationTime, expirationTime)

	if revokedCert.SerialNumber().Cmp(serialNumber) != 0 {
		t.Errorf("Serial number mismatch, got %v, want %v", revokedCert.SerialNumber(), serialNumber)
	}

	if !revokedCert.RevocationTime().Equal(revocationTime) {
		t.Errorf("Revocation time mismatch, got %v, want %v", revokedCert.RevocationTime(), revocationTime)
	}

	if !revokedCert.ExpirationTime().Equal(expirationTime) {
		t.Errorf("Expiration time mismatch, got %v, want %v", revokedCert.ExpirationTime(), expirationTime)
	}
}

func TestGetRevokedCertificate(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(12345)
	revocationTime := time.Now().Round(time.Second) // Use Round to normalize to seconds for comparison

	revokedCert := appmodels.NewRevokedCertificate(serialNumber, revocationTime, time.Time{})

	pkixRevokedCert := revokedCert.RevokedCertificate()

	if pkixRevokedCert.SerialNumber.Cmp(serialNumber) != 0 {
		t.Errorf("PKIX Serial number mismatch, got %v, want %v", pkixRevokedCert.SerialNumber, serialNumber)
	}

	if !pkixRevokedCert.RevocationTime.Equal(revocationTime) {
		t.Errorf("PKIX Revocation time mismatch, got %v, want %v", pkixRevokedCert.RevocationTime, revocationTime)
	}
}
