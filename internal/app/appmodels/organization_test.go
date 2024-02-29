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

	if org.ID() != orgID {
		t.Errorf("ID() = %s, want %s", org.ID(), orgID)
	}

	if len(org.Names()) != len(names) {
		t.Fatalf("Names() returned %d names; want %d", len(org.Names()), len(names))
	}

	for i, name := range org.Names() {
		if name != names[i] {
			t.Errorf("Names()[%d] = %s, want %s", i, name, names[i])
		}
	}
}

func TestOrganization_GetID(t *testing.T) {
	orgID := "org456"
	org := appmodels.NewOrganization(orgID, nil)

	if got := org.ID(); got != orgID {
		t.Errorf("ID() = %s, want = %s", got, orgID)
	}
}

func TestOrganization_GetName(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := appmodels.NewOrganization("org789", names)

	if got := org.Name(); got != names[0] {
		t.Errorf("Name() = %s, want = %s", got, names[0])
	}
}

func TestOrganization_GetName_NoNames(t *testing.T) {
	org := appmodels.NewOrganization("orgNoNames", []string{})
	if name := org.Name(); name != "" {
		t.Errorf("Name() with no names should return an empty string, got: %s", name)
	}
}

func TestOrganization_GetNames(t *testing.T) {
	names := []string{"Primary Name", "Secondary Name"}
	org := appmodels.NewOrganization("org101112", names)

	gotNames := org.Names()
	if len(gotNames) != len(names) || gotNames[0] != names[0] || gotNames[1] != names[1] {
		t.Errorf("Names() got = %v, want = %v", gotNames, names)
	}
}
