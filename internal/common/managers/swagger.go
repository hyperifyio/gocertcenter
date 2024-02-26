// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"context"
	"fmt"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/davidebianchi/gswagger/support/gorilla"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

// SwaggerManager implements apitypes.ISwaggerManager
type SwaggerManager struct {
	swaggerRouter *swagger.Router[gorilla.HandlerFunc, gorilla.Route]
}

func (r *SwaggerManager) GenerateAndExposeOpenapi() error {
	return r.swaggerRouter.GenerateAndExposeOpenapi()
}

func (r *SwaggerManager) AddRoute(method string, path string, handler http.HandlerFunc, definitions swagger.Definitions) (*mux.Route, error) {
	route, err := r.swaggerRouter.AddRoute(method, path, ToGorillaHandlerFunc(handler), definitions)
	if err != nil {
		return nil, fmt.Errorf("[SwaggerManager] Add Route %s %s failed: %w", method, path, err)
	}
	return FromGorillaRoute(route), nil
}

func NewSwaggerManager(
	router *mux.Router,
	context *context.Context,
	url string,
	description string,
	info *openapi3.Info,
) (ISwaggerManager, error) {

	swaggerRouter, err := swagger.NewRouter(
		gorilla.NewRouter(router),
		swagger.Options{
			Context: *context,
			Openapi: &openapi3.T{
				Info: info,
				Servers: []*openapi3.Server{
					{
						URL:         url,
						Description: description,
					},
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create swagger router: %v", err)
	}

	var manager ISwaggerManager = &SwaggerManager{
		swaggerRouter: swaggerRouter,
	}
	return manager, nil
}

var _ ISwaggerManager = (*SwaggerManager)(nil)
