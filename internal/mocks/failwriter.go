// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package mocks

import (
	"fmt"
	"net/http"
)

type FailWriter struct {
	http.ResponseWriter
}

func NewFailWriter(m http.ResponseWriter) *FailWriter {
	return &FailWriter{m}
}

func (fw *FailWriter) Write(_ []byte) (int, error) {
	return 0, fmt.Errorf("simulated write failure")
}
