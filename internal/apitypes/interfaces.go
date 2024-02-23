// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apitypes

import (
	swagger "github.com/davidebianchi/gswagger"
	"net/url"
)

// Route represents a single API route
type Route struct {
	Method      string
	Path        string
	Handler     RequestHandlerFunc
	Definitions swagger.Definitions
}

type IResponse interface {
	Send(statusCode int, data interface{})
	SendError(statusCode int, error string)
	SendMethodNotSupportedError()
	SendNotFoundError()
	SendConflictError(error string)
	SendInternalServerError(error string)
}

// IServer defines the methods available from the Server
// that are needed by the HTTP handlers.
type IServer interface {

	// IsStarted returns true if this service has been started
	IsStarted() bool

	// GetAddress returns the address where this service will listen on
	GetAddress() string

	// GetURL returns the full URL for the server
	GetURL() string

	// InitSetup can be called to initialize default values before calling
	// other Setup methods. This is not normally required, since it is called in
	// the NewServer()
	InitSetup()

	// SetupRoutes can be called to configure route handlers from an array
	SetupRoutes(routes []Route)

	// SetupHandler can be called to configure single route handler
	SetupHandler(
		method string,
		path string,
		apiHandler RequestHandlerFunc,
		definitions swagger.Definitions,
	)

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

type IRequest interface {
	IsMethodGet() bool
	GetMethod() string
	GetURL() *url.URL
	GetVars() map[string]string
}

// RequestHandlerFunc defines the type for handlers in this API.
type RequestHandlerFunc func(IResponse, IRequest, IServer)
