// Copyright (c) 2024. Hangover Games <info@hangover.games>. All rights reserved.

package models

// Key model
type Key struct {
	SerialNumber string
}

// NewKey creates a key model
func NewKey(SerialNumber string) *Key {
	return &Key{SerialNumber}
}
