// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetOrganizationCollectionDefinitions returns OpenAPI definitions
func (c *ApiController) GetOrganizationCollectionDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a collection of organization entities",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.OrganizationListDTO{}},
				},
			},
		},
	}
}

// GetOrganizationCollection handles a request
func (c *ApiController) GetOrganizationCollection(response apitypes.IResponse, request apitypes.IRequest) {
	list, err := c.appController.GetOrganizationCollection()
	if err != nil {
		response.SendError(500, "Could not get a collection")
	}
	log.Printf("[GetOrganizationCollection] Request: list = %d", len(list))
	data := apputils.ToOrganizationListDTO(list)
	response.Send(http.StatusOK, data)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetOrganizationCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetOrganizationCollection
