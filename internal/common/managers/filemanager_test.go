// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package managers_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestFileManager_ReadFile(t *testing.T) {
	// Setup: Create a temporary file with content
	content := []byte("Hello, FileManager!")
	tmpFile, err := ioutil.TempFile("", "example")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up

	_, err = tmpFile.Write(content)
	require.NoError(t, err)
	require.NoError(t, tmpFile.Close())

	fm := managers.NewFileManager()

	// Test
	readContent, err := fm.ReadFile(tmpFile.Name())
	require.NoError(t, err)
	assert.Equal(t, content, readContent)
}

func TestFileManager_MkdirAll_And_Remove(t *testing.T) {
	// Setup: Create a temporary directory path
	tmpDirPath := filepath.Join(os.TempDir(), "filemanagertest")
	defer os.RemoveAll(tmpDirPath) // Clean up

	fm := managers.NewFileManager()

	// Test MkdirAll
	err := fm.MkdirAll(tmpDirPath, 0755)
	require.NoError(t, err)

	// Verify the directory was created
	_, err = os.Stat(tmpDirPath)
	require.NoError(t, err)

	// Test Remove
	err = fm.Remove(tmpDirPath)
	require.NoError(t, err)

	// Verify the directory was removed
	_, err = os.Stat(tmpDirPath)
	require.True(t, os.IsNotExist(err))
}

func TestFileManager_CreateTemp_And_Chmod(t *testing.T) {
	fm := managers.NewFileManager()

	// Test CreateTemp
	tmpFile, err := fm.CreateTemp("", "tempfile-*")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name()) // Clean up

	// Test Chmod
	err = fm.Chmod(tmpFile.Name(), 0600)
	require.NoError(t, err)

	// Verify the permissions
	info, err := os.Stat(tmpFile.Name())
	require.NoError(t, err)
	assert.Equal(t, "-rw-------", info.Mode().String())
}

func TestFileManager_Rename(t *testing.T) {
	// Setup: Create a temporary file
	tmpDir, err := ioutil.TempDir("", "test")
	require.NoError(t, err)
	defer os.RemoveAll(tmpDir) // Clean up

	tmpFile, err := ioutil.TempFile(tmpDir, "original-")
	require.NoError(t, err)
	originalPath := tmpFile.Name()
	newPath := filepath.Join(tmpDir, "renamed-file")

	fm := managers.NewFileManager()

	// Test Rename
	err = fm.Rename(originalPath, newPath)
	require.NoError(t, err)

	// Verify the file was renamed
	_, err = os.Stat(newPath)
	require.NoError(t, err)
}
