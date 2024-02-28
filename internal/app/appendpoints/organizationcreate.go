// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateOrganizationDefinitions returns OpenAPI definitions
func (c *ApiController) CreateOrganizationDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Creates an organization",
		Description: "",
		RequestBody: &swagger.ContentValue{
			Description: "Organization data",
			Content: swagger.Content{
				"application/json": {
					Value: appdtos.OrganizationDTO{},
				},
			},
		},
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.OrganizationDTO{}},
				},
			},
		},
	}
}

// CreateOrganization handles a request
func (c *ApiController) CreateOrganization(response apitypes.IResponse, request apitypes.IRequest) error {

	body, err := c.DecodeOrganizationFromRequestBody(request)
	if err != nil {
		return c.sendBadRequest(response, request, "body invalid", err)
	}

	id := body.ID
	name := body.Name
	names := body.AllNames
	var certificates []appmodels.ICertificate
	var keys []appmodels.IPrivateKey

	if id == "" && name != "" {
		id = name
	}

	if name == "" && id != "" {
		name = id
	}

	if len(names) == 0 {
		names = append(names, name)
	}

	id = apputils.Slugify(id)

	model := appmodels.NewOrganization(id, names)

	savedModel, err := c.appController.NewOrganization(model)
	if err != nil {
		return c.sendConflict(response, request, err, "organization")
	}

	keyDTOs, err := apputils.ToPrivateKeyDTOList(c.certManager, keys)
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}

	dto := appdtos.NewOrganizationCreatedDTO(
		apputils.GetOrganizationDTO(savedModel),
		apputils.ToListOfCertificateDTO(certificates),
		keyDTOs,
	)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateOrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateOrganization
