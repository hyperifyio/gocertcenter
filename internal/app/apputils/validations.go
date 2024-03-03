// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hyperifyio/gocertcenter/internal/app/appmodels"
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

// ValidateServerCertificateCommonName checks if the provided common name adheres to specific rules for server certificates.
//   - Check if common name starts with a wildcard
//   - Check for invalid start or end characters
//   - Check for repeating periods
//   - Check for valid characters [a-z0-9\-\.]
//   - Check last TLD length
func ValidateServerCertificateCommonName(commonName string) error {

	if strings.HasPrefix(commonName, "*.") {
		commonName = commonName[2:] // Remove the wildcard for further validation.
	}

	if commonName == "" {
		return errors.New("cannot be empty")
	}

	if strings.HasPrefix(commonName, "-") || strings.HasPrefix(commonName, ".") || strings.HasSuffix(commonName, "-") || strings.HasSuffix(commonName, ".") {
		return errors.New("cannot start or end with '-' or '.'")
	}

	if strings.Contains(commonName, "..") {
		return errors.New("cannot have repeating '.'")
	}

	validCharsRegex := regexp.MustCompile(`^[a-z0-9\-.]+$`)
	if !validCharsRegex.MatchString(commonName) {
		return errors.New("contains invalid characters")
	}

	// Check last TLD length.
	parts := strings.Split(commonName, ".")
	if len(parts[len(parts)-1]) < 2 {
		return errors.New("last TLD must be at least two characters")
	}

	return nil
}

// ValidateDNSNames validates an array of DNS names using the rules for server certificates.
func ValidateDNSNames(dnsNames []string) error {
	for _, dnsName := range dnsNames {
		if err := ValidateServerCertificateCommonName(dnsName); err != nil {
			return err
		}
	}
	return nil
}

// ValidateOrganizationName checks if the provided name adheres to specific rules.
func ValidateOrganizationName(name string) error {
	if len(name) == 0 || len(strings.TrimSpace(name)) < 2 {
		return errors.New("must not be empty and must be at least two characters long")
	}
	if strings.TrimSpace(name) != name {
		return errors.New("must not have leading or trailing spaces")
	}
	if strings.Contains(name, "  ") {
		return errors.New("should not have repeating spaces")
	}
	if _, err := strconv.Atoi(name); err == nil {
		return errors.New("should not be full numbers")
	}
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_\-. ]+$`, name)
	if !matched {
		return errors.New("contains invalid characters")
	}
	return nil
}

// ValidateOrganizationNames validates an array of DNS names using the rules for server certificates.
func ValidateOrganizationNames(list []string) error {
	for _, item := range list {
		if err := ValidateOrganizationName(item); err != nil {
			return err
		}
	}
	return nil
}

// ValidateOrganizationSlug checks if the provided slug adheres to specific rules.
func ValidateOrganizationSlug(id string) error {
	if len(id) < 2 {
		return errors.New("must be at least two characters long")
	}
	if _, err := strconv.Atoi(id); err == nil {
		return errors.New("should not be full numbers")
	}
	matched, _ := regexp.MatchString(`[A-Z]`, id)
	if matched {
		return errors.New("contains uppercase characters")
	}
	matched, _ = regexp.MatchString(`^[a-z0-9\-.]+$`, id)
	if !matched || id != strings.Trim(id, " -.") {
		return errors.New("contains invalid characters, or has leading/trailing spaces, '-', or '.'")
	}
	return nil
}

func ValidateOrganizationModel(model appmodels.Organization) error {
	id := model.Slug()
	if err := ValidateOrganizationSlug(id); err != nil {
		return fmt.Errorf("slug: '%v': %v", id, err)
	}
	name := model.Name()
	if err := ValidateOrganizationName(name); err != nil {
		return fmt.Errorf("name: '%v': %v", name, err)
	}
	names := model.Names()
	if err := ValidateOrganizationNames(names); err != nil {
		return fmt.Errorf("names: '%v': %v", names, err)
	}
	return nil
}
