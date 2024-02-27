// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	swagger "github.com/davidebianchi/gswagger"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
)

// GetOrganizationCertificateDefinitions returns OpenAPI definitions
func (c *ApiController) GetOrganizationCertificateDefinitions() swagger.Definitions {
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

// GetOrganizationCertificate handles a request
func (c *ApiController) GetOrganizationCertificate(response apitypes.IResponse, request apitypes.IRequest) error {
	organization := request.GetVariable("organization")
	log.Printf("[GetOrganizationCertificate] organization = %s", organization)

	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[ApiController.GetOrganizationCertificate]: failed to parse serialNumber: %v", err)
	}
	log.Printf("[GetOrganizationCertificate] serialNumber = %s", serialNumber.String())

	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[ApiController.GetOrganizationCertificate]: failed to find organizationController: %v", err)
	}

	model, err := organizationController.GetCertificateModel(serialNumber)
	if err != nil {
		return fmt.Errorf("[ApiController.GetOrganizationCertificate]: failed to find model: %v", err)
	}
	log.Printf("[GetOrganizationCertificate] Request: model = %v", model)
	data := apputils.GetCertificateDTO(model)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetOrganizationCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetOrganizationCertificate
