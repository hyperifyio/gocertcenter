// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestNewOrganization(t *testing.T) {
	orgID := "org123"
	names := []string{"Test Org", "Test Org Department"}
	org := appmodels.NewOrganization(orgID, names)

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
	org := appmodels.NewOrganization(orgID, nil)

	if got := org.GetID(); got != orgID {
		t.Errorf("GetID() = %s, want = %s", got, orgID)
	}
}

func TestOrganization_GetName(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := appmodels.NewOrganization("org789", names)

	if got := org.GetName(); got != names[0] {
		t.Errorf("GetName() = %s, want = %s", got, names[0])
	}
}

func TestOrganization_GetName_NoNames(t *testing.T) {
	org := appmodels.NewOrganization("orgNoNames", []string{})
	if name := org.GetName(); name != "" {
		t.Errorf("GetName() with no names should return an empty string, got: %s", name)
	}
}

func TestOrganization_GetNames(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := appmodels.NewOrganization("org101112", names)

	gotNames := org.GetNames()
	if len(gotNames) != len(names) || gotNames[0] != names[0] || gotNames[1] != names[1] {
		t.Errorf("GetNames() got = %v, want = %v", gotNames, names)
	}
}
