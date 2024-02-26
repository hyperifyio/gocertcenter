// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package filerepository

import (
	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
	"github.com/hyperifyio/gocertcenter/internal/common/managers"
)

func NewCollection(
	certManager managers.ICertificateManager,
	fileManager managers.IFileManager,
	filePath string,
) *appmodels.Collection {
	return appmodels.NewCollection(
		NewOrganizationRepository(certManager, fileManager, filePath),
		NewCertificateRepository(certManager, fileManager, filePath),
		NewPrivateKeyRepository(certManager, fileManager, filePath),
	)
}
