// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package modelutils_test

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelutils"
	"math/big"
	"testing"
)

func TestNewSerialNumberWithErrors(t *testing.T) {
	mockRandomManager := &mocks.MockRandomManager{
		CreateBigIntFunc: func(max *big.Int) (*big.Int, error) {
			return nil, errors.New("Mocked error")
		},
	}

	_, err := modelutils.GenerateSerialNumber(mockRandomManager)
	if err == nil {
		t.Fatalf("Expected GenerateSerialNumber() to fail, did not fail")
	}
}

func TestNewSerialNumberWithMock(t *testing.T) {
	mockRandomManager := &mocks.MockRandomManager{
		CreateBigIntFunc: func(max *big.Int) (*big.Int, error) {
			return big.NewInt(12345), nil // Return a fixed serial number for testing
		},
	}

	serialNumber, err := modelutils.GenerateSerialNumber(mockRandomManager)
	if err != nil {
		t.Fatalf("GenerateSerialNumber failed: %v", err)
	}

	expected := big.NewInt(12345)
	if serialNumber.Value().Cmp(expected) != 0 {
		t.Errorf("Expected serial number %v, got %v", expected, serialNumber)
	}
}

// TestNewSerialNumberUniqueness checks the uniqueness of generated serial numbers.
// Note: This test does not guarantee uniqueness but reduces the likelihood of collisions.
func TestNewSerialNumberUniqueness(t *testing.T) {
	randomManager := managers.NewRandomManager()
	seen := make(map[string]bool)
	count := 100 // Number of serial numbers to generate for the test
	for i := 0; i < count; i++ {
		serialNumber, err := modelutils.GenerateSerialNumber(randomManager)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		serialStr := serialNumber.String()
		if seen[serialStr] {
			t.Fatalf("Duplicate serial number found: %s", serialStr)
		}
		seen[serialStr] = true
	}
}

// TestNewSerialNumber checks if GenerateSerialNumber returns a non-nil, positive
// serial number and does not return an error.
func TestNewSerialNumber(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, err := modelutils.GenerateSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if serialNumber == nil {
		t.Fatalf("Expected serial number to be non-nil")
	}
	if serialNumber.Sign() != 1 {
		t.Fatalf("Expected serial number to be positive, got %s", serialNumber.String())
	}
}
