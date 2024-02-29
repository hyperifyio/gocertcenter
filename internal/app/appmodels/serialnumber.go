// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"fmt"
	"math/big"
)

type SerialNumber struct {
	value *big.Int
}

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

func ParseSerialNumber(value string, base int) (ISerialNumber, error) {

	if value == "" {
		return nil, fmt.Errorf("[ParseSerialNumber]: no value provided")
	}

	if base <= 1 {
		return nil, fmt.Errorf("[ParseSerialNumber]: invalid base: %d", base)
	}

	newSerialNumber := new(SerialNumber)
	newSerialNumber.value = new(big.Int)
	_, success := newSerialNumber.value.SetString(value, base)
	if success {
		return newSerialNumber, nil
	} else {
		return nil, fmt.Errorf("[ParseSerialNumber]: failed to parse: %s", value)
	}
}

var _ ISerialNumber = (*SerialNumber)(nil)
