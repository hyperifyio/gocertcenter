// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors

import (
	"log"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// NotFound handles not found errors
func NotFound(response apitypes.Response, request apitypes.Request) error {
	log.Printf("[NotFound] %s %s", request.Method(), request.URL())
	response.SendNotFoundError()
	return nil
}

var _ apitypes.RequestHandlerFunc = NotFound
