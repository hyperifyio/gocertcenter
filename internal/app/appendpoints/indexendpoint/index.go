// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package indexendpoint

import (
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func IndexDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns information about the running server",
		Description: "This includes the software name and a version",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.IndexDTO{}},
				},
			},
		},
	}
}

// Index handles the GET requests at the root URL.
func Index(response apitypes.IResponse, request apitypes.IRequest, server apitypes.IServer) {

	log.Printf("[Index] Request")

	data := appdtos.NewIndexDTO(gocertcenter.Name, gocertcenter.Version)

	response.Send(http.StatusOK, data)

}
