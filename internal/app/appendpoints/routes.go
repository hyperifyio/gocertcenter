// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"net/http"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *ApiController) GetRoutes() []apitypes.Route {
	return []apitypes.Route{
		{
			Method:      http.MethodGet,
			Path:        "/",
			Handler:     c.GetIndex,
			Definitions: c.GetIndexDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations",
			Handler:     c.GetOrganizationCollection,
			Definitions: c.GetOrganizationCollectionDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations",
			Handler:     c.CreateOrganization,
			Definitions: c.CreateOrganizationDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}",
			Handler:     c.GetOrganization,
			Definitions: c.GetOrganizationDefinitions(),
		},
	}
}

var _ apitypes.ApplicationRoutesFunc = (*ApiController)(nil).GetRoutes
