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

	// Get organization ID
	organization := request.GetVariable("organization")

	// Parse root serial number
	rootSerialNumberString := request.GetVariable("rootSerialNumber")
	rootSerialNumber, err := appmodels.ParseSerialNumber(rootSerialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[GetCertificate]: failed to parse rootSerialNumber: %v", err)
	}
	log.Printf("[GetCertificate] rootSerialNumber = %s", rootSerialNumber.String())

	// Parse serial number
	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return fmt.Errorf("[GetCertificate]: failed to parse serialNumber: %v", err)
	}
	log.Printf("[GetCertificate] serialNumber = %s", serialNumber.String())

	// Get Organization controller
	organizationController, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return fmt.Errorf("[GetCertificate]: could not get a controller: %w", err)
	}

	// Fetch root certificate controller
	rootCertificateController, err := organizationController.GetCertificateController(rootSerialNumber)
	if err != nil {
		return fmt.Errorf("[GetCertificate]: failed to find root certificate controller: %v", err)
	}

	// Fetch root certificate controller
	certificateController, err := rootCertificateController.GetChildCertificateController(serialNumber)
	if err != nil {
		return fmt.Errorf("[GetCertificate]: failed to find child certificate controller: %v", err)
	}

	model := certificateController.GetCertificateModel()
	log.Printf("[GetCertificate] Request: model = %v", model)
	data := apputils.GetCertificateDTO(model)
	response.Send(http.StatusOK, data)
	return nil
}

var _ apitypes.RequestDefinitionsFunc = (*ApiController)(nil).GetCertificateDefinitions
var _ apitypes.RequestHandlerFunc = (*ApiController)(nil).GetCertificate
