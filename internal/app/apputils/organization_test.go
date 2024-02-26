// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package apputils_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

func TestGetOrganizationDTO(t *testing.T) {
	orgID := "org123"
	names := []string{"Test Org", "Test Org Department"}
	org := appmodels.NewOrganization(orgID, names)

	dto := apputils.GetOrganizationDTO(org)

	// Verify ID
	if dto.ID != orgID {
		t.Errorf("GetOrganizationDTO().ID = %s, want %s", dto.ID, orgID)
	}

	// Verify Name
	expectedName := names[0] // GetName returns the first name in the slice
	if dto.Name != expectedName {
		t.Errorf("GetOrganizationDTO().Name = %s, want %s", dto.Name, expectedName)
	}

	// Verify AllNames
	if len(dto.AllNames) != len(names) {
		t.Fatalf("GetOrganizationDTO().AllNames returned %d names; want %d", len(dto.AllNames), len(names))
	}

	for i, name := range dto.AllNames {
		if name != names[i] {
			t.Errorf("GetOrganizationDTO().AllNames[%d] = %s, want %s", i, name, names[i])
		}
	}
}
