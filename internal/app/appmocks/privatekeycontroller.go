// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmocks

import (
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

// MockPrivateKeyController is a mock implementation of appmodels.PrivateKeyController for testing purposes.
type MockPrivateKeyController struct {
	mock.Mock
}

func (m *MockPrivateKeyController) GetApplicationController() appmodels.ApplicationController {
	args := m.Called()
	return args.Get(0).(appmodels.ApplicationController)
}

func (m *MockPrivateKeyController) GetOrganizationID() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockPrivateKeyController) GetOrganizationModel() appmodels.Organization {
	args := m.Called()
	return args.Get(0).(appmodels.Organization)
}

func (m *MockPrivateKeyController) GetOrganizationController() appmodels.OrganizationController {
	args := m.Called()
	return args.Get(0).(appmodels.OrganizationController)
}

func (m *MockPrivateKeyController) GetCertificateModel() appmodels.Certificate {
	args := m.Called()
	return args.Get(0).(appmodels.Certificate)
}

func (m *MockPrivateKeyController) GetCertificateController() appmodels.CertificateController {
	args := m.Called()
	return args.Get(0).(appmodels.CertificateController)
}

var _ appmodels.PrivateKeyController = (*MockPrivateKeyController)(nil)
