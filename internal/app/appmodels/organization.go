// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"math/big"
)

// OrganizationModel model implements Organization
type OrganizationModel struct {
	id    *big.Int
	slug  string
	names []string
}

// ID returns the numeric unique identifier for this organization
func (o *OrganizationModel) ID() *big.Int {
	return o.id
}

// Slug returns unique, URL-friendly identifier for the organization
func (o *OrganizationModel) Slug() string {
	return o.slug
}

// Name returns the primary organization name
func (o *OrganizationModel) Name() string {
	slice := o.Names()
	if len(slice) > 0 {
		return slice[0]
	}
	return ""
}

// Names returns the full name of the organization including department
func (o *OrganizationModel) Names() []string {
	originalSlice := o.names
	sliceCopy := make([]string, len(originalSlice))
	copy(sliceCopy, originalSlice)
	return sliceCopy
}

// NewOrganization creates a organization model from existing data
func NewOrganization(
	id *big.Int,
	slug string,
	names []string,
) *OrganizationModel {
	return &OrganizationModel{
		id:    id,
		slug:  slug,
		names: names,
	}
}

// Compile time assertion for implementing the interface
var _ Organization = (*OrganizationModel)(nil)
