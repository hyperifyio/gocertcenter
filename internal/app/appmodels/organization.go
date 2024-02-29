// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

// OrganizationModel model implements Organization
type OrganizationModel struct {
	id    string
	names []string
}

// GetID returns unique identifier for this organization
func (o *OrganizationModel) ID() string {
	return o.id
}

// GetName returns the primary organization name
func (o *OrganizationModel) Name() string {
	slice := o.Names()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

// GetNames returns the full name of the organization including department
func (o *OrganizationModel) Names() []string {
	originalSlice := o.names
	sliceCopy := make([]string, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

// NewOrganization creates a organization model from existing data
func NewOrganization(
	id string,
	names []string,
) *OrganizationModel {
	return &OrganizationModel{
		id:    id,
		names: names,
	}
}

// Compile time assertion for implementing the interface
var _ Organization = (*OrganizationModel)(nil)
