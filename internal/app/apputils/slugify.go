// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apputils

import (
	"fmt"
	"regexp"
	"strings"
)

func Slugify(s string) string {

	// Make the string lowercase
	lower := strings.ToLower(s)

	// Replace spaces with hyphens
	slug := strings.ReplaceAll(lower, " ", "-")

	// Remove special characters
	reg, err := regexp.Compile("[^a-z0-9-]+")
	if err != nil {
		fmt.Println("Regex compile error:", err)
	}
	slug = reg.ReplaceAllString(slug, "")

	return slug
}
