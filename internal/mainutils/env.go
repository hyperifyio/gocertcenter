// Copyright (c) 2024. Heusala Group <info@hg.fi>. All rights reserved.

package mainutils

import "os"

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
