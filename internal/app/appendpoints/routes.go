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
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/certificates/{serialNumber}",
			Handler:     c.GetCertificate,
			Definitions: c.GetCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates/{serialNumber}/certificates",
			Handler:     c.GetCertificateCollection,
			Definitions: c.GetCertificateCollectionDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations/{organization}/certificates/{serialNumber}/certificates",
			Handler:     c.CreateCertificate,
			Definitions: c.CreateCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates/{serialNumber}",
			Handler:     c.GetRootCertificate,
			Definitions: c.GetRootCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates",
			Handler:     c.GetRootCertificateCollection,
			Definitions: c.GetRootCertificateCollectionDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations/{organization}/certificates",
			Handler:     c.CreateRootCertificate,
			Definitions: c.CreateRootCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}",
			Handler:     c.GetOrganization,
			Definitions: c.GetOrganizationDefinitions(),
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
			Path:        "/",
			Handler:     c.GetIndex,
			Definitions: c.GetIndexDefinitions(),
		},
	}
}

var _ apitypes.ApplicationRoutesFunc = (*ApiController)(nil).GetRoutes
