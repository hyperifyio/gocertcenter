// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors

import (
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"log"
)

// MethodNotAllowed handles not found errors
func MethodNotAllowed(response apitypes.IResponse, request apitypes.IRequest, server apitypes.IServer) {
	log.Printf("[MethodNotAllowed] %s %s", request.GetMethod(), request.GetURL())
	response.SendMethodNotSupportedError()
}
