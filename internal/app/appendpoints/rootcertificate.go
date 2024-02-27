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
	organization := request.GetVariable("organization")
	log.Printf("[GetRootCertificate] organization = %s", organization)

	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[ApiController.GetRootCertificate]: failed to parse serialNumber: %v", err)
	}
	log.Printf("[GetRootCertificate] serialNumber = %s", serialNumber.String())

	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[ApiController.GetRootCertificate]: failed to find organizationController: %v", err)
	}

	model, err := organizationController.GetCertificateModel(serialNumber)
	if err != nil {
		return fmt.Errorf("[ApiController.GetRootCertificate]: failed to find model: %v", err)
	}
	log.Printf("[GetRootCertificate] Request: model = %v", model)
	data := apputils.GetCertificateDTO(model)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetRootCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetRootCertificate
