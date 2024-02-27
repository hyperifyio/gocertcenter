// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// GetRootCertificateCollectionDefinitions returns OpenAPI definitions
func (c *ApiController) GetRootCertificateCollectionDefinitions() swagger.Definitions {
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

// GetRootCertificateCollection handles a request to get organization's certificates
func (c *ApiController) GetRootCertificateCollection(response apitypes.IResponse, request apitypes.IRequest) error {
	organization := request.GetVariable("organization")

	controller, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[GetRootCertificateCollection]: could not get a controller: %w", err)
	}

	list, err := controller.GetCertificateCollection()
	if err != nil {
		return fmt.Errorf("[GetRootCertificateCollection]: could not get a collection: %w", err)
	}

	log.Printf("[GetRootCertificateCollection]: Request: list = %d", len(list))
	data := apputils.ToCertificateListDTO(list)
	response.Send(http.StatusOK, data)

	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetRootCertificateCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetRootCertificateCollection
