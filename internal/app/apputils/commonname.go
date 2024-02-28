// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"errors"
	"regexp"
	"strings"
)

// ValidateRootCertificateCommonName checks if the provided common name adheres to specific rules:
// - Should not be full numbers
// - Should only have a-z, A-Z, _, -, ., space, 0-9
// - No repeating spaces
// - Must not be empty
// - No prefix or suffix spaces
func ValidateRootCertificateCommonName(commonName string) error {

	if commonName == "" {
		return errors.New("cannot be empty")
	}

	if strings.HasPrefix(commonName, " ") {
		return errors.New("cannot start with a space")
	}

	if strings.HasSuffix(commonName, " ") {
		return errors.New("cannot end with a space")
	}

	// Check for repeating spaces
	if strings.Contains(commonName, "  ") {
		return errors.New("should not have repeating spaces")
	}

	// Check if commonName is full numbers
	ullNumbersRegex := regexp.MustCompile(`^\d+$`)
	if ullNumbersRegex.MatchString(commonName) {
		return errors.New("should not be full numbers")
	}

	// Regular expression to check valid characters and no repeating spaces
	validCharsRegex := regexp.MustCompile(`^[a-zA-Z0-9_\-. ]+$`)
	if !validCharsRegex.MatchString(commonName) {
		return errors.New("contains invalid characters")
	}

	return nil
}

// ValidateClientCertificateCommonName checks if the provided common name adheres to specific rules:
// - Should only have characters: a-z, A-Z, _, -, ., space, 0-9, @
// - No repeating spaces
// - No leading spaces
// - No starting spaces
func ValidateClientCertificateCommonName(commonName string) error {

	if commonName == "" {
		return errors.New("cannot be empty")
	}

	if strings.HasPrefix(commonName, " ") {
		return errors.New("cannot start with a space")
	}

	if strings.HasSuffix(commonName, " ") {
		return errors.New("cannot end with a space")
	}

	// Check for repeating spaces
	if strings.Contains(commonName, "  ") {
		return errors.New("should not have repeating spaces")
	}

	// Regular expression to check for valid characters
	validCharsRegex := regexp.MustCompile(`^[a-zA-Z0-9_\-\. @]+$`)
	if !validCharsRegex.MatchString(commonName) {
		return errors.New("contains invalid characters")
	}

	return nil
}
