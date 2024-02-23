// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package api

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/api/indexapi"
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"net/http"
)

func GetOpenApiInfo() *openapi3.Info {
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
