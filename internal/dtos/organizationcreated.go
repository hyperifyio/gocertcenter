// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos

type OrganizationCreatedDTO struct {
	Organization OrganizationDTO `json:"organization"`

	// Certificates may be used to provide initial client certificates for
	// configuring a new organization
	Certificates []CertificateDTO `json:"certificates"`

	// PrivateKeys may be used to provide initial client certificates for
	// configuring a new organization
	PrivateKeys []PrivateKeyDTO `json:"privateKeys"`
}

func NewOrganizationCreatedDTO(
	organization OrganizationDTO,
	certificates []CertificateDTO,
	privateKeys []PrivateKeyDTO,
) OrganizationCreatedDTO {
	return OrganizationCreatedDTO{
		Organization: organization,
		Certificates: certificates,
		PrivateKeys:  privateKeys,
	}
}
