// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"crypto/rand"
	"math/big"
)

// RandomManager implements models.IRandomManager
type RandomManager struct {
}

// CreateBigInt wraps up a call to rand.Int with rand.Reader
func (r *RandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}

func NewRandomManager() *RandomManager {
	return &RandomManager{}
}

var _ IRandomManager = (*RandomManager)(nil)
