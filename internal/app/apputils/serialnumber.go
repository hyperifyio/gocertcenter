// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"math/big"

	"github.com/hyperifyio/gocertcenter/internal/common/managers"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func GenerateSerialNumber(randomManager managers.RandomManager) (appmodels.SerialNumber, error) {
	value, err := randomManager.CreateBigInt(new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return appmodels.NewSerialNumber(nil), err
	}
	return appmodels.NewSerialNumber(value), nil
}
