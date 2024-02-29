// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository

import (
	"path/filepath"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

const (
	OrganizationsDirectoryName = "organizations"
	CertificatesDirectoryName  = "certificates"
	OrganizationJsonName       = "organization.json"
	CertificatePemName         = "cert.pem"
	PrivateKeyPemName          = "privkey.pem"
)

// GetOrganizationDirectory returns a path like `{dir}/organizations/{organization}`
func GetOrganizationDirectory(dir, organization string) string {
	return filepath.Join(dir, OrganizationsDirectoryName, organization)
}

// GetOrganizationJsonPath returns a path like `{dir}/organizations/{organization}/organization.json`
func GetOrganizationJsonPath(dir, organization string) string {
	return filepath.Join(dir, OrganizationsDirectoryName, organization, OrganizationJsonName)
}

// GetPrivateKeyPemPath returns a path like `{dir}/organizations/{organization}(/certificates/{certificate})+/privkey.pem`
func GetPrivateKeyPemPath(dir, organization string, certificates []appmodels.SerialNumber) string {
	return filepath.Join(GetCertificateDirectory(dir, organization, certificates), PrivateKeyPemName)
}

// GetCertificatePemPath returns a path like `{dir}/organizations/{organization}(/certificates/{certificate})+/cert.pem`
func GetCertificatePemPath(dir, organization string, certificates []appmodels.SerialNumber) string {
	return filepath.Join(GetCertificateDirectory(dir, organization, certificates), CertificatePemName)
}

// GetCertificateDirectory returns a path like `{dir}/organizations/{organization}(/certificates/{certificate})+`
func GetCertificateDirectory(
	dir, organization string,
	certificates []appmodels.SerialNumber,
) string {
	parts := []string{dir, OrganizationsDirectoryName, organization}
	for _, certificate := range certificates {
		parts = append(parts, CertificatesDirectoryName, certificate.String())
	}
	return filepath.Join(parts...)
}
