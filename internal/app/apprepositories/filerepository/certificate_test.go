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

	organization := "TestOrg"
	serialNumbers := []appmodels.SerialNumber{
		appmodels.NewSerialNumber(big.NewInt(1)),
	}

	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)

	certPath := filerepository.GetCertificatePemPath(filePath, organization, serialNumbers)
	fmt.Println("Expected certificate path:", certPath)

	// Create a dummy certificate for testing
	cert := &x509.Certificate{
		SerialNumber: serialNumbers[0].Value(),
		Issuer:       pkix.Name{CommonName: "Test CA"},
		Subject:      pkix.Name{CommonName: "Test CertificateModel"},
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
	retrievedCert, err := repo.FindByOrganizationAndSerialNumbers(organization, serialNumbers)
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

	// Attempt to get an existing certificate with an empty serialNumbers slice
	organization := "TestOrg"
	certificates := []appmodels.SerialNumber{}

	retrievedCert, err := repo.FindByOrganizationAndSerialNumbers(organization, certificates)

	// Verify that an error is returned and that it contains the expected message
	assert.Error(t, err, "Expected an error when no certificate serial numbers are provided")
	assert.Nil(t, retrievedCert, "Expected no certificate to be returned")
	assert.EqualError(t, err, "no certificate serial numbers provided", "Error message should indicate that no serial numbers were provided")
}

func TestCertificateRepository_GetExistingCertificate_ReadFail(t *testing.T) {
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)
	fileManager := managers.NewFileManager()

	repo := filerepository.NewCertificateRepository(certManager, fileManager, tempDir)

	// Setup a scenario where the certificate file will not exist
	organization := "NonExistentOrg"
	serialNumbers := []appmodels.SerialNumber{appmodels.NewSerialNumber(big.NewInt(1))}

	// Attempt to get a certificate that does not exist, which should fail
	retrievedCert, err := repo.FindByOrganizationAndSerialNumbers(organization, serialNumbers)

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
		Subject:      pkix.Name{CommonName: "Test CertificateModel"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour), // 1 year
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	assert.NoError(t, err)

	cert, err := x509.ParseCertificate(certBytes)
	assert.NoError(t, err)

	mockCertificate := &appmocks.MockCertificate{}
	mockCertificate.On("GetCertificate").Return(cert, nil)
	mockCertificate.On("GetOrganizationID").Return("testOrg")
	mockCertificate.On("GetParents").Return([]appmodels.SerialNumber{})
	mockCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(template.SerialNumber))
	mockCertificate.On("GetID").Return("")

	// Attempt to save the certificate.
	_, err = repo.Save(mockCertificate)

	// Verify that the certificate was saved without error.
	assert.NoError(t, err)

	// You can extend this test to retrieve the saved certificate using FindByOrganizationAndSerialNumbers
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
	mockCertificate.On("GetCertificate").Return(&x509.Certificate{}, nil)
	mockCertificate.On("GetOrganizationID").Return("TestOrg")
	mockCertificate.On("GetParents").Return([]appmodels.SerialNumber{})
	mockCertificate.On("GetSerialNumber").Return(appmodels.NewSerialNumber(big.NewInt(1)))

	// Attempt to save the certificate, expecting a failure
	_, err = repo.Save(&mockCertificate)

	// Verify that an error is returned
	assert.Error(t, err, "Expected an error due to failed certificate save")
	assert.Contains(t, err.Error(), "failed to save certificate", "Error message should indicate a failure in saving the certificate")
}
