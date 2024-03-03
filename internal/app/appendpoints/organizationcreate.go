// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateOrganizationDefinitions returns OpenAPI definitions
func (c *HttpApiController) CreateOrganizationDefinitions() swagger.Definitions {
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
func (c *HttpApiController) CreateOrganization(response apitypes.Response, request apitypes.Request) error {

	body, err := c.DecodeOrganizationFromRequestBody(request)
	if err != nil {
		return c.badRequest(response, request, "body invalid", err)
	}

	slug := body.Slug
	name := body.Name
	names := body.AllNames
	var certificates []appmodels.Certificate
	var keys []appmodels.PrivateKey

	if slug == "" && name != "" {
		slug = name
	}

	if name == "" && slug != "" {
		name = slug
	}

	if len(names) == 0 {
		names = append(names, name)
	}

	randomManager := c.certManager.RandomManager()

	newOrgId, err := apputils.GenerateSerialNumber(randomManager)
	if err != nil {
		return c.internalServerError(response, request, fmt.Errorf("CreateOrganization: failed: %w", err))
	}

	slug = apputils.Slugify(slug)

	model := appmodels.NewOrganization(newOrgId, slug, names)

	savedModel, err := c.appController.NewOrganization(model)
	if err != nil {
		return c.conflict(response, request, err, "organization")
	}

	keyDTOs, err := apputils.ToPrivateKeyDTOList(c.certManager, keys)
	if err != nil {
		return c.internalServerError(response, request, err)
	}

	dto := appdtos.NewOrganizationCreatedDTO(
		apputils.ToOrganizationDTO(savedModel),
		apputils.ToListOfCertificateDTO(certificates),
		keyDTOs,
	)
	return c.ok(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*HttpApiController)(nil).CreateOrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*HttpApiController)(nil).CreateOrganization
