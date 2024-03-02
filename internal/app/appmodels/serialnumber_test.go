// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"math/big"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestSerialNumberToBigInt(t *testing.T) {
	// Setup: Create a new serial number directly for testing purposes
	expectedBigInt := big.NewInt(1234567890)
	serialNumber := appmodels.NewSerialNumber(1234567890)

	// Verify: The result should be equal to the initial *big.Int used to create Int64SerialNumber
	if serialNumber.Cmp(expectedBigInt) != 0 {
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
			// Given a Int64SerialNumber with a specific value
			serialNumber := appmodels.NewSerialNumber(tt.input)

			// When calling Sign()
			sign := serialNumber.Sign()

			// Then the sign should be correctly identified
			if sign != tt.expected {
				t.Errorf("Expected Sign() to return %d for input %d, got %d", tt.expected, tt.input, sign)
			}
		})
	}
}

func TestParseSerialNumber(t *testing.T) {
	tests := []struct {
		name      string
		value     string
		base      int
		expectErr bool
	}{
		{"Valid Hex", "1A3B5C7D", 16, false},
		{"Valid Decimal", "1234567890", 10, false},
		{"Invalid Hex", "ZZZ", 16, true},
		{"Empty Value", "", 10, true},
		{"Invalid Base", "123", 1, true}, // Base cannot be less than 2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := appmodels.ParseSerialNumber(tt.value, tt.base)
			if tt.expectErr {
				if err == nil {
					t.Errorf("ParseSerialNumber(%q, %d) expected an error, but got nil", tt.value, tt.base)
				}
			} else {
				if err != nil {
					t.Errorf("ParseSerialNumber(%q, %d) unexpected error: %v", tt.value, tt.base, err)
				}
			}
		})
	}
}
