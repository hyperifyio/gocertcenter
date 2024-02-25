// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package memoryrepository

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
	"strings"
)

func getCertificateLocator(organization string, certificates []models.ISerialNumber) string {
	parts := []string{organization}
	for _, certificate := range certificates {
		parts = append(parts, certificate.String())
	}
	return strings.Join(parts, "/")
}
