// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package modelutils

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
	"math/big"
)

func GenerateSerialNumber(randomManager models.IRandomManager) (models.ISerialNumber, error) {
	value, err := randomManager.CreateBigInt(new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return models.NewSerialNumber(nil), err
	}
	return models.NewSerialNumber(value), nil
}
