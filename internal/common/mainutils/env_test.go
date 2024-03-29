// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mainutils_test

import (
	"os"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/common/mainutils"
)

func TestGetEnvOrDefault_VariableSet(t *testing.T) {
	const testKey = "TEST_ENV_VAR"
	const testValue = "expected value"

	// Set up the environment variable for the test
	err := os.Setenv(testKey, testValue)
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}

	// Ensure the environment variable is cleaned up after the test
	defer os.Unsetenv(testKey)

	// Test EnvOrDefault
	got := mainutils.EnvOrDefault(testKey, "default value")
	if got != testValue {
		t.Errorf("EnvOrDefault(%q, %q) = %q, want %q", testKey, "default value", got, testValue)
	}
}

func TestGetEnvOrDefault_VariableNotSet(t *testing.T) {
	const testKey = "NONEXISTENT_TEST_ENV_VAR"
	const defaultValue = "default value"

	// Ensure the environment variable is not set
	err := os.Unsetenv(testKey)
	if err != nil {
		t.Fatalf("Failed to unset environment variable: %v", err)
	}

	// Test EnvOrDefault
	got := mainutils.EnvOrDefault(testKey, defaultValue)
	if got != defaultValue {
		t.Errorf("EnvOrDefault(%q, %q) = %q, want %q", testKey, defaultValue, got, defaultValue)
	}
}
