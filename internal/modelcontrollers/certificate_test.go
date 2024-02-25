// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers_test

import (
	"crypto/x509"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/hyperifyio/gocertcenter/internal/modelcontrollers"
	"github.com/hyperifyio/gocertcenter/internal/models"
	"math/big"
	"testing"
)

func TestNewCertificateController(t *testing.T) {

	mockService := &mocks.MockCertificateService{}
	controller := modelcontrollers.NewCertificateController(mockService)

	if !controller.UsesCertificateService(mockService) {
		t.Fatalf("Expected the certificate controller to use the mockService, got false")
	}

}

func TestCertificateController_GetExistingCertificate(t *testing.T) {
	expectedCert := models.NewCertificate(
		"TestOrg",
		[]models.ISerialNumber{models.NewSerialNumber(big.NewInt(4567))},
		&x509.Certificate{SerialNumber: big.NewInt(1234)},
	)
	mockService := &mocks.MockCertificateService{
		GetExistingCertificateFunc: func(organization string, certificates []models.ISerialNumber) (models.ICertificate, error) {
			if (len(certificates)) >= 1 {
				serialNumber := certificates[len(certificates)-1]
				if serialNumber.Cmp(expectedCert.GetSerialNumber()) == 0 {
					return expectedCert, nil
				}
			}
			return nil, nil // Simplified; in a real test, handle not found or error scenarios
		},
	}

	controller := modelcontrollers.NewCertificateController(mockService)
	cert, err := controller.GetExistingCertificate(
		"TestOrg",
		[]models.ISerialNumber{
			models.NewSerialNumber(big.NewInt(4567)),
			models.NewSerialNumber(big.NewInt(1234)),
		},
	)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
	if cert.GetSerialNumber().String() != expectedCert.GetSerialNumber().String() {
		t.Errorf("Expected certificate serial number %v, got %v", expectedCert.GetSerialNumber(), cert.GetSerialNumber())
	}
}

func TestCertificateController_CreateCertificate(t *testing.T) {
	newCert := models.NewCertificate(
		"NewOrg",
		[]models.ISerialNumber{models.NewSerialNumber(big.NewInt(4567))},
		&x509.Certificate{SerialNumber: big.NewInt(5678)},
	)
	mockService := &mocks.MockCertificateService{
		CreateCertificateFunc: func(certificate models.ICertificate) (models.ICertificate, error) {
			return certificate, nil // Echo back the input for simplicity
		},
	}

	controller := modelcontrollers.NewCertificateController(mockService)
	createdCert, err := controller.CreateCertificate(newCert)
	if err != nil {
		t.Fatalf("Did not expect an error, got %v", err)
	}
	if createdCert.GetSerialNumber().String() != newCert.GetSerialNumber().String() {
		t.Errorf("Expected certificate serial number %v, got %v", newCert.GetSerialNumber(), createdCert.GetSerialNumber())
	}
}
