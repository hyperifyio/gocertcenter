// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos_test

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"reflect"
	"testing"
)

func TestNewOrganizationDTO(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		id       string
		orgName  string
		allNames []string
		want     dtos.OrganizationDTO
	}{
		{
			name:     "Single name",
			id:       "org1",
			orgName:  "Organization One",
			allNames: []string{"Organization One"},
			want: dtos.OrganizationDTO{
				ID:       "org1",
				Name:     "Organization One",
				AllNames: []string{"Organization One"},
			},
		},
		{
			name:     "Multiple names",
			id:       "org2",
			orgName:  "Organization Two",
			allNames: []string{"Organization Two", "Org 2"},
			want: dtos.OrganizationDTO{
				ID:       "org2",
				Name:     "Organization Two",
				AllNames: []string{"Organization Two", "Org 2"},
			},
		},
		// Add more test cases as needed
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dtos.NewOrganizationDTO(tt.id, tt.orgName, tt.allNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrganizationDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
