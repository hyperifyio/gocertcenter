// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

func TestSlugify(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Basic", "Hello World", "hello-world"},
		{"Prefix special characters", "-Hello World", "hello-world"},
		{"Suffix special characters", "Hello World-", "hello-world"},
		{"Directory characters", "Hello/..//../World-", "hello-world"},
		{"Special Characters", "Hello, World!", "hello-world"},
		{"Multiple Spaces", "Hello   World", "hello-world"},
		{"Uppercase and Special", "HELLO @ WORLD", "hello-world"},
		{"Numbers and Hyphens", "123 - Hello - 456", "123-hello-456"},
		{"Edge Case: Only Special Characters", "@!#$%", ""},
		{"Non-ASCII Characters", "Café ÉtéöäåÖÄÅ", "cafe-eteoaaoaa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := apputils.Slugify(tt.input)
			if got != tt.expected {
				t.Errorf("Slugify(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
