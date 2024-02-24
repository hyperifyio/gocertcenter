// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/managers"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"math/big"
	"testing"
)

// TestNewSerialNumber checks if GenerateSerialNumber returns a non-nil, positive
// serial number and does not return an error.
func TestNewSerialNumber(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, err := models.GenerateSerialNumber(randomManager)
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

// TestNewSerialNumberUniqueness checks the uniqueness of generated serial numbers.
// Note: This test does not guarantee uniqueness but reduces the likelihood of collisions.
func TestNewSerialNumberUniqueness(t *testing.T) {
	randomManager := managers.NewRandomManager()
	seen := make(map[string]bool)
	count := 100 // Number of serial numbers to generate for the test
	for i := 0; i < count; i++ {
		serialNumber, err := models.GenerateSerialNumber(randomManager)
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

func TestNewSerialNumberWithMock(t *testing.T) {
	mockRandomManager := &mocks.MockRandomManager{
		CreateBigIntFunc: func(max *big.Int) (*big.Int, error) {
			return big.NewInt(12345), nil // Return a fixed serial number for testing
		},
	}

	serialNumber, err := models.GenerateSerialNumber(mockRandomManager)
	if err != nil {
		t.Fatalf("GenerateSerialNumber failed: %v", err)
	}

	expected := big.NewInt(12345)
	if serialNumber.Value().Cmp(expected) != 0 {
		t.Errorf("Expected serial number %v, got %v", expected, serialNumber)
	}
}

func TestNewSerialNumberWithErrors(t *testing.T) {
	mockRandomManager := &mocks.MockRandomManager{
		CreateBigIntFunc: func(max *big.Int) (*big.Int, error) {
			return nil, errors.New("Mocked error")
		},
	}

	_, err := models.GenerateSerialNumber(mockRandomManager)
	if err == nil {
		t.Fatalf("Expected GenerateSerialNumber() to fail, did not fail")
	}
}

func TestSerialNumberToBigInt(t *testing.T) {
	// Setup: Create a new serial number directly for testing purposes
	expectedBigInt := big.NewInt(1234567890) // Example serial number
	serialNumber := models.NewSerialNumber(expectedBigInt)

	// Verify: The result should be equal to the initial *big.Int used to create SerialNumber
	if serialNumber.Value().Cmp(expectedBigInt) != 0 {
		t.Errorf("SerialNumberToBigInt did not return the expected *big.Int value. Expected: %s, Got: %s", expectedBigInt.String(), serialNumber.String())
	}
}
