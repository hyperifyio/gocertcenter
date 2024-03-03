// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository

import (
	"fmt"
	"math/big"
)

func getCertificateLocator(organization string, certificate *big.Int) string {
	return fmt.Sprintf("%s/%s", organization, certificate.String())
}
