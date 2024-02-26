// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func NewCollection(filePath string) *appmodels.Collection {
	return appmodels.NewCollection(
		NewOrganizationRepository(filePath),
		NewCertificateRepository(filePath),
		NewPrivateKeyRepository(filePath),
	)
}
