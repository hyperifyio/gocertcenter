// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository

import (
	"math/big"
	"strings"
)

func getCertificateLocator(organization string, certificates []*big.Int) string {
	parts := []string{organization}
	for _, certificate := range certificates {
		parts = append(parts, certificate.String())
	}
	return strings.Join(parts, "/")
}
