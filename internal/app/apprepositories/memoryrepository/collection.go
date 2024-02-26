// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package memoryrepository

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
)

func NewCollection() *appmodels.Collection {
	return appmodels.NewCollection(
		NewOrganizationRepository(),
		NewCertificateRepository(),
		NewPrivateKeyRepository(),
	)
}
