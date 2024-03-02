// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func GenerateSerialNumber(randomManager managers.RandomManager) (*big.Int, error) {
	value, err := randomManager.CreateBigInt(new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return nil, err
	}
	return value, nil
}
