// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"fmt"
	"math/big"
)

func ParseBigInt(value string, base int) (*big.Int, error) {

	if value == "" {
		return nil, fmt.Errorf("[ParseBigInt]: no value provided")
	}

	if base <= 1 {
		return nil, fmt.Errorf("[ParseBigInt]: invalid base: %d", base)
	}

	newSerialNumber := new(big.Int)
	_, success := newSerialNumber.SetString(value, base)
	if success {
		return newSerialNumber, nil
	} else {
		return nil, fmt.Errorf("[ParseBigInt]: failed to parse: %s", value)
	}
}
