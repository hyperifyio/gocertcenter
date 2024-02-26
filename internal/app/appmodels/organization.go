// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

// Organization model implements IOrganization
type Organization struct {
	id    string
	names []string
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
	originalSlice := o.names
	sliceCopy := make([]string, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
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

// Compile time assertion for implementing the interface
var _ IOrganization = (*Organization)(nil)
