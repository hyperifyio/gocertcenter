// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package hashutils_test

import (
	"errors"
	"github.com/hyperifyio/gocertcenter/internal/mocks"
	"github.com/stretchr/testify/mock"
	"hash/fnv"
	"testing"

	"github.com/hyperifyio/gocertcenter/internal/hashutils"
)

func TestToUint64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0xcbf29ce484222325, // Expected FNV-1a hash of an empty string
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: 8618312879776256743, // Expected FNV-1a hash of "hello world"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashutils.ToUint64(tt.input, fnv.New64a())
			if err != nil {
				t.Errorf("ToUint64(%s) = got error %v", tt.input, err)
			}
			if got != tt.expected {
				t.Errorf("ToUint64(%s) = got %v, expected %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestToUint64_WriteError(t *testing.T) {

	// Create an instance of the mock Hash64
	mockHash := new(mocks.MockHash64)

	// Simulate an error on Write
	mockHash.On("Write", mock.Anything).Return(0, errors.New("write error"))

	// Since the function logs and exits on error, we can't assert the function's behavior directly
	// Instead, we assert that the mock was called with the expected parameters
	// This is more of demonstrating the mock setup rather than a practical test for log.Fatalf
	testString := "test"

	// Normally, you would call hashutils.ToUint64 here, but because it calls log.Fatalf, doing so would exit the test runner
	_, err := hashutils.ToUint64(testString, mockHash)
	if err == nil {
		t.Errorf("ToUint64 = got unexpected success")
	}

	// Assert that Write was called with the correct parameters
	mockHash.AssertCalled(t, "Write", []byte(testString))

	// Additional cleanup or assertions as needed
}
