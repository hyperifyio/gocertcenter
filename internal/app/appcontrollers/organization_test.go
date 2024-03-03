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
	mockOrganization.On("ID").Return(organizationID)

	// Simulate existing serial number
	mockCertificateRepository.On("FindByOrganizationAndSerialNumber", organizationID, mock.Anything).Return(&appmocks.MockCertificate{}, nil)

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
	serialNumber := appmodels.NewSerialNumber(12345)

	// Simulate failure to fetch the certificate model
	mockCertificateRepository.On("FindByOrganizationAndSerialNumber", mock.Anything, serialNumber).Return(nil, fmt.Errorf("fetch fail"))

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

	_, err := controller.CertificateController(serialNumber)
	if err == nil || !strings.Contains(err.Error(), "failed") {
		t.Errorf("Expected a failure fetching certificate model, got: %v", err)
	}
}

func TestOrganizationController_GetCertificateCollection_Success(t *testing.T) {
	organizationID := "testOrg"
	mockCertificateRepository := new(appmocks.MockCertificateService)
	mockCertificates := []appmodels.Certificate{
		&appmocks.MockCertificate{}, // Assume this is properly initialized in your setup
	}

	// Setup mock to return a list of certificates
	mockCertificateRepository.On("FindAllByOrganization", organizationID).Return(mockCertificates, nil)

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

	certificates, err := controller.CertificateCollection()
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

	_, err := controller.CertificateCollection()
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

	model := controller.Organization()
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

	parent := controller.ApplicationController()
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

func TestOrganizationController_GetCertificateCollection_Failure(t *testing.T) {
	organizationID := "testOrg"
	expectedErr := fmt.Errorf("database error")

	// Mocking Organization and its repositories
	mockOrganization := &appmocks.MockOrganization{}
	mockCertificateRepository := new(appmocks.MockCertificateService)
	mockPrivateKeyRepository := new(appmocks.MockPrivateKeyService)
	mockCertManager := commonmocks.NewMockCertificateManager()
	mockRandomManager := commonmocks.NewMockRandomManager()
	mockApplicationController := new(appmocks.MockApplicationController)

	// Setup: CertOrganizationController with mocked dependencies
	controller := appcontrollers.NewOrganizationController(
		organizationID,
		mockOrganization,
		new(appmocks.MockOrganizationService),
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		24*time.Hour,
		mockApplicationController,
	)

	// Mock the behavior of FindAllByOrganizationAndSignedBy to return an error
	mockCertificateRepository.On("FindAllByOrganization", organizationID).Return(nil, expectedErr)

	// Execute: Call the method we're testing
	certificates, err := controller.CertificateCollection()

	// Assert: Check that the error is as expected and no certificates are returned
	assert.Nil(t, certificates, "Expected no certificates to be returned on error")
	assert.Error(t, err, "Expected an error to be returned")
	assert.Contains(t, err.Error(), "database error", "Expected the error to be propagated from the certificate repository")

	// Verify that the mocked method was called with expected parameters
	mockCertificateRepository.AssertCalled(t, "FindAllByOrganization", organizationID)
}

func TestOrganizationController_GetCertificateController(t *testing.T) {
	organizationID := "testOrg"
	serialNumber := appmodels.NewSerialNumber(12345)

	// Mocking Organization and its repositories
	mockCertificate := &appmocks.MockCertificate{}
	mockOrganization := &appmocks.MockOrganization{}
	mockCertificateRepository := new(appmocks.MockCertificateService)
	mockPrivateKeyRepository := new(appmocks.MockPrivateKeyService)
	mockCertManager := commonmocks.NewMockCertificateManager()
	mockRandomManager := commonmocks.NewMockRandomManager()
	mockApplicationController := new(appmocks.MockApplicationController)

	// Setup: CertOrganizationController with mocked dependencies
	controller := appcontrollers.NewOrganizationController(
		organizationID,
		mockOrganization,
		new(appmocks.MockOrganizationService),
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		24*time.Hour,
		mockApplicationController,
	)

	// Mock the behavior of FindByOrganizationAndSerialNumber to return an error
	mockCertificateRepository.On("FindByOrganizationAndSerialNumber", organizationID, serialNumber).Return(mockCertificate, nil)

	// Execute: Call the method we're testing
	certificateController, err := controller.CertificateController(serialNumber)

	// Assert: Check that the error is as expected and no certificate controller is returned
	assert.NotNil(t, certificateController, "Expected no certificate controller to be returned on error")
	assert.NoError(t, err, "Expected success")

	// Verify that the mocked method was called with expected parameters
	mockCertificateRepository.AssertCalled(t, "FindByOrganizationAndSerialNumber", organizationID, serialNumber)
}

func TestNewRootCertificate_SerialNumberGenerationFail(t *testing.T) {
	// Mock dependencies
	mockRandomManager := new(commonmocks.MockRandomManager)
	mockOrganization := new(appmocks.MockOrganization)
	mockCertificateRepository := new(appmocks.MockCertificateService)
	mockPrivateKeyRepository := new(appmocks.MockPrivateKeyService)
	mockCertManager := commonmocks.NewMockCertificateManager()
	organizationID := "testOrg"

	// Setup the CertOrganizationController with mocked dependencies
	controller := appcontrollers.NewOrganizationController(
		organizationID,
		mockOrganization,
		new(appmocks.MockOrganizationService),
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	// Simulate failure in GenerateSerialNumber
	mockRandomManager.On("CreateBigInt", mock.Anything).Return(nil, fmt.Errorf("random generation fail"))

	_, err := controller.NewRootCertificate("Common Name")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create serial number")
}

func TestNewRootCertificate_NoCertificateRepository(t *testing.T) {
	// Setup the CertOrganizationController without a certificate repository
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		new(appmocks.MockOrganization),
		new(appmocks.MockOrganizationService),
		nil, // No certificate repository
		new(appmocks.MockPrivateKeyService),
		commonmocks.NewMockCertificateManager(),
		commonmocks.NewMockRandomManager(),
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	_, err := controller.NewRootCertificate("Common Name")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no certificate repository")
}

func TestNewRootCertificate_NoPrivateKeyRepository(t *testing.T) {

	// Setup the CertOrganizationController without a private key repository
	controller := appcontrollers.NewOrganizationController(
		"testOrg",
		new(appmocks.MockOrganization),
		new(appmocks.MockOrganizationService),
		new(appmocks.MockCertificateService),
		nil, // No private key repository
		commonmocks.NewMockCertificateManager(),
		commonmocks.NewMockRandomManager(),
		24*time.Hour,
		new(appmocks.MockApplicationController),
	)

	_, err := controller.NewRootCertificate("Common Name")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "no certificate repository")
}
