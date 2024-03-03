// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"reflect"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewOrganizationDTO(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string // Test case name
		id       string
		slug     string
		orgName  string
		allNames []string
		want     appdtos.OrganizationDTO
	}{
		{
			name:     "Single name",
			id:       "1001",
			slug:     "org1",
			orgName:  "Organization One",
			allNames: []string{"Organization One"},
			want: appdtos.OrganizationDTO{
				ID:       "1001",
				Slug:     "org1",
				Name:     "Organization One",
				AllNames: []string{"Organization One"},
			},
		},
		{
			name:     "Multiple names",
			id:       "1002",
			slug:     "org2",
			orgName:  "Organization Two",
			allNames: []string{"Organization Two", "Org 2"},
			want: appdtos.OrganizationDTO{
				ID:       "1002",
				Slug:     "org2",
				Name:     "Organization Two",
				AllNames: []string{"Organization Two", "Org 2"},
			},
		},
		// Add more test cases as needed
	}

	// Execute tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appdtos.NewOrganizationDTO(tt.id, tt.slug, tt.orgName, tt.allNames)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrganizationDTO() = %v, want %v", got, tt.want)
			}
		})
	}
}
