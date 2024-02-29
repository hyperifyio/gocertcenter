// Copyright (c) 2024. Heusala roup Oy <info@heusalagroup.fi>. All rights reserved.

package appcontrollers_test

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewOrganizationController(t *testing.T) {

	organization := "testorg"
	model := &appmocks.MockOrganization{}
	mockOrganizationRepository := &appmocks.MockOrganizationService{}
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}
	certManager := commonmocks.NewMockCertificateManager()
	randomManager := commonmocks.NewMockRandomManager()

	controller := appcontrollers.NewOrganizationController(
		organization,
		model,
		mockOrganizationRepository,
		mockCertificateRepository,
		mockPrivateKeyRepository,
		certManager,
		randomManager,
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	if !controller.UsesOrganizationService(mockOrganizationRepository) {
		t.Fatalf("Expected the organization controller to use the mockService, got false")
	}

}

func TestOrganizationController_NewRootCertificate_SerialNumberExists(t *testing.T) {
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}
	mockCertManager := commonmocks.NewMockCertificateManager()
	mockRandomManager := commonmocks.NewMockRandomManager()

	mockRandomManager.On("CreateBigInt", mock.Anything).Return(big.NewInt(123), nil)

	mockOrganization := &appmocks.MockOrganization{}
	organizationID := "testorg"
	mockOrganization.On("GetID").Return(organizationID)

	// Simulate existing serial number
	mockCertificateRepository.On("FindByOrganizationAndSerialNumbers", organizationID, mock.Anything).Return(&appmocks.MockCertificate{}, nil)

	controller := appcontrollers.NewOrganizationController(
		organizationID,
		mockOrganization,
		&appmocks.MockOrganizationService{},
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	_, err := controller.NewRootCertificate("Common Name")
	if err == nil || !strings.Contains(err.Error(), "serial number exists already") {
		t.Errorf("Expected an error about existing serial number, got: %v", err)
	}
}

func TestOrganizationController_GetCertificateController_FetchFail(t *testing.T) {
	mockCertificateRepository := &appmocks.MockCertificateService{}
	serialNumber := appmodels.NewSerialNumber(big.NewInt(12345))

	// Simulate failure to fetch the certificate model
	mockCertificateRepository.On("FindByOrganizationAndSerialNumbers", mock.Anything, []appmodels.ISerialNumber{serialNumber}).Return(nil, fmt.Errorf("fetch fail"))

	controller := appcontrollers.NewOrganizationController(
		"testorg",
		&appmocks.MockOrganization{},
		&appmocks.MockOrganizationService{},
		mockCertificateRepository,
		&appmocks.MockPrivateKeyService{},
		commonmocks.NewMockCertificateManager(),
		commonmocks.NewMockRandomManager(),
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	_, err := controller.GetCertificateController(serialNumber)
	if err == nil || !strings.Contains(err.Error(), "failed") {
		t.Errorf("Expected a failure fetching certificate model, got: %v", err)
	}
}

func TestOrganizationController_GetCertificateCollection_Success(t *testing.T) {
	organizationID := "testOrg"
	mockCertificateRepository := new(appmocks.MockCertificateService)
	mockCertificates := []appmodels.ICertificate{
		&appmocks.MockCertificate{}, // Assume this is properly initialized in your setup
	}

	// Setup mock to return a list of certificates
	mockCertificateRepository.On("FindAllByOrganizationAndSerialNumbers", organizationID, []appmodels.ISerialNumber{}).Return(mockCertificates, nil)

	controller := appcontrollers.NewOrganizationController(
		organizationID,
		&appmocks.MockOrganization{},
		&appmocks.MockOrganizationService{},
		mockCertificateRepository,
		&appmocks.MockPrivateKeyService{},
		commonmocks.NewMockCertificateManager(),
		commonmocks.NewMockRandomManager(),
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	certificates, err := controller.GetCertificateCollection()
	assert.NoError(t, err)
	assert.Equal(t, mockCertificates, certificates, "The returned certificates should match the mock certificates")
	mockCertificateRepository.AssertExpectations(t)
}

func TestOrganizationController_GetCertificateCollection_NoCertificateRepository(t *testing.T) {
	organizationID := "testOrg"

	controller := appcontrollers.NewOrganizationController(
		organizationID,
		&appmocks.MockOrganization{},
		&appmocks.MockOrganizationService{},
		nil, // No certificate repository provided
		&appmocks.MockPrivateKeyService{},
		commonmocks.NewMockCertificateManager(),
		commonmocks.NewMockRandomManager(),
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	_, err := controller.GetCertificateCollection()
	if err == nil {
		t.Fatal("Expected an error due to no certificate repository, got nil")
	}
	assert.Contains(t, err.Error(), "no certificate repository", "Error message should indicate the absence of a certificate repository")
}

func TestOrganizationController_GetOrganizationModel(t *testing.T) {
	mockModel := new(appmocks.MockOrganization)
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		mockModel, // This is the model we expect to retrieve
		nil, nil, nil, nil, nil, 0,
		new(appmocks.MockApplicationController),
	)

	model := controller.GetOrganizationModel()
	assert.Equal(t, mockModel, model, "Expected organization model does not match")
}

func TestOrganizationController_GetApplicationController(t *testing.T) {
	mockParent := new(appmocks.MockApplicationController)
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		nil,
		nil, nil, nil,
		nil, nil,
		0,
		mockParent,
	)

	parent := controller.GetApplicationController()
	assert.Equal(t, mockParent, parent, "Expected application controller does not match")
}

func TestOrganizationController_SetExpirationDuration(t *testing.T) {
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		nil,
		nil, nil, nil,
		nil, nil,
		24*time.Hour, // initial duration
		new(appmocks.MockApplicationController),
	)

	newDuration := 48 * time.Hour
	controller.SetExpirationDuration(newDuration)

	assert.Equal(t, newDuration, controller.ExpirationDuration(), "Expiration duration was not updated correctly")
}

func TestOrganizationController_UsesApplicationController(t *testing.T) {
	mockParent := new(appmocks.MockApplicationController)
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		nil,
		nil, nil, nil,
		nil, nil,
		0,
		mockParent,
	)

	assert.True(t, controller.UsesApplicationController(mockParent), "Controller should use the provided application controller")
	assert.False(t, controller.UsesApplicationController(new(appmocks.MockApplicationController)), "Controller should not use a different application controller")
}

func TestOrganizationController_RevokeCertificate_NotImplemented(t *testing.T) {
	t.Skip("RevokeCertificate method is not implemented yet")

	// The following is a placeholder structure for the future test
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		nil,
		nil, nil, nil,
		nil, nil,
		0,
		new(appmocks.MockApplicationController),
	)
	mockCertificate := new(appmocks.MockCertificate)
	// Assume mockCertificate setup is done here

	_, err := controller.RevokeCertificate(mockCertificate)
	assert.Error(t, err, "Expected not implemented error for RevokeCertificate")
	// Further assertions once implemented
}
