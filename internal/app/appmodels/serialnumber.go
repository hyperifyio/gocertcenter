// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"fmt"
	"math/big"
)

type Int64SerialNumber struct {
	value *big.Int
}

func (s *Int64SerialNumber) String() string {
	return s.value.String()
}

func (s *Int64SerialNumber) Value() *big.Int {
	return s.value
}

func (s *Int64SerialNumber) Cmp(s2 SerialNumber) int {
	return s.value.Cmp(s2.Value())
}

func (s *Int64SerialNumber) Sign() int {
	return s.value.Sign()
}

func NewSerialNumber(value *big.Int) SerialNumber {
	return &Int64SerialNumber{value: value}
}

func ParseSerialNumber(value string, base int) (SerialNumber, error) {

	if value == "" {
		return nil, fmt.Errorf("[ParseSerialNumber]: no value provided")
	}

	if base <= 1 {
		return nil, fmt.Errorf("[ParseSerialNumber]: invalid base: %d", base)
	}

	newSerialNumber := new(Int64SerialNumber)
	newSerialNumber.value = new(big.Int)
	_, success := newSerialNumber.value.SetString(value, base)
	if success {
		return newSerialNumber, nil
	} else {
		return nil, fmt.Errorf("[ParseSerialNumber]: failed to parse: %s", value)
	}
}

var _ SerialNumber = (*Int64SerialNumber)(nil)
