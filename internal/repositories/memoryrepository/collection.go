// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"github.com/hyperifyio/gocertcenter/internal/models"
)

func NewCollection() *models.Collection {
	return models.NewCollection(
		NewOrganizationRepository(),
		NewCertificateRepository(),
		NewPrivateKeyRepository(),
	)
}
