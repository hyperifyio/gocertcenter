// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"strings"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *ApiController) getRequestOrganization(request apitypes.IRequest) string {
	organization := request.GetVariable("organization")
	organization = strings.Trim(strings.ToLower(organization), " ")
	if organization != "" {
		if err := apputils.ValidateOrganizationID(organization); err != nil {
			c.logf(request, "invalid organization: '%s': %v", organization, err)
			return ""
		}
		c.logf(request, "organization = '%s'", organization)
	} else {
		c.logf(request, "no organization found")
	}
	return organization
}

func (c *ApiController) getRootSerialNumber(request apitypes.IRequest) (appmodels.ISerialNumber, error) {
	serialNumberString := request.GetVariable("rootSerialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to parse rootSerialNumber: %v", request.GetMethod(), request.GetURL(), err)
	}
	c.logf(request, "rootSerialNumber = %s", serialNumber.String())
	return serialNumber, nil
}

func (c *ApiController) getSerialNumber(request apitypes.IRequest) (appmodels.ISerialNumber, error) {
	serialNumberString := request.GetVariable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to parse serialNumber: %v", request.GetMethod(), request.GetURL(), err)
	}
	c.logf(request, "serialNumber = %s", serialNumber.String())
	return serialNumber, nil
}

func (c *ApiController) getOrganizationController(request apitypes.IRequest) (appmodels.IOrganizationController, error) {
	organization := c.getRequestOrganization(request)
	if organization == "" {
		return nil, fmt.Errorf("[%s %s]: failed to find organization id", request.GetMethod(), request.GetURL())
	}
	controller, err := c.appController.GetOrganizationController(organization)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find organization controller: %v", request.GetMethod(), request.GetURL(), err)
	}
	return controller, nil
}

func (c *ApiController) getRootCertificateController(request apitypes.IRequest) (appmodels.ICertificateController, error) {

	controller, err := c.getOrganizationController(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find organization controller: %v", request.GetMethod(), request.GetURL(), err)
	}

	rootSerialNumber, err := c.getRootSerialNumber(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find root serial number: %v", request.GetMethod(), request.GetURL(), err)
	}

	rootCertificateController, err := controller.GetCertificateController(rootSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find root certificate controller: %v", request.GetMethod(), request.GetURL(), err)
	}

	return rootCertificateController, nil
}

func (c *ApiController) getInnerCertificateController(request apitypes.IRequest) (appmodels.ICertificateController, error) {

	rootCertificateController, err := c.getRootCertificateController(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find root certificate controller: %v", request.GetMethod(), request.GetURL(), err)
	}

	serialNumber, err := c.getSerialNumber(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find inner serial number: %v", request.GetMethod(), request.GetURL(), err)
	}

	certificateController, err := rootCertificateController.GetChildCertificateController(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find inner certificate controller: %v", request.GetMethod(), request.GetURL(), err)
	}

	return certificateController, nil
}
