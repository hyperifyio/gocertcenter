// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apitypes

import "net/url"

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
	Start() error
	SetupRoutes()
	GetAddress() string
}

type IRequest interface {
	IsMethodGet() bool
	GetMethod() string
	GetURL() *url.URL
	GetVars() map[string]string
}

// RequestHandlerFunc defines the type for handlers in this API.
type RequestHandlerFunc func(IResponse, IRequest, IServer)
