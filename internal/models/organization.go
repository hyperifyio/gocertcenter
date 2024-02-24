// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"github.com/hyperifyio/gocertcenter/internal/dtos"
)

// Organization model implements IOrganization
type Organization struct {
	id    string
	names []string
}

// Compile time assertion for implementing the interface
var _ IOrganization = (*Organization)(nil)

func (o *Organization) GetDTO() dtos.OrganizationDTO {
	return dtos.NewOrganizationDTO(
		o.GetID(),
		o.GetName(),
		o.GetNames(),
	)
}

// GetID returns unique identifier for this organization
func (o *Organization) GetID() string {
	return o.id
}

// GetName returns the primary organization name
func (o *Organization) GetName() string {
	slice := o.GetNames()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

// GetNames returns the full name of the organization including department
func (o *Organization) GetNames() []string {
	return o.names
}

// NewOrganization creates a organization model from existing data
func NewOrganization(
	id string,
	names []string,
) *Organization {
	return &Organization{
		id:    id,
		names: names,
	}
}
