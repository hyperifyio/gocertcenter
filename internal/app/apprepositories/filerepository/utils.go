// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperifyio/gocertcenter/internal/app/appdtos"
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
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
	dto appdtos.OrganizationDTO,
) error {
	jsonData, err := json.MarshalIndent(dto, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal organization data into JSON: %w", err)
	}
	return SaveFile(fileName, jsonData)
}

// ReadOrganizationJsonFile reads JSON data from a file and unmarshals it into the provided data structure.
func ReadOrganizationJsonFile(fileName string) (*appdtos.OrganizationDTO, error) {

	fileData, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, fmt.Errorf("failed to read organization JSON file: %w", err)
	}

	// Initialize the DTO; use a pointer here
	dto := &appdtos.OrganizationDTO{}
	if err := json.Unmarshal(fileData, dto); err != nil {
		return nil, fmt.Errorf("failed to unmarshal organization JSON data: %w", err)
	}

	return dto, nil
}

// ReadPrivateKeyFile reads a private key from a PEM file
func ReadPrivateKeyFile(fileName string) (any, appmodels.KeyType, error) {

	// Read file contents
	data, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("failed to read certificate file: %w", err)
	}

	// Decode PEM data
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, appmodels.NIL_KEY_TYPE, errors.New("failed to decode PEM block containing the certificate")
	}

	// Parse the certificate
	if block.Type == "PRIVATE KEY" {
		return parsePKCS8PrivateKey(block.Bytes)
	} else if block.Type == "RSA PRIVATE KEY" {
		return parseRSAPrivateKey(block.Bytes)
	} else if block.Type == "EC PRIVATE KEY" {
		return parseECPrivateKey(block.Bytes)
	}
	return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("unsupported key type or invalid key: %w", err)

}

func parsePKCS8PrivateKey(bytes []byte) (any, appmodels.KeyType, error) {
	privateKey, err := x509.ParsePKCS8PrivateKey(bytes)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("failed to parse private key: %w", err)
	}
	keyType, err := determineKeyType(privateKey)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("could not detect key type: %w", err)
	}
	return privateKey, keyType, nil
}

func determineKeyType(privateKey any) (appmodels.KeyType, error) {
	switch key := privateKey.(type) {

	case *rsa.PrivateKey:
		keyType, err := determineRSATypeFromKey(key)
		if err != nil {
			return appmodels.NIL_KEY_TYPE, fmt.Errorf("determineKeyType: could not detect RSA key type: %w", err)
		}
		return keyType, nil

	case *ecdsa.PrivateKey:
		keyType, err := determineECDSACurve(key.Curve)
		if err != nil {
			return appmodels.NIL_KEY_TYPE, fmt.Errorf("determineKeyType: could not detect ecdsa key type: %w", err)
		}
		return keyType, nil

	case ed25519.PrivateKey:
		return appmodels.Ed25519, nil

	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("determineKeyType: unknown or unsupported key type")
	}
}

func readRSAKeySize(key *rsa.PrivateKey) int {
	return key.N.BitLen()
}

func determineRSATypeFromSize(keySize int) (appmodels.KeyType, error) {
	switch keySize {
	case 1024:
		return appmodels.RSA_1024, nil
	case 2048:
		return appmodels.RSA_2048, nil
	case 3072:
		return appmodels.RSA_3072, nil
	case 4096:
		return appmodels.RSA_4096, nil
	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("determineRSAKeySize: RSA key size not supported: %d", keySize)
	}
}

func determineRSATypeFromKey(privateKey any) (appmodels.KeyType, error) {
	switch key := privateKey.(type) {
	case *rsa.PrivateKey:
		keySize := readRSAKeySize(key)
		keyType, err := determineRSATypeFromSize(keySize)
		if err != nil {
			return appmodels.NIL_KEY_TYPE, fmt.Errorf("determineRSAKeySize: failed: %w", err)
		}
		return keyType, nil
	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("determineRSAKeySize: not an RSA key")
	}
}

func parseRSAPrivateKey(bytes []byte) (any, appmodels.KeyType, error) {
	privateKey, err := x509.ParsePKCS1PrivateKey(bytes)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("parseRSAPrivateKey: failed to parse RSA private key: %w", err)
	}
	keyType, err := determineRSATypeFromKey(privateKey)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("parseRSAPrivateKey: failed to parse RSA type: %w", err)
	}
	return privateKey, keyType, nil
}

func parseECPrivateKey(bytes []byte) (any, appmodels.KeyType, error) {
	privateKey, err := x509.ParseECPrivateKey(bytes)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("failed to parse EC private key: %w", err)
	}
	// Determine the curve
	keyType, err := determineECDSACurve(privateKey.Curve)
	if err != nil {
		return nil, appmodels.NIL_KEY_TYPE, fmt.Errorf("could not detect key type: %w", err)
	}
	return privateKey, keyType, nil
}

// determineECDSACurve determines the ECDSA key type from curve
func determineECDSACurve(curve elliptic.Curve) (appmodels.KeyType, error) {
	switch curve {
	case elliptic.P224():
		return appmodels.ECDSA_P224, nil
	case elliptic.P256():
		return appmodels.ECDSA_P256, nil
	case elliptic.P384():
		return appmodels.ECDSA_P384, nil
	case elliptic.P521():
		return appmodels.ECDSA_P521, nil
	default:
		return appmodels.NIL_KEY_TYPE, fmt.Errorf("unsupported EC curve")
	}
}

// privateKeyToPEM converts a private key to PEM format.
// This function needs to handle different key types accordingly.
func privateKeyToPEM(key appmodels.IPrivateKey) ([]byte, error) {
	var pemBlock *pem.Block

	switch typedKey := key.GetPrivateKey().(type) {
	case *rsa.PrivateKey:
		pemBlock = &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(typedKey)}
	case *ecdsa.PrivateKey:
		ecBytes, err := x509.MarshalECPrivateKey(typedKey)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal ECDSA private key: %w", err)
		}
		pemBlock = &pem.Block{Type: "EC PRIVATE KEY", Bytes: ecBytes}
	case ed25519.PrivateKey:
		// Ed25519 keys are handled differently as they do not have an explicit marshal function in the x509 package.
		return nil, fmt.Errorf("Ed25519 private key saving not supported")
	default:
		return nil, fmt.Errorf("unsupported private key type")
	}

	return pem.EncodeToMemory(pemBlock), nil
}
