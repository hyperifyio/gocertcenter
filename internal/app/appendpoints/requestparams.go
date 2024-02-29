// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"strings"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *HttpApiController) requestOrganization(request apitypes.Request) string {
	organization := request.Variable("organization")
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

func (c *HttpApiController) rootSerialNumber(request apitypes.Request) (appmodels.SerialNumber, error) {
	serialNumberString := request.Variable("rootSerialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to parse rootSerialNumber: %v", request.Method(), request.URL(), err)
	}
	c.logf(request, "rootSerialNumber = %s", serialNumber.String())
	return serialNumber, nil
}

func (c *HttpApiController) serialNumber(request apitypes.Request) (appmodels.SerialNumber, error) {
	serialNumberString := request.Variable("serialNumber")
	serialNumber, err := appmodels.ParseSerialNumber(serialNumberString, 10)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to parse serialNumber: %v", request.Method(), request.URL(), err)
	}
	c.logf(request, "serialNumber = %s", serialNumber.String())
	return serialNumber, nil
}

func (c *HttpApiController) organizationController(request apitypes.Request) (appmodels.OrganizationController, error) {
	organization := c.requestOrganization(request)
	if organization == "" {
		return nil, fmt.Errorf("[%s %s]: failed to find organization id", request.Method(), request.URL())
	}
	controller, err := c.appController.OrganizationController(organization)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find organization controller: %v", request.Method(), request.URL(), err)
	}
	return controller, nil
}

func (c *HttpApiController) rootCertificateController(request apitypes.Request) (appmodels.CertificateController, error) {

	controller, err := c.organizationController(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find organization controller: %v", request.Method(), request.URL(), err)
	}

	rootSerialNumber, err := c.rootSerialNumber(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find root serial number: %v", request.Method(), request.URL(), err)
	}

	rootCertificateController, err := controller.CertificateController(rootSerialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find root certificate controller: %v", request.Method(), request.URL(), err)
	}

	return rootCertificateController, nil
}

func (c *HttpApiController) innerCertificateController(request apitypes.Request) (appmodels.CertificateController, error) {

	rootCertificateController, err := c.rootCertificateController(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find root certificate controller: %v", request.Method(), request.URL(), err)
	}

	serialNumber, err := c.serialNumber(request)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find inner serial number: %v", request.Method(), request.URL(), err)
	}

	certificateController, err := rootCertificateController.ChildCertificateController(serialNumber)
	if err != nil {
		return nil, fmt.Errorf("[%s %s]: failed to find inner certificate controller: %v", request.Method(), request.URL(), err)
	}

	return certificateController, nil
}
