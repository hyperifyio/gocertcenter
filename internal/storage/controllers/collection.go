// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package controllers

// ControllerCollection contains all the controller instances
type ControllerCollection struct {
	Certificate ICertificateService
	PrivateKey  IPrivateKeyService
}

// NewControllerCollection returns a new ControllerCollection instance
func NewControllerCollection(
	certificate ICertificateService,
	privateKey IPrivateKeyService,
) *ControllerCollection {
	return &ControllerCollection{
		Certificate: certificate,
		PrivateKey:  privateKey,
	}
}
