// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"math/big"
	"testing"
)

func TestRandomManager_CreateBigInt(t *testing.T) {
	manager := NewRandomManager()
	max := big.NewInt(100) // Define a maximum value for the random number

	for i := 0; i < 10; i++ { // Perform multiple iterations to check the randomness
		result, err := manager.CreateBigInt(max)
		if err != nil {
			t.Fatalf("CreateBigInt returned an unexpected error: %v", err)
		}

		if result.Cmp(big.NewInt(0)) == -1 || result.Cmp(max) >= 0 {
			t.Errorf("Expected result to be >= 0 and < %v, got %v", max, result)
		}
	}
}

func TestRandomManager_Randomness(t *testing.T) {
	manager := NewRandomManager()
	max := big.NewInt(1 << 62) // Use a large max to reduce the chance of collisions
	iterations := 100
	results := make(map[string]struct{}, iterations)

	for i := 0; i < iterations; i++ {
		result, err := manager.CreateBigInt(max)
		if err != nil {
			t.Fatalf("CreateBigInt returned an unexpected error: %v", err)
		}
		// Use the string representation of the result as the map key
		resultStr := result.String()
		if _, exists := results[resultStr]; exists {
			t.Fatalf("Duplicate result found, indicating lack of randomness: %v", result)
		}
		results[resultStr] = struct{}{}
	}
}
