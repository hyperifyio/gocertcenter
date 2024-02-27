// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *ApiController) GetInfo() *openapi3.Info {
	return &openapi3.Info{
		Title:   gocertcenter.Name,
		Version: gocertcenter.Version,
		License: &openapi3.License{
			Name: gocertcenter.LicenseName,
			URL:  gocertcenter.LicenseURL,
		},
		Description: gocertcenter.Description,
		Contact: &openapi3.Contact{
			Name:  gocertcenter.SupportName,
			URL:   gocertcenter.SupportURL,
			Email: gocertcenter.SupportEmail,
		},
	}
}

var _ apitypes.ApplicationInfoFunc = (*ApiController)(nil).GetInfo
