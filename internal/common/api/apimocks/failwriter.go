// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apimocks

import (
	"fmt"
	"net/http"
)

// FailWriter implements http.ResponseWriter
type FailWriter struct {
	http.ResponseWriter
}

var _ http.ResponseWriter = (*FailWriter)(nil)

func NewFailWriter(m http.ResponseWriter) *FailWriter {
	return &FailWriter{m}
}

func (fw *FailWriter) Write(_ []byte) (int, error) {
	return 0, fmt.Errorf("simulated write failure")
}
