// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package filerepository_test

import (
	"os"
	"testing"
)

// setupTempDir is utility function for test in filerepository_test package
func setupTempDir(t *testing.T) (string, func()) {
	t.Helper()
	// Create a temporary directory
	tempDir, err := os.MkdirTemp("", "orgRepoTest")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	// Return the path and a cleanup function
	cleanup := func() {
		os.RemoveAll(tempDir) // clean up the temp directory
	}

	return tempDir, cleanup
}
