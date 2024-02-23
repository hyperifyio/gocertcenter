// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"testing"
)

func TestNewIndexDTO(t *testing.T) {
	tests := []struct {
		version string
		want    dtos.IndexDTO
	}{
		{"1.0.0", dtos.IndexDTO{Version: "1.0.0"}},
		{"2.0.1", dtos.IndexDTO{Version: "2.0.1"}},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			got := dtos.NewIndexDTO(tt.version)
			if got.Version != tt.want.Version {
				t.Errorf("NewIndexDTO(%s) = %v, want %v", tt.version, got, tt.want)
			}
		})
	}
}
