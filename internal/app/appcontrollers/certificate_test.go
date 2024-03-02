// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"crypto/rsa"
	"crypto/x509"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewCertificateController(t *testing.T) {

	serialNumber := appmodels.NewSerialNumber(1)
	model := &appmocks.MockCertificate{}
	mockOrganizationController := &appmocks.MockOrganizationController{}
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}

	mockCertManager := &managers.SystemCertificateManager{}
	mockRandomManager := &commonmocks.MockRandomManager{}

	controller := appcontrollers.NewCertificateController(
		mockOrganizationController,
		nil,
		serialNumber,
		model,
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		time.Second,
	)

	if !controller.UsesCertificateService(mockCertificateRepository) {
		t.Fatalf("Expected the certificate controller to use the mockService, got false")
	}

}

func TestCertificateController_GetOrganizationController(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(1)
	model := &appmocks.MockCertificate{}
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}

	mockCertManager := &managers.SystemCertificateManager{}
	mockRandomManager := &commonmocks.MockRandomManager{}

	mockOrgController := new(appmocks.MockOrganizationController)
	controller := appcontrollers.NewCertificateController(
		mockOrgController,
		nil,
		serialNumber,
		model,
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		time.Second,
	)

	// Success test
	retrievedOrgController := controller.OrganizationController()
	assert.Equal(t, mockOrgController, retrievedOrgController)

	// Panic test
	assert.Panics(t, func() {

		_ = appcontrollers.NewCertificateController(
			nil,
			nil,
			serialNumber,
			model,
			mockCertificateRepository,
			mockPrivateKeyRepository,
			mockCertManager,
			mockRandomManager,
			time.Second,
		)

	}, "Expected panic when parent organization controller is nil")
}

func TestCertificateController_GetCertificateModel(t *testing.T) {
	serialNumber := appmodels.NewSerialNumber(1)
	mockCert := new(appmocks.MockCertificate)
	mockOrganizationController := &appmocks.MockOrganizationController{}
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}

	mockCertManager := &managers.SystemCertificateManager{}
	mockRandomManager := &commonmocks.MockRandomManager{}

	controller := appcontrollers.NewCertificateController(
		mockOrganizationController,
		nil,
		serialNumber,
		mockCert,
		mockCertificateRepository,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		time.Second,
	)

	retrievedCert := controller.Certificate()
	assert.Equal(t, mockCert, retrievedCert)
}

func TestCertificateController_GetChildCertificateCollection(t *testing.T) {
	mockCertRepo := new(appmocks.MockCertificateService)
	mockCert := new(appmocks.MockCertificate)
	serialNumber := appmodels.NewSerialNumber(123)
	mockCert.On("Parents").Return([]*big.Int{})
	certs := []appmodels.Certificate{mockCert}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}
	mockOrganizationController := &appmocks.MockOrganizationController{}
	mockOrganizationController.On("OrganizationID").Return("exmaple")

	mockCertManager := &managers.SystemCertificateManager{}
	mockRandomManager := &commonmocks.MockRandomManager{}

	mockCertRepo.On("FindAllByOrganizationAndSerialNumbers", mock.Anything, mock.Anything).Return(certs, nil)

	controller := appcontrollers.NewCertificateController(
		mockOrganizationController,
		nil,
		serialNumber,
		mockCert,
		mockCertRepo,
		mockPrivateKeyRepository,
		mockCertManager,
		mockRandomManager,
		time.Second,
	)

	// Success test
	retrievedCerts, err := controller.ChildCertificateCollection("")
	assert.NoError(t, err)
	assert.Equal(t, certs, retrievedCerts)

	// Test with certificateType filter (assuming implementation of FilterCertificatesByType)
	// You need to adjust based on actual implementation details
}

func TestCertificateController_NewIntermediateCertificate(t *testing.T) {
	orgID := "exampleOrg"
	commonName := "example.com"
	mockCert := new(appmocks.MockCertificate)
	mockPrivateKey := new(appmocks.MockPrivateKey)
	mockCertRepo := new(appmocks.MockCertificateService)
	mockPrivateKeyRepo := new(appmocks.MockPrivateKeyService)
	mockCertManager := new(commonmocks.MockCertificateManager)
	mockOrganization := new(appmocks.MockOrganization)
	mockRandomManager := new(commonmocks.MockRandomManager)
	mockOrgController := new(appmocks.MockOrganizationController)

	mockOrganization.On("ID").Return("example")
	mockOrganization.On("Name").Return("Example")
	mockOrganization.On("Names").Return([]string{"Example"})

	mockOrgController.On("OrganizationID").Return(orgID)
	mockOrgController.On("Organization").Return(mockOrganization)

	// Simulating serial number generation
	serialNumber := appmodels.NewSerialNumber(123)
	newSerialNumber := appmodels.NewSerialNumber(456)

	// Assuming we have a way to simulate or directly set the serial number for testing
	mockRandomManager.On("CreateBigInt", mock.Anything).Return(newSerialNumber, nil)

	mockCertManager.On("CreateCertificate", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return([]byte("certBytes"), nil)
	mockCertManager.On("ParseCertificate", []byte("certBytes")).Return(&x509.Certificate{SerialNumber: newSerialNumber}, nil)

	mockCert.On("Certificate").Return(&x509.Certificate{})
	mockCert.On("SerialNumber").Return(serialNumber)
	mockCert.On("Parents").Return([]*big.Int{})

	mockPrivateKey.On("PrivateKey").Return(&rsa.PrivateKey{})

	mockPrivateKeyRepo.On("FindByOrganizationAndSerialNumbers", orgID, []*big.Int{serialNumber}).Return(mockPrivateKey, nil)
	mockCertRepo.On("Save", mock.Anything).Return(mockCert, nil)

	controller := appcontrollers.NewCertificateController(
		mockOrgController,
		nil,
		serialNumber,
		mockCert,
		mockCertRepo,
		mockPrivateKeyRepo,
		mockCertManager,
		mockRandomManager,
		time.Hour*24,
	)

	// Test success path
	createdCert, createdPrivateKey, err := controller.NewIntermediateCertificate(commonName)
	assert.NoError(t, err)
	assert.NotNil(t, createdCert)
	assert.NotNil(t, createdPrivateKey)

	// Verify interactions
	mockCertManager.AssertExpectations(t)
	mockPrivateKeyRepo.AssertExpectations(t)
	mockCertRepo.AssertExpectations(t)
}
