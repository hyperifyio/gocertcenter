// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appcontrollers_test

import (
	"testing"
	"time"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
)

func TestNewCertificateController(t *testing.T) {

	mockCertificateService := &appmocks.MockCertificateService{}
	mockCertManager := &managers.CertificateManager{}
	mockRandomManager := &commonmocks.MockRandomManager{}
	controller := appcontrollers.NewCertificateController(
		mockCertificateService,
		mockCertManager,
		mockRandomManager,
		time.Second,
	)

	if !controller.UsesCertificateService(mockCertificateService) {
		t.Fatalf("Expected the certificate controller to use the mockService, got false")
	}

}
