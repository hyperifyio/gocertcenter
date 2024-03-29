// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func ToOrganizationDTO(o appmodels.Organization) appdtos.OrganizationDTO {
	return appdtos.NewOrganizationDTO(
		o.ID().String(),
		o.Slug(),
		o.Name(),
		o.Names(),
	)
}

func ToListOfOrganizationDTO(list []appmodels.Organization) []appdtos.OrganizationDTO {
	result := make([]appdtos.OrganizationDTO, len(list))
	for i, v := range list {
		result[i] = ToOrganizationDTO(v)
	}
	return result
}

func ToOrganizationListDTO(list []appmodels.Organization) appdtos.OrganizationListDTO {
	payload := ToListOfOrganizationDTO(list)
	return appdtos.NewOrganizationListDTO(payload)
}
