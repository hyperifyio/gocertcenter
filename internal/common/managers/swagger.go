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

// GorillaSwaggerManager implements SwaggerManager
type GorillaSwaggerManager struct {
	swaggerRouter *swagger.Router[gorilla.HandlerFunc, gorilla.Route]
}

func (r *GorillaSwaggerManager) GenerateAndExposeOpenapi() error {
	return r.swaggerRouter.GenerateAndExposeOpenapi()
}

func (r *GorillaSwaggerManager) AddRoute(method string, path string, handler http.HandlerFunc, definitions swagger.Definitions) (*mux.Route, error) {
	route, err := r.swaggerRouter.AddRoute(method, path, ToGorillaHandlerFunc(handler), definitions)
	if err != nil {
		return nil, fmt.Errorf("[GorillaSwaggerManager] Add Route %s %s failed: %w", method, path, err)
	}
	return FromGorillaRoute(route), nil
}

func NewSwaggerManager(
	router *mux.Router,
	context *context.Context,
	url string,
	description string,
	info *openapi3.Info,
) (SwaggerManager, error) {

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

	var manager SwaggerManager = &GorillaSwaggerManager{
		swaggerRouter: swaggerRouter,
	}
	return manager, nil
}

var _ SwaggerManager = (*GorillaSwaggerManager)(nil)
