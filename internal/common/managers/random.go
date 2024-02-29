// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"crypto/rand"
	"math/big"
)

// RandRandomManager implements RandomManager
type RandRandomManager struct {
}

// CreateBigInt wraps up a call to rand.Int with rand.Reader
func (r *RandRandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}

func NewRandomManager() *RandRandomManager {
	return &RandRandomManager{}
}

var _ RandomManager = (*RandRandomManager)(nil)
