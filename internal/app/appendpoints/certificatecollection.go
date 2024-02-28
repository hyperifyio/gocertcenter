// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetCertificateCollectionDefinitions returns OpenAPI definitions
func (c *ApiController) GetCertificateCollectionDefinitions() swagger.Definitions {
	return swagger.Definitions{
		Summary:     "Returns a collection of root certificate entities",
		Description: "",
		Responses: map[int]swagger.ContentValue{
			200: {
				Content: swagger.Content{
					"application/json": {Value: appdtos.CertificateListDTO{}},
				},
			},
		},
	}
}

// GetCertificateCollection handles a request to get organization's certificates
func (c *ApiController) GetCertificateCollection(response apitypes.IResponse, request apitypes.IRequest) error {

	// certificateType is server, client, root or intermediate
	certificateType := request.GetQueryParam("type")
	if !(certificateType == "" || certificateType == "server" || certificateType == "client" || certificateType == "root") {
		return c.sendBadRequest(response, request, "query param invalid: type", nil)
	}

	// Fetch root certificate controller
	controller, err := c.getRootCertificateController(request)
	if err != nil {
		return c.sendNotFound(response, request, err)
	}

	// Get certificate list
	list, err := controller.GetChildCertificateCollection(certificateType)
	if err != nil {
		return c.sendInternalServerError(response, request, err)
	}

	c.logf(request, "list len = %d", len(list))
	dto := apputils.ToCertificateListDTO(list)
	return c.sendOK(response, dto)
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetCertificateCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetCertificateCollection
