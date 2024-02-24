// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"testing"
)

func TestNewIndexDTO(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    dtos.IndexDTO
	}{
		{"gocertcenter", "1.0.0", dtos.IndexDTO{Name: "gocertcenter", Version: "1.0.0"}},
		{"gocertcenter", "2.0.1", dtos.IndexDTO{Name: "gocertcenter", Version: "2.0.1"}},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			got := dtos.NewIndexDTO(tt.name, tt.version)
			if got.Version != tt.want.Version {
				t.Errorf("NewIndexDTO(%s, %s) = %v, want %v", tt.name, tt.version, got, tt.want)
			}
		})
	}
}
