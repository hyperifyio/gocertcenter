// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package managers

import (
	"crypto/rand"
	"math/big"
)

type RandomManager struct {
}

func NewRandomManager() *RandomManager {
	return &RandomManager{}
}

func (r *RandomManager) CreateBigInt(max *big.Int) (*big.Int, error) {
	return rand.Int(rand.Reader, max)
}
