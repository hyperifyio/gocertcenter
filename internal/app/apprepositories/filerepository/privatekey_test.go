// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmocks"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/filerepository"
	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/fsutils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestPrivateKeyRepository_GetExistingPrivateKey(t *testing.T) {

	fileManager := managers.NewFileManager()
	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)

	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	repo := filerepository.NewPrivateKeyRepository(certManager, fileManager, tempDir)
	organization := "TestOrg"
	certificates := []appmodels.SerialNumber{appmodels.NewSerialNumber(big.NewInt(1))}

	fileName := filerepository.PrivateKeyPemPath(tempDir, organization, certificates)

	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)
	rsaPrivBytes := x509.MarshalPKCS1PrivateKey(rsaPrivKey)

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: rsaPrivBytes})
	err = fsutils.SaveBytes(fileManager, fileName, keyPEM, 0600, 0700)
	assert.NoError(t, err)

	privateKey, err := repo.FindByOrganizationAndSerialNumbers(organization, certificates)
	assert.NoError(t, err)
	assert.NotNil(t, privateKey)

	// Perform more detailed assertions to verify the properties of privateKey
}

func TestPrivateKeyRepository_CreatePrivateKey(t *testing.T) {

	fileManager := managers.NewFileManager()
	randomManager := managers.NewRandomManager()
	certManager := managers.NewCertificateManager(randomManager)

	// Setup
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()

	repo := filerepository.NewPrivateKeyRepository(certManager, fileManager, tempDir)

	parentSerialNumber := appmodels.NewSerialNumber(big.NewInt(1))
	serialNumber := appmodels.NewSerialNumber(big.NewInt(2))

	rsaPrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.NoError(t, err)
	assert.NotNil(t, rsaPrivKey)

	mockPrivateKey := &appmocks.MockPrivateKey{}
	mockPrivateKey.On("OrganizationID").Return("TestOrg")
	mockPrivateKey.On("Parents").Return([]appmodels.SerialNumber{parentSerialNumber})
	mockPrivateKey.On("SerialNumber").Return(serialNumber)
	mockPrivateKey.On("PrivateKey").Return(rsaPrivKey)

	// Act
	createdKey, err := repo.Save(mockPrivateKey)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdKey)

	// Verify the file was created with the correct contents
	// This might involve reading the file directly and comparing the contents
}

func TestPrivateKeyRepository_GetExistingPrivateKey_Nonexistent(t *testing.T) {
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()
	certManager := &commonmocks.MockCertificateManager{}

	fileManager := managers.NewFileManager()

	repo := filerepository.NewPrivateKeyRepository(certManager, fileManager, tempDir)
	organization := "NonexistentOrg"
	certificates := []appmodels.SerialNumber{appmodels.NewSerialNumber(big.NewInt(1))}

	privateKey, err := repo.FindByOrganizationAndSerialNumbers(organization, certificates)
	assert.Error(t, err)
	assert.Nil(t, privateKey)
}

func TestPrivateKeyRepository_GetExistingPrivateKey_EmptyCertificates(t *testing.T) {
	tempDir, cleanup := setupTempDir(t)
	defer cleanup()
	certManager := &commonmocks.MockCertificateManager{}

	fileManager := managers.NewFileManager()

	repo := filerepository.NewPrivateKeyRepository(certManager, fileManager, tempDir)
	organization := "NonexistentOrg"
	var certificates []appmodels.SerialNumber

	privateKey, err := repo.FindByOrganizationAndSerialNumbers(organization, certificates)
	assert.Error(t, err)
	assert.Nil(t, privateKey)
}

func TestPrivateKeyRepository_GetFilePath(t *testing.T) {
	expectedFilePath := "/expected/file/path"
	certManager := &commonmocks.MockCertificateManager{}
	fileManager := managers.NewFileManager()
	repo := filerepository.NewPrivateKeyRepository(certManager, fileManager, expectedFilePath)

	actualFilePath := repo.FilePath()

	assert.Equal(t, expectedFilePath, actualFilePath, "The file path returned by FilePath should match the path set during repository creation")
}
