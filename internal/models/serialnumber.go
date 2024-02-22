// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package models

import (
	"math/big"
)

type SerialNumber *big.Int

func NewSerialNumber(randomManager IRandomManager) (SerialNumber, error) {
	serialNumber, err := randomManager.CreateBigInt(new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}
	return serialNumber, nil
}
