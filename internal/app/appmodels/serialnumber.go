// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package appmodels

import (
	"math/big"
)

func NewSerialNumber(value int64) *big.Int {
	return big.NewInt(value)
}
