// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	swagger "github.com/davidebianchi/gswagger"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

type IApiController interface {
	GetInfo() *openapi3.Info
	GetRoutes() []apitypes.Route

	GetIndexDefinitions() swagger.Definitions
	GetIndex(response apitypes.IResponse, request apitypes.IRequest) error

	GetOrganizationCollection(response apitypes.IResponse, request apitypes.IRequest) error
	GetOrganizationCollectionDefinitions() swagger.Definitions

	GetRootCertificateCollection(response apitypes.IResponse, request apitypes.IRequest) error
	GetRootCertificateCollectionDefinitions() swagger.Definitions

	CreateRootCertificate(response apitypes.IResponse, request apitypes.IRequest) error
	CreateRootCertificateDefinitions() swagger.Definitions
}
