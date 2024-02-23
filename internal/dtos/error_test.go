// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"testing"
)

func TestNewErrorDTO(t *testing.T) {
	tests := []struct {
		code  int
		error string
		want  dtos.ErrorDTO
	}{
		{404, "Not Found", dtos.ErrorDTO{404, "Not Found"}},
		{500, "Internal Server Error", dtos.ErrorDTO{500, "Internal Server Error"}},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.error, func(t *testing.T) {
			got := dtos.NewErrorDTO(tt.code, tt.error)
			if got.Code != tt.want.Code || got.Error != tt.want.Error {
				t.Errorf("NewErrorDTO(%d, %s) = %v, want %v", tt.code, tt.error, got, tt.want)
			}
		})
	}
}
