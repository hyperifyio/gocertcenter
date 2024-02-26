// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers

import (
	"fmt"
	"os"
	"path/filepath"
)

// CertificateManager implements operations to manage x509 certificates by
// implementing models.ICertificateManager. This is intended to wrap low level
// external library operations for easier testing by using mocks. Any higher
// level operations shouldn't be implemented inside it.
type FileManager struct {
}

// ReadBytes reads bytes from a file
//   - fileName string: The file where to read
//
// Returns the bytes read or error
func (f FileManager) ReadBytes(fileName string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, fmt.Errorf("failed to read a file: %s: %w", fileName, err)
	}
	return data, nil
}

// SaveBytes saves bytes to a file. It will first create any parent directories.
//   - fileName string: The file where to save
//   - data []byte: The data to save
//   - filePerms os.FileMode: Permissions for file
//   - dirPerms os.FileMode: Permissions for directories
//
// Returns nil or error
func (f FileManager) SaveBytes(fileName string, data []byte, filePerms, dirPerms os.FileMode) error {

	fileName = filepath.Clean(fileName)

	// Ensure the directory exists
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, dirPerms); err != nil {
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

	// FIXME: Set file permissions

	// Write the data to the temporary file
	if _, err := tmpFile.Write(data); err != nil {
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

func NewFileManager() FileManager {
	return FileManager{}
}

var _ IFileManager = (*FileManager)(nil)
