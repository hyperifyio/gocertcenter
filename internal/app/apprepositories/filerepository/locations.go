// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository

import (
	"math/big"
	"path/filepath"
)

const (
	OrganizationsDirectoryName = "organizations"
	CertificatesDirectoryName  = "certificates"
	OrganizationJsonName       = "organization.json"
	CertificatePemName         = "cert.pem"
	PrivateKeyPemName          = "privkey.pem"
)

// OrganizationDirectory returns a path like `{dir}/organizations/{organization}`
func OrganizationDirectory(dir string, organization *big.Int) string {
	return filepath.Join(dir, OrganizationsDirectoryName, organization.String())
}

// OrganizationJsonPath returns a path like `{dir}/organizations/{organization}/organization.json`
func OrganizationJsonPath(dir string, organization *big.Int) string {
	return filepath.Join(dir, OrganizationsDirectoryName, organization.String(), OrganizationJsonName)
}

// PrivateKeyPemPath returns a path like `{dir}/organizations/{organization}/certificates/{certificate}/privkey.pem`
func PrivateKeyPemPath(dir string, organization, certificate *big.Int) string {
	return filepath.Join(CertificateDirectory(dir, organization, certificate), PrivateKeyPemName)
}

// CertificatePemPath returns a path like `{dir}/organizations/{organization}/certificates/{certificate}/cert.pem`
func CertificatePemPath(dir string, organization, certificate *big.Int) string {
	return filepath.Join(CertificateDirectory(dir, organization, certificate), CertificatePemName)
}

// CertificateDirectory returns a path like `{dir}/organizations/{organization}/certificates/{certificate}`
func CertificateDirectory(
	dir string,
	organization, certificate *big.Int,
) string {
	parts := []string{dir, OrganizationsDirectoryName, organization.String(), CertificatesDirectoryName, certificate.String()}
	return filepath.Join(parts...)
}
