// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// GetCertificateDefinitions returns OpenAPI definitions
func (c *ApiController) GetCertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a certificate entity owned by a root certificate",
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

// GetCertificate handles a request
func (c *ApiController) GetCertificate(response apitypes.IResponse, request apitypes.IRequest) error {

	// Fetch the certificate controller
	controller, err := c.getInnerCertificateController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}

	model := controller.GetCertificateModel()
	dto := apputils.GetCertificateDTO(model)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetCertificate
