// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// OrganizationDefinitions returns OpenAPI definitions
func (c *HttpApiController) OrganizationDefinitions() swagger.Definitions {
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

// Organization handles a request
func (c *HttpApiController) Organization(response apitypes.Response, request apitypes.Request) error {
	controller, err := c.organizationController(request)
	if err != nil {
		return c.notFound(response, request, err)
	}
	model := controller.Organization()
	c.logf(request, "model = %v", model)
	dto := apputils.ToOrganizationDTO(model)
	return c.ok(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).OrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).Organization
