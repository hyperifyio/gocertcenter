// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apierrors

import (
	"log"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// NotFound handles not found errors
func NotFound(response apitypes.IResponse, request apitypes.IRequest) {
	log.Printf("[NotFound] %s %s", request.GetMethod(), request.GetURL())
	response.SendNotFoundError()
}
