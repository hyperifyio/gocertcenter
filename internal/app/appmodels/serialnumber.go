// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"fmt"
	"math/big"
)

func NewSerialNumber(value int64) *big.Int {
	return big.NewInt(value)
}

func ParseSerialNumber(value string, base int) (*big.Int, error) {

	if value == "" {
		return nil, fmt.Errorf("[ParseSerialNumber]: no value provided")
	}

	if base <= 1 {
		return nil, fmt.Errorf("[ParseSerialNumber]: invalid base: %d", base)
	}

	newSerialNumber := new(big.Int)
	_, success := newSerialNumber.SetString(value, base)
	if success {
		return newSerialNumber, nil
	} else {
		return nil, fmt.Errorf("[ParseSerialNumber]: failed to parse: %s", value)
	}
}
