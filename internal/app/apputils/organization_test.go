// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
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

func TestToListOfOrganizationDTO(t *testing.T) {
	// Mock organization models
	orgID1 := "org123"
	name1 := "Test Org 1"
	names1 := []string{name1, "Test Org Department 1"}
	org1 := new(appmocks.MockOrganization)
	org1.On("GetID").Return(orgID1)
	org1.On("GetName").Return(name1)
	org1.On("GetNames").Return(names1)

	orgID2 := "org456"
	name2 := "Test Org 2"
	names2 := []string{name2, "Test Org Department 2"}
	org2 := new(appmocks.MockOrganization)
	org2.On("GetID").Return(orgID2)
	org2.On("GetName").Return(name2)
	org2.On("GetNames").Return(names2)

	orgList := []appmodels.Organization{org1, org2}

	// Convert to DTOs
	dtoList := apputils.ToListOfOrganizationDTO(orgList)

	// Assert length of the list
	assert.Len(t, dtoList, 2, "The length of DTO list should match the length of organization list")

	// Assert contents
	for i, dto := range dtoList {
		org := orgList[i]
		assert.Equal(t, org.GetID(), dto.ID, "ID should match")
		assert.Equal(t, org.GetName(), dto.Name, "Name should match")
		assert.True(t, reflect.DeepEqual(org.GetNames(), dto.AllNames), "AllNames should match")
	}
}

func TestToOrganizationListDTO(t *testing.T) {
	// Mock organization models similar to the previous test
	orgID1 := "org789"
	name1 := "Test Org 3"
	names1 := []string{name1, "Test Org Department 3"}
	org1 := new(appmocks.MockOrganization)
	org1.On("GetID").Return(orgID1)
	org1.On("GetName").Return(name1)
	org1.On("GetNames").Return(names1)

	orgList := []appmodels.Organization{org1}

	// Convert to OrganizationListDTO
	listDTO := apputils.ToOrganizationListDTO(orgList)

	// Assert payload size
	assert.Len(t, listDTO.Payload, 1, "The size of Organizations in OrganizationListDTO should be 1")

	// Assert contents
	dto := listDTO.Payload[0]
	assert.Equal(t, orgID1, dto.ID, "ID should match")
	assert.Equal(t, name1, dto.Name, "Name should match")
	assert.True(t, reflect.DeepEqual(names1, dto.AllNames), "AllNames should match")
}
