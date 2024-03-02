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
func OrganizationDirectory(dir, organization string) string {
	return filepath.Join(dir, OrganizationsDirectoryName, organization)
}

// OrganizationJsonPath returns a path like `{dir}/organizations/{organization}/organization.json`
func OrganizationJsonPath(dir, organization string) string {
	return filepath.Join(dir, OrganizationsDirectoryName, organization, OrganizationJsonName)
}

// PrivateKeyPemPath returns a path like `{dir}/organizations/{organization}(/certificates/{certificate})+/privkey.pem`
func PrivateKeyPemPath(dir, organization string, certificates []*big.Int) string {
	return filepath.Join(CertificateDirectory(dir, organization, certificates), PrivateKeyPemName)
}

// CertificatePemPath returns a path like `{dir}/organizations/{organization}(/certificates/{certificate})+/cert.pem`
func CertificatePemPath(dir, organization string, certificates []*big.Int) string {
	return filepath.Join(CertificateDirectory(dir, organization, certificates), CertificatePemName)
}

// CertificateDirectory returns a path like `{dir}/organizations/{organization}(/certificates/{certificate})+`
func CertificateDirectory(
	dir, organization string,
	certificates []*big.Int,
) string {
	parts := []string{dir, OrganizationsDirectoryName, organization}
	for _, certificate := range certificates {
		parts = append(parts, CertificatesDirectoryName, certificate.String())
	}
	return filepath.Join(parts...)
}
