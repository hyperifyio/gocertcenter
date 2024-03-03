// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels_test

import (
	"math/big"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func TestNewOrganization(t *testing.T) {
	orgID := big.NewInt(123)
	orgSlug := "org789"
	names := []string{"Test Org", "Test Org Department"}
	org := appmodels.NewOrganization(orgID, orgSlug, names)

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

func TestOrganization_ID(t *testing.T) {
	orgID := big.NewInt(1)
	orgSlug := "org456"
	org := appmodels.NewOrganization(orgID, orgSlug, nil)

	if got := org.ID(); got != orgID {
		t.Errorf("ID() = %s, want = %s", got, orgID)
	}
}

func TestOrganization_Slug(t *testing.T) {
	orgID := big.NewInt(1)
	orgSlug := "org456"
	org := appmodels.NewOrganization(orgID, orgSlug, nil)

	if got := org.Slug(); got != orgSlug {
		t.Errorf("ID() = %s, want = %s", got, orgID)
	}
}

func TestOrganization_Name(t *testing.T) {
	orgID := big.NewInt(1)
	orgSlug := "org789"
	names := []string{"Primary Name", "Secondary Name"}
	org := appmodels.NewOrganization(orgID, orgSlug, names)

	if got := org.Name(); got != names[0] {
		t.Errorf("Name() = %s, want = %s", got, names[0])
	}
}

func TestOrganization_Name_NoNames(t *testing.T) {
	orgID := big.NewInt(1)
	orgSlug := "orgNoNames"
	org := appmodels.NewOrganization(orgID, orgSlug, []string{})
	if name := org.Name(); name != "" {
		t.Errorf("Name() with no names should return an empty string, got: %s", name)
	}
}

func TestOrganization_Names(t *testing.T) {
	orgID := big.NewInt(1)
	orgSlug := "org101112"
	names := []string{"Primary Name", "Secondary Name"}
	org := appmodels.NewOrganization(orgID, orgSlug, names)

	gotNames := org.Names()
	if len(gotNames) != len(names) || gotNames[0] != names[0] || gotNames[1] != names[1] {
		t.Errorf("Names() got = %v, want = %v", gotNames, names)
	}
}
