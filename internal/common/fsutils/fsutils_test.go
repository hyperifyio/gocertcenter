// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package fsutils_test

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/hyperifyio/gocertcenter/internal/common/commonmocks"
	"github.com/hyperifyio/gocertcenter/internal/common/fsutils"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestSaveBytes(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	mockFile := new(commonmocks.MockFile)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)

	// Setup mock expectations for OSFileManager
	mockFS.On("MkdirAll", mock.AnythingOfType("string"), dirPerms).Return(nil)
	mockFS.On("CreateTemp", dir, "*.tmp").Return(mockFile, nil)
	mockFS.On("Chmod", mock.AnythingOfType("string"), filePerms).Return(nil)
	mockFS.On("Rename", mock.AnythingOfType("string"), fileName).Return(nil)

	// Setup mock expectations for OSFile
	mockTempFileName := "tempfile.tmp"
	mockFile.On("Write", data).Return(len(data), nil)
	mockFile.On("Close").Return(nil)
	mockFile.On("Name").Return(mockTempFileName)

	// Call SaveBytes with the mock file system
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)
	assert.NoError(t, err)

	// Verify that all expectations were met
	mockFS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestSaveBytes_FailToCreateDirectory(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)

	// Mock MkdirAll to return an error
	mockFS.On("MkdirAll", dir, dirPerms).Return(errors.New("failed to create directory"))

	// Call SaveBytes expecting it to handle the directory creation error
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)

	// Assert that an error was returned
	assert.Error(t, err, "Expected an error due to failed directory creation")

	// Verify the error message
	assert.Contains(t, err.Error(), "failed to create a directory", "Error message should reflect the failed directory creation")

	// Verify that all expectations were met
	mockFS.AssertExpectations(t)
}

func TestSaveBytes_FailToCreateTempFile(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)

	// Setup mock expectations for OSFileManager
	// Simulate success for directory creation
	mockFS.On("MkdirAll", mock.AnythingOfType("string"), dirPerms).Return(nil)

	// Simulate failure for temporary file creation
	mockFS.On("CreateTemp", dir, "*.tmp").Return(managers.NewFile(nil), errors.New("temp file creation failed"))

	// Call SaveBytes with the mock file system expecting a failure in temporary file creation
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)

	// Verify that an error was returned due to the failure in creating a temporary file
	assert.Error(t, err, "Expected an error due to failed temporary file creation")
	assert.Contains(t, err.Error(), "failed to create a temporary file", "Error message should reflect the failed temporary file creation")

	// Verify that all expectations were met
	mockFS.AssertExpectations(t)
}

func TestSaveBytes_FailToSetFilePermissions(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	mockFile := new(commonmocks.MockFile)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)
	mockTempFileName := "tempfile.tmp"

	// Mock expectations for OSFileManager to simulate directory creation and temporary file creation
	mockFile.On("Name").Return(mockTempFileName)
	mockFile.On("Close").Return(nil)

	mockFS.On("MkdirAll", mock.AnythingOfType("string"), dirPerms).Return(nil)
	mockFS.On("CreateTemp", dir, "*.tmp").Return(mockFile, nil)
	mockFS.On("Remove", mockTempFileName).Return(nil)

	// Simulate failure for setting file permissions
	mockFS.On("Chmod", mockTempFileName, filePerms).Return(errors.New("chmod failed"))

	// Call SaveBytes expecting it to handle the file permission setting error
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)

	// Verify that an error was returned due to the failure in setting file permissions
	assert.Error(t, err, "Expected an error due to failed file permission setting")
	assert.Contains(t, err.Error(), "failed to set file permissions", "Error message should reflect the failed file permission setting")

	// Verify that all expectations were met
	mockFS.AssertExpectations(t)
	mockFile.AssertExpectations(t) // Ensure expectations for the file operations were also met
}

func TestSaveBytes_FailToWriteToTempFile(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	mockFile := new(commonmocks.MockFile)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)
	mockTempFileName := "tempfile.tmp"

	// Setup mock expectations
	mockFile.On("Close").Return(nil)
	mockFS.On("MkdirAll", mock.AnythingOfType("string"), dirPerms).Return(nil)
	mockFS.On("CreateTemp", dir, "*.tmp").Return(mockFile, nil)
	mockFS.On("Remove", mockTempFileName).Return(nil)
	mockFS.On("Chmod", mockTempFileName, filePerms).Return(nil)

	// Simulate failure for setting file permissions
	mockFile.On("Write", data).Return(0, errors.New("write failed"))
	mockFile.On("Name").Return("tempfile.tmp")

	// Call SaveBytes expecting a failure in writing to the temporary file
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)

	// Verify the error
	assert.Error(t, err, "Expected an error due to failed writing to temporary file")
	assert.Contains(t, err.Error(), "failed to write to the temporary file", "Error message should reflect the failed write operation")

	mockFS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestSaveBytes_FailToCloseTempFile(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	mockFile := new(commonmocks.MockFile)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)
	mockTempFileName := "tempfile.tmp"

	// Setup mock expectations
	mockFS.On("Remove", mockTempFileName).Return(nil)
	mockFS.On("Chmod", mockTempFileName, filePerms).Return(nil)
	mockFS.On("MkdirAll", mock.AnythingOfType("string"), dirPerms).Return(nil)
	mockFS.On("CreateTemp", dir, "*.tmp").Return(mockFile, nil)
	mockFile.On("Write", data).Return(len(data), nil)
	mockFile.On("Close").Return(errors.New("close failed"))
	mockFile.On("Name").Return("tempfile.tmp")
	mockFS.On("Remove", mock.AnythingOfType("string")).Return(nil) // Mock cleanup

	// Call SaveBytes expecting a failure when closing the temporary file
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)

	// Verify the error
	assert.Error(t, err, "Expected an error due to failed closing of temporary file")
	assert.Contains(t, err.Error(), "failed to close the temporary file", "Error message should reflect the failed close operation")

	mockFS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestSaveBytes_FailToMoveTempFile(t *testing.T) {
	mockFS := new(commonmocks.MockFileManager)
	mockFile := new(commonmocks.MockFile)
	dir, fileName, data := "/test/dir", "/test/dir/file.txt", []byte("data")
	filePerms, dirPerms := os.FileMode(0644), os.FileMode(0755)
	mockTempFileName := "tempfile.tmp"

	// Setup mock expectations
	mockFS.On("MkdirAll", mock.AnythingOfType("string"), dirPerms).Return(nil)
	mockFS.On("CreateTemp", dir, "*.tmp").Return(mockFile, nil)
	mockFile.On("Write", data).Return(len(data), nil)
	mockFile.On("Close").Return(nil)
	mockFile.On("Name").Return(mockTempFileName)
	mockFS.On("Chmod", mockTempFileName, filePerms).Return(nil)

	// Simulate failure for moving the temporary file to its final location
	mockFS.On("Rename", mockTempFileName, fileName).Return(errors.New("rename failed"))
	mockFS.On("Remove", mockTempFileName).Return(nil) // Expect cleanup action

	// Call SaveBytes expecting a failure when moving the temporary file
	err := fsutils.SaveBytes(mockFS, fileName, data, filePerms, dirPerms)

	// Verify the error
	assert.Error(t, err, "Expected an error due to failed moving of temporary file to final location")
	assert.Contains(t, err.Error(), "failed to move the temporary file", "Error message should reflect the failed move operation")

	// Ensure all mock expectations were met
	mockFS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}
