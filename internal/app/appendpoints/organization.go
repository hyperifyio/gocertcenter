// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// GetOrganizationDefinitions returns OpenAPI definitions
func (c *ApiController) GetOrganizationDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a collection of organization entities",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.OrganizationDTO{}},
				},
			},
		},
	}
}

// GetOrganization handles a request
func (c *ApiController) GetOrganization(response apitypes.IResponse, request apitypes.IRequest) error {
	organization := request.GetVariable("organization")
	model, err := c.appController.GetOrganizationModel(organization)
	if err != nil {
		return fmt.Errorf("ApiController.GetOrganization: failed to find model: %v", err)
	}
	log.Printf("[GetOrganization] Request: model = %v", model)
	data := apputils.GetOrganizationDTO(model)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetOrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetOrganization
