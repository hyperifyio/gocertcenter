// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// CreateOrganizationDefinitions returns OpenAPI definitions
func (c *ApiController) CreateOrganizationDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a collection of organization entities",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.IndexDTO{}},
				},
			},
		},
	}
}

// CreateOrganization handles a request
func (c *ApiController) CreateOrganization(response apitypes.IResponse, request apitypes.IRequest) {

	body, err := c.DecodeOrganizationFromRequestBody(request)
	if err != nil {
		log.Printf("Request body invalid: %v", err)
		response.SendError(400, "request body invalid")
		return
	}

	id := body.ID
	names := body.AllNames
	var certificates []appmodels.ICertificate
	var keys []appmodels.IPrivateKey

	model := appmodels.NewOrganization(id, names)

	savedModel, err := c.appController.NewOrganization(model)
	if err != nil {
		response.SendError(500, "Failed to create an organization")
		return
	}

	keyDTOs, err := apputils.ToPrivateKeyDTOList(c.certManager, keys)
	if err != nil {
		response.SendError(500, "Failed to create an organization")
		return
	}

	data := appdtos.NewOrganizationCreatedDTO(
		apputils.GetOrganizationDTO(savedModel),
		apputils.ToCertificateDTOList(certificates),
		keyDTOs,
	)
	response.Send(http.StatusOK, data)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateOrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateOrganization
