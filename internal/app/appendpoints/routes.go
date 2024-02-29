// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"net/http"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *HttpApiController) Routes() []apitypes.Route {
	return []apitypes.Route{
		{
			Method:      http.MethodDelete,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/certificates/{serialNumber}",
			Handler:     c.RevokeCertificate,
			Definitions: c.RevokeCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/certificates/{serialNumber}",
			Handler:     c.Certificate,
			Definitions: c.CertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/certificates",
			Handler:     c.CertificateCollection,
			Definitions: c.CertificateCollectionDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/certificates",
			Handler:     c.CreateCertificate,
			Definitions: c.CreateCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/crl",
			Handler:     c.CertificateRevocationList,
			Definitions: c.CertificateRevocationListDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}/crl",
			Handler:     c.UpdateCertificateRevocationList,
			Definitions: c.UpdateCertificateRevocationListDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates/{rootSerialNumber}",
			Handler:     c.RootCertificate,
			Definitions: c.RootCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}/certificates",
			Handler:     c.RootCertificateCollection,
			Definitions: c.RootCertificateCollectionDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations/{organization}/certificates",
			Handler:     c.CreateRootCertificate,
			Definitions: c.CreateRootCertificateDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations/{organization}",
			Handler:     c.Organization,
			Definitions: c.OrganizationDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/organizations",
			Handler:     c.OrganizationCollection,
			Definitions: c.OrganizationCollectionDefinitions(),
		},
		{
			Method:      http.MethodPost,
			Path:        "/organizations",
			Handler:     c.CreateOrganization,
			Definitions: c.CreateOrganizationDefinitions(),
		},
		{
			Method:      http.MethodGet,
			Path:        "/",
			Handler:     c.Index,
			Definitions: c.IndexDefinitions(),
		},
	}
}

var _ apitypes.ApplicationRoutesFunc = (*HttpApiController)(nil).Routes
