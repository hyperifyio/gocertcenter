// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type OrganizationDTO struct {
	ID       string   `json:"id"`
	Slug     string   `json:"slug"`
	Name     string   `json:"name"`
	AllNames []string `json:"allNames"`
}

func NewOrganizationDTO(
	id, slug, name string,
	allNames []string,
) OrganizationDTO {
	return OrganizationDTO{
		ID:       id,
		Slug:     slug,
		Name:     name,
		AllNames: allNames,
	}
}
