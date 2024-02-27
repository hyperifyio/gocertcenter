// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewCertificateController(t *testing.T) {

	serialNumber := appmodels.NewSerialNumber(big.NewInt(1))

	model := &appmocks.MockCertificate{}

	mockOrganizationController := &appmocks.MockOrganizationController{}
	mockCertificateRepository := &appmocks.MockCertificateService{}
	mockPrivateKeyRepository := &appmocks.MockPrivateKeyService{}

	mockCertManager := &managers.CertificateManager{}
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
