// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appdtos

type OrganizationListDTO struct {
	Payload []OrganizationDTO `json:"payload" jsonschema:"title=Organization Payload DTOs,required"`
}

func NewOrganizationListDTO(
	payload []OrganizationDTO,
) OrganizationListDTO {
	return OrganizationListDTO{
		Payload: payload,
	}
}
