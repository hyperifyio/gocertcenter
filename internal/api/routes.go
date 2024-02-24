// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package api

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/api/indexapi"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
)

func GetInfo() *openapi3.Info {
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

func GetRoutes() []apitypes.Route {
	return []apitypes.Route{
		{
			Method:      http.MethodGet,
			Path:        "/",
			Handler:     indexapi.Index,
			Definitions: indexapi.IndexDefinitions(),
		},
	}
}