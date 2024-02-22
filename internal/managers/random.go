// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package managers

import (
	"crypto/rand"
	"math/big"
)

// IRandomManager describes operations to create random values
type IRandomManager interface {
	CreateBigInt(max *big.Int) (*big.Int, error)
}

type RandomManager struct {
}

func NewRandomManager() *RandomManager {
	return &RandomManager{}
}

func (r *RandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}
