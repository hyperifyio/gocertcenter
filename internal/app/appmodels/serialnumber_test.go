// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"math/big"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestSerialNumberToBigInt(t *testing.T) {
	// Setup: Create a new serial number directly for testing purposes
	expectedBigInt := big.NewInt(1234567890) // Example serial number
	serialNumber := appmodels.NewSerialNumber(expectedBigInt)

	// Verify: The result should be equal to the initial *big.Int used to create SerialNumber
	if serialNumber.Value().Cmp(expectedBigInt) != 0 {
		t.Errorf("SerialNumberToBigInt did not return the expected *big.Int value. Expected: %s, Got: %s", expectedBigInt.String(), serialNumber.String())
	}
}

func TestSerialNumber_Sign(t *testing.T) {
	tests := []struct {
		name     string
		input    int64 // Using int64 for simplicity, as NewInt can directly use it
		expected int
	}{
		{"Positive Value", 123, 1},
		{"Negative Value", -123, -1},
		{"Zero Value", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given a SerialNumber with a specific value
			serialNumber := appmodels.NewSerialNumber(big.NewInt(tt.input))

			// When calling Sign()
			sign := serialNumber.Sign()

			// Then the sign should be correctly identified
			if sign != tt.expected {
				t.Errorf("Expected Sign() to return %d for input %d, got %d", tt.expected, tt.input, sign)
			}
		})
	}
}
