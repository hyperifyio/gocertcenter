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

// TestNewSerialNumber checks if NewSerialNumber returns a non-nil, positive
// serial number and does not return an error.
func TestNewSerialNumber(t *testing.T) {
	randomManager := managers.NewRandomManager()
	serialNumber, err := models.NewSerialNumber(randomManager)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if serialNumber == nil {
		t.Fatalf("Expected serial number to be non-nil")
	}
	bigIntSerialNumber := (*big.Int)(serialNumber)
	if bigIntSerialNumber.Sign() != 1 {
		t.Fatalf("Expected serial number to be positive, got %s", bigIntSerialNumber)
	}
}

// TestNewSerialNumberUniqueness checks the uniqueness of generated serial numbers.
// Note: This test does not guarantee uniqueness but reduces the likelihood of collisions.
func TestNewSerialNumberUniqueness(t *testing.T) {
	randomManager := managers.NewRandomManager()
	seen := make(map[string]bool)
	count := 100 // Number of serial numbers to generate for the test
	for i := 0; i < count; i++ {
		serialNumber, err := models.NewSerialNumber(randomManager)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		bigIntSerialNumber := (*big.Int)(serialNumber)
		serialStr := bigIntSerialNumber.String()
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

	serialNumber, err := models.NewSerialNumber(mockRandomManager)
	if err != nil {
		t.Fatalf("NewSerialNumber failed: %v", err)
	}

	bigIntSSerialNumber := (*big.Int)(serialNumber)

	expected := big.NewInt(12345)
	if bigIntSSerialNumber.Cmp(expected) != 0 {
		t.Errorf("Expected serial number %v, got %v", expected, serialNumber)
	}
}

func TestNewSerialNumberWithErrors(t *testing.T) {
	mockRandomManager := &mocks.MockRandomManager{
		CreateBigIntFunc: func(max *big.Int) (*big.Int, error) {
			return nil, errors.New("Mocked error")
		},
	}

	_, err := models.NewSerialNumber(mockRandomManager)
	if err == nil {
		t.Fatalf("Expected NewSerialNumber() to fail, did not fail")
	}
}

func TestSerialNumberToBigInt(t *testing.T) {
	// Setup: Create a new serial number directly for testing purposes
	expectedBigInt := big.NewInt(1234567890) // Example serial number
	var serialNumber models.SerialNumber = expectedBigInt

	// Execute: Convert SerialNumber to *big.Int
	resultBigInt := models.SerialNumberToBigInt(serialNumber)

	// Verify: The result should be equal to the initial *big.Int used to create SerialNumber
	if resultBigInt.Cmp(expectedBigInt) != 0 {
		t.Errorf("SerialNumberToBigInt did not return the expected *big.Int value. Expected: %v, Got: %v", expectedBigInt, resultBigInt)
	}
}
