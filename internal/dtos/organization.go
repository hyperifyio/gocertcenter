// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package dtos

type OrganizationDTO struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	AllNames []string `json:"allNames"`
}

func NewOrganizationDTO(
	id string,
	name string,
	allNames []string,
) OrganizationDTO {
	return OrganizationDTO{
		ID:       id,
		Name:     name,
		AllNames: allNames,
	}
}
