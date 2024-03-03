// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

func TestParseBigInt(t *testing.T) {
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
			_, err := apputils.ParseBigInt(tt.value, tt.base)
			if tt.expectErr {
				if err == nil {
					t.Errorf("ParseBigInt(%q, %d) expected an error, but got nil", tt.value, tt.base)
				}
			} else {
				if err != nil {
					t.Errorf("ParseBigInt(%q, %d) unexpected error: %v", tt.value, tt.base, err)
				}
			}
		})
	}
}
