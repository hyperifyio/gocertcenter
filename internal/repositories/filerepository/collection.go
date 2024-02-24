// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

func NewCollection(filePath string) *models.Collection {
	return models.NewCollection(
		NewOrganizationRepository(filePath),
		NewCertificateRepository(filePath),
		NewPrivateKeyRepository(filePath),
	)
}
