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
