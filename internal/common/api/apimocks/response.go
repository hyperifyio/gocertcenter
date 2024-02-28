// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apimocks

import (
	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// MockResponse implements the apitypes.IResponse interface
type MockResponse struct {
	SentData                interface{}
	SentStatusCode          int
	SentErrorMessage        string
	NotFoundError           bool
	MethodNotSupportedError bool
	ConflictError           string
	InternalServerError     string
}

func (m *MockResponse) SendBytes(bytes []byte) error {
	// TODO implement me
	panic("implement me")
}

func (m *MockResponse) SetHeader(name, value string) {
	// TODO implement me
	panic("implement me")
}

func (m *MockResponse) Send(statusCode int, data interface{}) {
	m.SentStatusCode = statusCode
	m.SentData = data
}

func (m *MockResponse) SendError(statusCode int, errorMessage string) {
	m.SentStatusCode = statusCode
	m.SentErrorMessage = errorMessage
}

func (m *MockResponse) SendMethodNotSupportedError() {
	m.MethodNotSupportedError = true
}

func (m *MockResponse) SendNotFoundError() {
	m.NotFoundError = true
}

func (m *MockResponse) SendConflictError(error string) {
	m.ConflictError = error
}

func (m *MockResponse) SendInternalServerError(error string) {
	m.InternalServerError = error
}

var _ apitypes.IResponse = (*MockResponse)(nil)
