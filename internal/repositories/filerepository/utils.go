// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository

import (
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"os"
	"path/filepath"
)

// ReadCertificateFile reads a certificate from a PEM file
func ReadCertificateFile(fileName string) (*x509.Certificate, error) {

	// Read file contents
	data, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}

	// Decode PEM data
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the certificate")
	}

	// Parse the certificate
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse certificate: %w", err)
	}

	return cert, nil

}

// SaveCertificateFile saves the certificate to a PEM file
func SaveCertificateFile(
	fileName string,
	cert *x509.Certificate,
) error {
	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	if pemData == nil {
		return fmt.Errorf("failed to encode certificate to PEM")
	}
	return SaveFile(fileName, pemData)
}

// SaveFile saves bytes as a file by creating a temporary file first
func SaveFile(
	fileName string,
	pemData []byte,
) error {

	fileName = filepath.Clean(fileName)

	// Ensure the directory exists
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("failed to create a directory %s: %w", dir, err)
	}

	// Create a temporary file within the final file's directory
	tmpFile, err := os.CreateTemp(dir, "*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create a temporary file: %w", err)
	}
	defer tmpFile.Close()

	// Flag to check if rename was successful
	renameSuccessful := false

	// Cleanup function to remove the temporary file if rename fails
	defer func() {
		if !renameSuccessful {
			_ = os.Remove(tmpFile.Name()) // Ignore error here as it's cleanup
		}
	}()

	// Write the PEM data to the temporary file
	if _, err := tmpFile.Write(pemData); err != nil {
		return fmt.Errorf("failed to write to the temporary file: %w", err)
	}

	// Close the file to ensure all writes are flushed to disk
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close the temporary file: %w", err)
	}

	// Move the temporary file to the final location
	if err := os.Rename(tmpFile.Name(), fileName); err != nil {
		return fmt.Errorf("failed to move the temporary file to %s: %w", fileName, err)
	}

	renameSuccessful = true
	return nil

}

// SaveOrganizationJsonFile marshals data into JSON and saves it using SaveFile
func SaveOrganizationJsonFile(
	fileName string,
	dto dtos.OrganizationDTO,
) error {
	jsonData, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal organization data into JSON: %w", err)
	}
	return SaveFile(fileName, jsonData)
}

// ReadOrganizationJsonFile reads JSON data from a file and unmarshals it into the provided data structure.
func ReadOrganizationJsonFile(fileName string) (*dtos.OrganizationDTO, error) {

	fileData, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, fmt.Errorf("failed to read organization JSON file: %w", err)
	}

	// Initialize the DTO; use a pointer here
	dto := &dtos.OrganizationDTO{}
	if err := json.Unmarshal(fileData, dto); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization JSON data: %w", err)
	}

	return dto, nil
}
