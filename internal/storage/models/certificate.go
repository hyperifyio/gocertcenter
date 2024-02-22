// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package models

// Certificate model
type Certificate struct {
	SerialNumber string
}

// NewCertificate creates a certificate model
func NewCertificate(SerialNumber string) *Certificate {
	return &Certificate{SerialNumber}
}
