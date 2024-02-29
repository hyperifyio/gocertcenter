// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// isMn checks if a rune is a non-spacing mark
func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: Mark, Nonspacing
}

func Slugify(s string) string {

	// Normalize string to NFD (Normalization Form Decomposition)
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, s)

	var builder strings.Builder
	for _, char := range strings.ToLower(result) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || char == '-' {
			builder.WriteRune(char)
		} else {
			builder.WriteRune('-')
		}
	}

	// Convert multiple hyphens to a single hyphen
	slug := builder.String()
	multiHyphenRegex := regexp.MustCompile(`-+`)
	slug = multiHyphenRegex.ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	return slug
}
