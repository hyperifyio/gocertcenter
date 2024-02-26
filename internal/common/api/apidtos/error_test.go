// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apidtos_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apidtos"
)

func TestNewErrorDTO(t *testing.T) {
	tests := []struct {
		code  int
		error string
		want  apidtos.ErrorDTO
	}{
		{404, "Not Found", apidtos.ErrorDTO{404, "Not Found"}},
		{500, "Internal Server Error", apidtos.ErrorDTO{500, "Internal Server Error"}},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.error, func(t *testing.T) {
			got := apidtos.NewErrorDTO(tt.code, tt.error)
			if got.Code != tt.want.Code || got.Error != tt.want.Error {
				t.Errorf("NewErrorDTO(%d, %s) = %v, want %v", tt.code, tt.error, got, tt.want)
			}
		})
	}
}
