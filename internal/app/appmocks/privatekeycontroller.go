// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockPrivateKeyController is a mock implementation of appmodels.IPrivateKeyController for testing purposes.
type MockPrivateKeyController struct {
	mock.Mock
}

func (m *MockPrivateKeyController) GetApplicationController() appmodels.IApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.IApplicationController)
}

func (m *MockPrivateKeyController) GetOrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockPrivateKeyController) GetOrganizationModel() appmodels.IOrganization {
	args := m.Called()
	return args.Get(0).(appmodels.IOrganization)
}

func (m *MockPrivateKeyController) GetOrganizationController() appmodels.IOrganizationController {
	args := m.Called()
	return args.Get(0).(appmodels.IOrganizationController)
}

func (m *MockPrivateKeyController) GetCertificateModel() appmodels.ICertificate {
	args := m.Called()
	return args.Get(0).(appmodels.ICertificate)
}

func (m *MockPrivateKeyController) GetCertificateController() appmodels.ICertificateController {
	args := m.Called()
	return args.Get(0).(appmodels.ICertificateController)
}

var _ appmodels.IPrivateKeyController = (*MockPrivateKeyController)(nil)
