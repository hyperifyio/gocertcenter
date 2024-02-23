// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package indexapi

import (
	"github.com/hyperifyio/gocertcenter/internal/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"log"
	"net/http"
)

// Index handles the GET requests at the root URL.
func Index(response apitypes.IResponse, request apitypes.IRequest, server apitypes.IServer) {

	if !request.IsMethodGet() {
		response.SendMethodNotSupportedError()
		return
	}

	log.Printf("[Index] Request")

	data := dtos.NewIndexDTO("0.0.1")

	response.Send(http.StatusOK, data)

}
