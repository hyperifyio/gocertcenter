// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package modelcontrollers

import "github.com/hyperifyio/gocertcenter/internal/models"

// ControllerCollection contains all the controller instances
type ControllerCollection struct {
	Certificate models.ICertificateService
	PrivateKey  models.IPrivateKeyService
}

// NewControllerCollection returns a new ControllerCollection instance
func NewControllerCollection(
	certificate models.ICertificateService,
	privateKey models.IPrivateKeyService,
) *ControllerCollection {
	return &ControllerCollection{
		Certificate: certificate,
		PrivateKey:  privateKey,
	}
}
