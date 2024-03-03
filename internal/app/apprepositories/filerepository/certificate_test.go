// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/fsutils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/filerepository"
)

func TestCertificateRepository_GetExistingCertificate(t *testing.T) {

	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)
	fileManager := managers.NewFileManager()

	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	filePath := tempDir
	repo := filerepository.NewCertificateRepository(certManager, fileManager, filePath)

	organization := big.NewInt(123)
	serialNumber := appmodels.NewSerialNumber(1)

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	certPath := filerepository.CertificatePemPath(filePath, organization, serialNumber)
	fmt.Println("Expected certificate path:", certPath)

	// Create a dummy certificate for testing
	cert := &x509.Certificate{
		SerialNumber: serialNumber,
		Issuer:       pkix.Name{CommonName: "Test CA"},
		Subject:      pkix.Name{CommonName: "Test Certificate"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &privKey.PublicKey, privKey)
	assert.NoError(t, err)

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certBytes})
	err = fsutils.SaveBytes(fileManager, certPath, certPEM, 0600, 0700)
	assert.NoError(t, err)

	// Test
	retrievedCert, err := repo.FindByOrganizationAndSerialNumber(organization, serialNumber)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedCert)
	// Add more assertions as necessary, e.g., comparing serial numbers, issuer names, etc.

}

func TestCertificateRepository_GetExistingCertificate_EmptySerialNumbers(t *testing.T) {
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	randomManager := commonmocks.NewMockRandomManager()
	fileManager := commonmocks.NewMockFileManager()
	certManager := commonmocks.NewMockCertificateManager()

	certManager.On("RandRandomManager").Return(randomManager)

	repo := filerepository.NewCertificateRepository(certManager, fileManager, tempDir)

	// Attempt to get an existing certificate with no serialNumber
	organization := big.NewInt(123)

	retrievedCert, err := repo.FindByOrganizationAndSerialNumber(organization, nil)

	// Verify that an error is returned and that it contains the expected message
	assert.Error(t, err, "Expected an error when no certificate serial numbers are provided")
	assert.Nil(t, retrievedCert, "Expected no certificate to be returned")
	assert.EqualError(t, err, "no certificate serial number provided", "Error message should indicate that no serial number were provided")
}

func TestCertificateRepository_GetExistingCertificate_ReadFail(t *testing.T) {
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)
	fileManager := managers.NewFileManager()

	repo := filerepository.NewCertificateRepository(certManager, fileManager, tempDir)

	// Setup a scenario where the certificate file will not exist
	organization := big.NewInt(0)
	serialNumber := appmodels.NewSerialNumber(1)

	// Attempt to get a certificate that does not exist, which should fail
	retrievedCert, err := repo.FindByOrganizationAndSerialNumber(organization, serialNumber)

	// Verify that an error is returned due to the failure in reading the certificate file
	assert.Error(t, err, "Expected an error due to failure in reading the certificate file")
	assert.Nil(t, retrievedCert, "Expected no certificate to be returned when read fails")

	// Optionally, you can check that the error message contains certain keywords, such as "failed to read certificate"
	assert.Contains(t, err.Error(), "failed to read certificate", "Error message should indicate a failure to read the certificate")
}

func TestCertificateRepository_CreateCertificate(t *testing.T) {

	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)
	fileManager := managers.NewFileManager()

	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	repo := filerepository.NewCertificateRepository(certManager, fileManager, tempDir)

	// Generate a new RSA private key.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	// Create a dummy certificate for testing.
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "Test Certificate"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 1 year
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	assert.NoError(t, err)

	cert, err := x509.ParseCertificate(certBytes)
	assert.NoError(t, err)

	mockCertificate := &appmocks.MockCertificate{}
	mockCertificate.On("Certificate").Return(cert, nil)
	mockCertificate.On("OrganizationID").Return(big.NewInt(123))
	mockCertificate.On("SerialNumber").Return(template.SerialNumber)
	mockCertificate.On("ID").Return("")

	// Attempt to save the certificate.
	_, err = repo.Save(mockCertificate)

	// Verify that the certificate was saved without error.
	assert.NoError(t, err)

	// You can extend this test to retrieve the saved certificate using FindByOrganizationAndSerialNumber
	// and verify its properties match those of the mockCertificate.
}

func TestCertificateRepository_CreateCertificate_SaveFail(t *testing.T) {

	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)
	fileManager := managers.NewFileManager()

	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	// Make the temp directory read-only to simulate a save failure.
	// Note: This approach might not work on Windows due to different permission models.
	err := os.Chmod(tempDir, 0400)
	if err != nil {
		t.Fatalf("Failed to make temp directory read-only: %v", err)
	}
	defer func() {
		// Attempt to restore write permissions to allow cleanup
		_ = os.Chmod(tempDir, 0700)
	}()

	repo := filerepository.NewCertificateRepository(certManager, fileManager, tempDir)

	// Use a mock or a simple certificate for testing
	mockCertificate := appmocks.MockCertificate{}
	mockCertificate.On("Certificate").Return(&x509.Certificate{}, nil)
	mockCertificate.On("OrganizationID").Return(big.NewInt(123))
	mockCertificate.On("SerialNumber").Return(appmodels.NewSerialNumber(1))

	// Attempt to save the certificate, expecting a failure
	_, err = repo.Save(&mockCertificate)

	// Verify that an error is returned
	assert.Error(t, err, "Expected an error due to failed certificate save")
	assert.Contains(t, err.Error(), "failed to save certificate", "Error message should indicate a failure in saving the certificate")
}
