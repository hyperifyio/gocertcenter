// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package filerepository

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// readCertificateFile reads a certificate from a PEM file
func readCertificateFile(fileName string) (*x509.Certificate, error) {

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

// saveCertificateFile saves the certificate as a PEM file
func saveCertificateFile(
	fileName string,
	cert *x509.Certificate,
) error {
	// Convert the certificate to PEM format
	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	})
	if pemData == nil {
		return fmt.Errorf("failed to encode certificate to PEM")
	}
	return saveFile(fileName, pemData)
}

// saveFile saves a bytes as a PEM file
func saveFile(
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
