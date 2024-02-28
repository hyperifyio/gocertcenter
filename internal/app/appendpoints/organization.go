// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// GetOrganizationDefinitions returns OpenAPI definitions
func (c *ApiController) GetOrganizationDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns an organization entity",
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
	controller, err := c.getOrganizationController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}
	model := controller.GetOrganizationModel()
	c.logf(request, "model = %v", model)
	dto := apputils.GetOrganizationDTO(model)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetOrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetOrganization
