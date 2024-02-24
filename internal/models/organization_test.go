// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models_test

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	orgID := "org123"
	names := []string{"Test Org", "Test Org Department"}
	org := models.NewOrganization(orgID, names)

	if org.GetID() != orgID {
		t.Errorf("GetID() = %s, want %s", org.GetID(), orgID)
	}

	if len(org.GetNames()) != len(names) {
		t.Fatalf("GetNames() returned %d names; want %d", len(org.GetNames()), len(names))
	}

	for i, name := range org.GetNames() {
		if name != names[i] {
			t.Errorf("GetNames()[%d] = %s, want %s", i, name, names[i])
		}
	}
}

func TestOrganization_GetID(t *testing.T) {
	orgID := "org456"
	org := models.NewOrganization(orgID, nil)

	if got := org.GetID(); got != orgID {
		t.Errorf("GetID() = %s, want = %s", got, orgID)
	}
}

func TestOrganization_GetName(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := models.NewOrganization("org789", names)

	if got := org.GetName(); got != names[0] {
		t.Errorf("GetName() = %s, want = %s", got, names[0])
	}
}

func TestOrganization_GetName_NoNames(t *testing.T) {
	org := models.NewOrganization("orgNoNames", []string{})
	if name := org.GetName(); name != "" {
		t.Errorf("GetName() with no names should return an empty string, got: %s", name)
	}
}

func TestOrganization_GetNames(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := models.NewOrganization("org101112", names)

	gotNames := org.GetNames()
	if len(gotNames) != len(names) || gotNames[0] != names[0] || gotNames[1] != names[1] {
		t.Errorf("GetNames() got = %v, want = %v", gotNames, names)
	}
}

func TestOrganization_GetDTO(t *testing.T) {
	orgID := "org123"
	names := []string{"Test Org", "Test Org Department"}
	org := models.NewOrganization(orgID, names)

	dto := org.GetDTO()

	// Verify ID
	if dto.ID != orgID {
		t.Errorf("GetDTO().ID = %s, want %s", dto.ID, orgID)
	}

	// Verify Name
	expectedName := names[0] // GetName returns the first name in the slice
	if dto.Name != expectedName {
		t.Errorf("GetDTO().Name = %s, want %s", dto.Name, expectedName)
	}

	// Verify AllNames
	if len(dto.AllNames) != len(names) {
		t.Fatalf("GetDTO().AllNames returned %d names; want %d", len(dto.AllNames), len(names))
	}

	for i, name := range dto.AllNames {
		if name != names[i] {
			t.Errorf("GetDTO().AllNames[%d] = %s, want %s", i, name, names[i])
		}
	}
}
