// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors

import (
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"log"
)

// NotFound handles not found errors
func NotFound(response apitypes.IResponse, request apitypes.IRequest, server apitypes.IServer) {
	log.Printf("[NotFound] %s %s", request.GetMethod(), request.GetURL())
	response.SendNotFoundError()
}
