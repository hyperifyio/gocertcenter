// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// GetRootCertificateDefinitions returns OpenAPI definitions
func (c *ApiController) GetRootCertificateDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns an root certificate",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.CertificateDTO{}},
				},
			},
		},
	}
}

// GetRootCertificate handles a request
func (c *ApiController) GetRootCertificate(response apitypes.IResponse, request apitypes.IRequest) error {

	controller, err := c.getRootCertificateController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}

	model := controller.GetCertificateModel()
	c.logf(request, "model = %v", model)

	dto := apputils.ToCertificateDTO(model)
	return c.sendOK(response, dto)

}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetRootCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetRootCertificate
