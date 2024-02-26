// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package managers

import (
	"crypto/rand"
	"math/big"
)

// RandomManager implements models.IRandomManager
type RandomManager struct {
}

var _ IRandomManager = (*RandomManager)(nil)

func NewRandomManager() *RandomManager {
	return &RandomManager{}
}

func (r *RandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}
