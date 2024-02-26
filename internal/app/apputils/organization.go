// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

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
