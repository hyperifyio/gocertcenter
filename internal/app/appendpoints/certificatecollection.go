// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
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

	// Get organization ID
	organization := request.GetVariable("organization")

	// Parse serial number
	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[GetCertificateCollection]: failed to parse serialNumber: %v", err)
	}
	log.Printf("[GetCertificateCollection] serialNumber = %s", serialNumber.String())

	// Get Organization controller
	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[GetCertificateCollection]: could not get a controller: %w", err)
	}

	// Fetch certificate controller
	certificateController, err := organizationController.GetCertificateController(serialNumber)
	if err != nil {
		return fmt.Errorf("[GetCertificateCollection]: failed to find certificate controller: %v", err)
	}

	// Get certificate list
	list, err := certificateController.GetChildCertificateCollection(certificateType)
	if err != nil {
		return fmt.Errorf("[GetCertificateCollection]: could not get a collection: %w", err)
	}

	log.Printf("[GetCertificateCollection]: Request: list = %d", len(list))

	data := apputils.ToCertificateListDTO(list)
	response.Send(http.StatusOK, data)

	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetCertificateCollectionDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetCertificateCollection
