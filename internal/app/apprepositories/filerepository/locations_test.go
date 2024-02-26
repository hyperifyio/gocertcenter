// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/filerepository"
)

func TestGetOrganizationDirectory(t *testing.T) {
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames"
	result := filerepository.GetOrganizationDirectory(dir, organization)
	assert.Equal(t, expected, result)
}

func TestGetOrganizationJsonPath(t *testing.T) {
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/organization.json"
	result := filerepository.GetOrganizationJsonPath(dir, organization)
	assert.Equal(t, expected, result)
}

func TestGetPrivateKeyPemPathWithTwoCertificates(t *testing.T) {
	a := appmodels.NewSerialNumber(big.NewInt(123))
	b := appmodels.NewSerialNumber(big.NewInt(456))
	dir := "/data"
	organization := "HangoverGames"
	certificates := []appmodels.ISerialNumber{a, b}
	expected := "/data/organizations/HangoverGames/certificates/123/certificates/456/privkey.pem"
	result := filerepository.GetPrivateKeyPemPath(dir, organization, certificates)
	assert.Equal(t, expected, result)
}

func TestGetCertificatePemPathWithTwoCertificates(t *testing.T) {
	a := appmodels.NewSerialNumber(big.NewInt(123))
	b := appmodels.NewSerialNumber(big.NewInt(456))
	dir := "/data"
	organization := "HangoverGames"
	certificates := []appmodels.ISerialNumber{a, b}
	expected := "/data/organizations/HangoverGames/certificates/123/certificates/456/cert.pem"
	result := filerepository.GetCertificatePemPath(dir, organization, certificates)
	assert.Equal(t, expected, result)
}

func TestGetCertificateDirectoryWithTwoCertificates(t *testing.T) {
	a := appmodels.NewSerialNumber(big.NewInt(123))
	b := appmodels.NewSerialNumber(big.NewInt(456))
	dir := "/data"
	organization := "HangoverGames"
	certificates := []appmodels.ISerialNumber{a, b}
	expected := "/data/organizations/HangoverGames/certificates/123/certificates/456"
	result := filerepository.GetCertificateDirectory(dir, organization, certificates)
	assert.Equal(t, expected, result)
}

func TestGetPrivateKeyPemPathWithOneCertificate(t *testing.T) {
	a := appmodels.NewSerialNumber(big.NewInt(123))
	dir := "/data"
	organization := "HangoverGames"
	certificates := []appmodels.ISerialNumber{a}
	expected := "/data/organizations/HangoverGames/certificates/123/privkey.pem"
	result := filerepository.GetPrivateKeyPemPath(dir, organization, certificates)
	assert.Equal(t, expected, result)
}

func TestGetCertificatePemPathWithOneCertificate(t *testing.T) {
	a := appmodels.NewSerialNumber(big.NewInt(123))
	dir := "/data"
	organization := "HangoverGames"
	certificates := []appmodels.ISerialNumber{a}
	expected := "/data/organizations/HangoverGames/certificates/123/cert.pem"
	result := filerepository.GetCertificatePemPath(dir, organization, certificates)
	assert.Equal(t, expected, result)
}

func TestGetCertificateDirectoryWithOneCertificate(t *testing.T) {
	a := appmodels.NewSerialNumber(big.NewInt(123))
	dir := "/data"
	organization := "HangoverGames"
	certificates := []appmodels.ISerialNumber{a}
	expected := "/data/organizations/HangoverGames/certificates/123"
	result := filerepository.GetCertificateDirectory(dir, organization, certificates)
	assert.Equal(t, expected, result)
}
