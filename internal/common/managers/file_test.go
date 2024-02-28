// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package managers_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func TestFile_Close(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile_")
	require.NoError(t, err)

	file := managers.NewFile(tmpFile)
	err = file.Close()
	assert.NoError(t, err, "File close should not error")

	// Try reopening to ensure it was closed
	_, err = os.Open(tmpFile.Name())
	assert.NoError(t, err, "Reopening the closed file should succeed")
}

func TestFile_Name(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile_")
	require.NoError(t, err)
	defer tmpFile.Close()

	file := managers.NewFile(tmpFile)
	name := file.Name()
	assert.Equal(t, tmpFile.Name(), name, "File name should match")
}

func TestFile_Write(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile_")
	require.NoError(t, err)
	defer tmpFile.Close()

	file := managers.NewFile(tmpFile)
	data := []byte("Hello, world!")
	n, err := file.Write(data)
	assert.NoError(t, err, "File write should not error")
	assert.Equal(t, len(data), n, "Write should return correct byte count")

	// Verify the data was written correctly
	content, err := ioutil.ReadFile(tmpFile.Name())
	require.NoError(t, err)
	assert.Equal(t, data, content, "File content should match written data")
}
