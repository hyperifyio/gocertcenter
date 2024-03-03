// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"

	"github.com/hyperifyio/gocertcenter/internal/app/apprepositories/filerepository"
)

func TestGetOrganizationDirectory(t *testing.T) {
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames"
	result := filerepository.OrganizationDirectory(dir, organization)
	assert.Equal(t, expected, result)
}

func TestGetOrganizationJsonPath(t *testing.T) {
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/organization.json"
	result := filerepository.OrganizationJsonPath(dir, organization)
	assert.Equal(t, expected, result)
}

func TestGetPrivateKeyPemPathWithTwoCertificates(t *testing.T) {
	certificate := appmodels.NewSerialNumber(456)
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/certificates/456/privkey.pem"
	result := filerepository.PrivateKeyPemPath(dir, organization, certificate)
	assert.Equal(t, expected, result)
}

func TestGetCertificatePemPathWithTwoCertificates(t *testing.T) {
	certificate := appmodels.NewSerialNumber(456)
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/certificates/456/cert.pem"
	result := filerepository.CertificatePemPath(dir, organization, certificate)
	assert.Equal(t, expected, result)
}

func TestGetCertificateDirectoryWithTwoCertificates(t *testing.T) {
	certificate := appmodels.NewSerialNumber(456)
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/certificates/456"
	result := filerepository.CertificateDirectory(dir, organization, certificate)
	assert.Equal(t, expected, result)
}

func TestGetPrivateKeyPemPathWithOneCertificate(t *testing.T) {
	certificate := appmodels.NewSerialNumber(123)
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/certificates/123/privkey.pem"
	result := filerepository.PrivateKeyPemPath(dir, organization, certificate)
	assert.Equal(t, expected, result)
}

func TestGetCertificatePemPathWithOneCertificate(t *testing.T) {
	certificate := appmodels.NewSerialNumber(123)
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/certificates/123/cert.pem"
	result := filerepository.CertificatePemPath(dir, organization, certificate)
	assert.Equal(t, expected, result)
}

func TestGetCertificateDirectoryWithOneCertificate(t *testing.T) {
	certificate := appmodels.NewSerialNumber(123)
	dir := "/data"
	organization := "HangoverGames"
	expected := "/data/organizations/HangoverGames/certificates/123"
	result := filerepository.CertificateDirectory(dir, organization, certificate)
	assert.Equal(t, expected, result)
}
