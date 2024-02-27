// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors

import (
	"log"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// MethodNotAllowed handles not found errors
func MethodNotAllowed(response apitypes.IResponse, request apitypes.IRequest) error {
	log.Printf("[MethodNotAllowed] %s %s", request.GetMethod(), request.GetURL())
	response.SendMethodNotSupportedError()
	return nil
}

var _ apitypes.RequestHandlerFunc = MethodNotAllowed
