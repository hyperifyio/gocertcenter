// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package memoryrepository

import (
	"strings"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func getCertificateLocator(organization string, certificates []appmodels.ISerialNumber) string {
	parts := []string{organization}
	for _, certificate := range certificates {
		parts = append(parts, certificate.String())
	}
	return strings.Join(parts, "/")
}
