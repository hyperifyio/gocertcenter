// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package fsutils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

// SaveBytes writes the given byte slice to a specified file, ensuring that any
// necessary parent directories are created before writing. This function
// abstracts filesystem operations through the managers.IFileManager interface,
// allowing for easier testing and integration with different filesystems.
//
// Parameters:
//   - fs: An implementation of the managers.IFileManager interface, which
//     encapsulates filesystem operations.
//   - fileName: The path to the file where the data will be saved. If the file
//     does not exist, it will be created along with any necessary parent
//     directories.
//   - data: The byte slice containing the data to be written to the file.
//   - filePerms: The file permissions to set for the newly created file. This
//     parameter is ignored if the file already exists.
//   - dirPerms: The permissions to apply when creating any new directories.
//     This parameter is only used if new directories need to be created.
//
// Returns:
//   - nil if the operation was successful.
//   - An error if there was a problem creating the directories, creating the file, setting permissions, or writing the data.
func SaveBytes(fs managers.IFileManager, fileName string, data []byte, filePerms, dirPerms os.FileMode) error {

	fileName = filepath.Clean(fileName)

	// Ensure the directory exists
	dir := filepath.Dir(fileName)
	if err := fs.MkdirAll(dir, dirPerms); err != nil {
		return fmt.Errorf("failed to create a directory %s: %w", dir, err)
	}

	// Flag to check if rename was successful
	renameSuccessful := false

	// Create a temporary file within the final file's directory
	tmpFile, err := fs.CreateTemp(dir, "*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create a temporary file: %w", err)
	}

	// Cleanup function to remove the temporary file if rename fails
	defer func() {
		_ = tmpFile.Close()

		if !renameSuccessful {
			_ = fs.Remove(tmpFile.Name()) // Ignore error here as it's cleanup
		}
	}()

	// Set the file permissions as desired
	if err := fs.Chmod(tmpFile.Name(), filePerms); err != nil {
		return fmt.Errorf("failed to set file permissions for %s: %w", tmpFile.Name(), err)
	}

	// Write the data to the temporary file
	if _, err := tmpFile.Write(data); err != nil {
		return fmt.Errorf("failed to write to the temporary file: %w", err)
	}

	// Close the file to ensure all writes are flushed to disk
	if err := tmpFile.Close(); err != nil {
		return fmt.Errorf("failed to close the temporary file: %w", err)
	}

	// Move the temporary file to the final location
	if err := fs.Rename(tmpFile.Name(), fileName); err != nil {
		return fmt.Errorf("failed to move the temporary file to %s: %w", fileName, err)
	}

	renameSuccessful = true
	return nil

}
