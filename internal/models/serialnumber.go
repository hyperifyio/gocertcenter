// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"math/big"
)

type SerialNumber struct {
	value *big.Int
}

var _ ISerialNumber = (*SerialNumber)(nil)

func (s *SerialNumber) String() string {
	return s.value.String()
}

func (s *SerialNumber) Value() *big.Int {
	return s.value
}

func (s *SerialNumber) Cmp(s2 ISerialNumber) int {
	return s.value.Cmp(s2.Value())
}

func (s *SerialNumber) Sign() int {
	return s.value.Sign()
}

func NewSerialNumber(value *big.Int) ISerialNumber {
	return &SerialNumber{value: value}
}
