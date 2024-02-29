// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/app/apputils"
	"github.com/hyperifyio/gocertcenter/internal/common/fsutils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// ReadCertificateFile reads a certificate from a PEM file
func ReadCertificateFile(
	fileManager managers.FileManager,
	certManager managers.CertificateManager,
	fileName string,
) (*x509.Certificate, error) {

	// Read file contents
	data, err := fileManager.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}

	// Decode PEM data
	block, _ := certManager.DecodePEM(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the certificate")
	}

	// Parse the certificate
	cert, err := certManager.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}

	return cert, nil

}

// SaveCertificateFile saves the certificate to a PEM file
func SaveCertificateFile(
	fileManager managers.FileManager,
	certManager managers.CertificateManager,
	fileName string,
	cert *x509.Certificate,
) error {
	pemData := certManager.EncodePEMToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	if pemData == nil {
		return fmt.Errorf("failed to encode certificate to PEM")
	}
	return fsutils.SaveBytes(fileManager, fileName, pemData, 0600, 0700)
}

// SaveOrganizationJsonFile marshals data into JSON and saves it using fileManager.SaveBytes
func SaveOrganizationJsonFile(
	fileManager managers.FileManager,
	fileName string,
	dto appdtos.OrganizationDTO,
) error {
	jsonData, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal organization data into JSON: %w", err)
	}
	return fsutils.SaveBytes(fileManager, fileName, jsonData, 0600, 0700)
}

// ReadOrganizationJsonFile reads JSON data from a file and unmarshals it into the provided data structure.
func ReadOrganizationJsonFile(fileManager managers.FileManager, fileName string) (*appdtos.OrganizationDTO, error) {

	fileData, err := fileManager.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read organization JSON file: %w", err)
	}

	dto := &appdtos.OrganizationDTO{}
	if err := json.Unmarshal(fileData, dto); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization JSON data: %w", err)
	}

	return dto, nil
}

// ReadPrivateKeyFile reads a private key from a PEM file
func ReadPrivateKeyFile(
	fileManager managers.FileManager,
	certManager managers.CertificateManager,
	fileName string,
) (any, appmodels.KeyType, error) {

	data, err := fileManager.ReadFile(fileName)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("ReadPrivateKeyFile: failed to read certificate file: %w", err)
	}

	privateKey, keyType, err := apputils.ParsePrivateKeyFromPEMBytes(certManager, data)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("ReadPrivateKeyFile: failed to decode PEM block containing the private key: %w", err)
	}

	return privateKey, keyType, nil
}
