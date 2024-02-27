// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// ApiController implements apitypes.IAppController and IApiController
type ApiController struct {
	server        apitypes.IServer
	appController appmodels.IApplicationController
	certManager   managers.ICertificateManager
}

func NewApiController(
	server apitypes.IServer,
	appController appmodels.IApplicationController,
	certManager managers.ICertificateManager,
) *ApiController {
	return &ApiController{
		server:        server,
		appController: appController,
		certManager:   certManager,
	}
}

var _ apitypes.IAppController = (*ApiController)(nil)
var _ IApiController = (*ApiController)(nil)
