// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
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
		log.Printf("Request body invalid: %v", err)
		response.SendError(400, "request body invalid")
		return nil
	}

	id := body.ID
	names := body.AllNames
	var certificates []appmodels.ICertificate
	var keys []appmodels.IPrivateKey

	model := appmodels.NewOrganization(id, names)

	savedModel, err := c.appController.NewOrganization(model)
	if err != nil {
		return fmt.Errorf("ApiController.CreateOrganization: saving failed: %v", err)
	}

	keyDTOs, err := apputils.ToPrivateKeyDTOList(c.certManager, keys)
	if err != nil {
		return fmt.Errorf("ApiController.CreateOrganization: ToPrivateKeyDTOList failed: %v", err)
	}

	data := appdtos.NewOrganizationCreatedDTO(
		apputils.GetOrganizationDTO(savedModel),
		apputils.ToListOfCertificateDTO(certificates),
		keyDTOs,
	)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).CreateOrganizationDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).CreateOrganization
