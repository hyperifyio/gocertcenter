// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func GetOrganizationDTO(o appmodels.IOrganization) appdtos.OrganizationDTO {
	return appdtos.NewOrganizationDTO(
		o.GetID(),
		o.GetName(),
		o.GetNames(),
	)
}

func ToListOfOrganizationDTO(list []appmodels.IOrganization) []appdtos.OrganizationDTO {
	result := make([]appdtos.OrganizationDTO, len(list))
	for i, v := range list {
		result[i] = GetOrganizationDTO(v)
	}
	return result
}

func ToOrganizationListDTO(list []appmodels.IOrganization) appdtos.OrganizationListDTO {
	payload := ToListOfOrganizationDTO(list)
	return appdtos.NewOrganizationListDTO(payload)
}
