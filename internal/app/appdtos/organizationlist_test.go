// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func TestNewOrganizationListDTO(t *testing.T) {
	// Prepare a sample payload of OrganizationDTO instances
	payload := []appdtos.OrganizationDTO{
		{
			ID:       "1",
			Name:     "Organization One",
			AllNames: []string{"Org One", "Organization 1"},
		},
		{
			ID:       "2",
			Name:     "Organization Two",
			AllNames: []string{"Org Two", "Organization 2"},
		},
	}

	// Create an OrganizationListDTO instance using the constructor
	organizationListDTO := appdtos.NewOrganizationListDTO(payload)

	// Assert that the payload is correctly assigned
	assert.Equal(t, payload, organizationListDTO.Payload, "Payload should match the input payload")
}
