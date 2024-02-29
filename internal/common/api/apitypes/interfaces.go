// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apitypes

import (
	"context"
	"hash"
	"io"
	"net/url"

	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// NewSwaggerManagerFunc is a factory function for SwaggerManager instances
type NewSwaggerManagerFunc func(
	router *mux.Router,
	context *context.Context,
	url string,
	description string,
	info *openapi3.Info,
) (managers.SwaggerManager, error)

// Route represents a single API route
type Route struct {
	Method      string
	Path        string
	Handler     RequestHandlerFunc
	Definitions swagger.Definitions
}

type Response interface {
	Send(statusCode int, data interface{})
	SendError(statusCode int, error string)
	SendMethodNotSupportedError()
	SendBytes([]byte) error
	SendNotFoundError()
	SendConflictError(error string)
	SendInternalServerError(error string)
	SetHeader(name, value string)
}

// Server defines the methods available from the Server
// that are needed by the HTTP handlers.
type Server interface {

	// IsStarted returns true if this service has been started
	IsStarted() bool

	// GetAddress returns the address where this service will listen on
	GetAddress() string

	// GetURL returns the full URL for the server
	GetURL() string

	// InitSetup can be called to initialize default values before calling
	// other Setup methods. This is not normally required, since it is called in
	// the NewServer()
	InitSetup() error

	// SetupRoutes can be called to configure route handlers from an array
	SetupRoutes(routes []Route) error

	// SetupHandler can be called to configure single route handler
	SetupHandler(
		method string,
		path string,
		apiHandler RequestHandlerFunc,
		definitions swagger.Definitions,
	) error

	// SetupNotFoundHandler can be used to configure the Not Found error handler
	SetupNotFoundHandler(handler RequestHandlerFunc)

	// SetupMethodNotAllowedHandler can be used to configure the Method Not Allowed handler
	SetupMethodNotAllowedHandler(handler RequestHandlerFunc)

	// FinalizeSetup can be called after calling other Setup routes to finalize
	// the route configuration. This is not normally needed since it is called
	// automatically in Start()
	FinalizeSetup() error

	// Start can be used to start the service
	Start() error

	// Stop can be used to stop the service
	Stop() error
}

type Request interface {
	IsMethodGet() bool
	GetMethod() string
	GetURL() *url.URL
	GetVariable(name string) string
	GetQueryParam(name string) string
	Body() io.ReadCloser
	GetBodyBytes() ([]byte, error)
	GetHeader(name string) string
}

// RequestHandlerFunc defines the type for handlers in this API.
type RequestHandlerFunc func(Response, Request) error

// RequestDefinitionsFunc defines the type for OpenAPI definitions function
type RequestDefinitionsFunc func() swagger.Definitions

// ApplicationInfoFunc defines the type for OpenAPI info structure
type ApplicationInfoFunc func() *openapi3.Info

// ApplicationRoutesFunc defines a function which returns application routes
type ApplicationRoutesFunc func() []Route

// NewServerManagerFunc is a factory function for ServerManager instances
type NewServerManagerFunc func(
	address string,
	handler *mux.Router,
) managers.ServerManager

type Hash64FactoryFunc func() hash.Hash64

// AppController defines common methods for each application end-point
// controller
type AppController interface {
	GetInfo() *openapi3.Info

	GetRoutes() []Route
}
