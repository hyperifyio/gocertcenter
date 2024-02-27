// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter"
	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
)

func (c *ApiController) GetIndexDefinitions() swagger.Definitions {
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

// GetIndex handles the GET requests at the root URL.
func (c *ApiController) GetIndex(response apitypes.IResponse, request apitypes.IRequest) error {
	log.Printf("[IndexController] Request")
	data := appdtos.NewIndexDTO(gocertcenter.Name, gocertcenter.Version)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetIndexDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetIndex
