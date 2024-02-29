// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apimocks

import (
	"bytes"
	"io"
	"net/url"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// MockRequest implements the apitypes.Request interface
type MockRequest struct {
	MockIsGet            bool
	MockMethod           string
	MockURL              *url.URL
	MockVars             map[string]string
	MockQueryParams      map[string]string
	MockBodyContent      []byte
	MockBodyContentError error
}

func (m *MockRequest) Header(name string) string {
	// TODO implement me
	panic("implement me")
}

func (m *MockRequest) Body() io.ReadCloser {
	return io.NopCloser(bytes.NewBufferString(string(m.MockBodyContent)))
}

func (m *MockRequest) BodyBytes() ([]byte, error) {
	return m.MockBodyContent, m.MockBodyContentError
}

func (m *MockRequest) IsGet() bool {
	return m.MockIsGet
}

func (m *MockRequest) URL() *url.URL {
	return m.MockURL
}

func (m *MockRequest) Method() string {
	return m.MockMethod
}

func (m *MockRequest) Variable(name string) string {
	return m.MockVars[name]
}

func (m *MockRequest) QueryParam(name string) string {
	return m.MockQueryParams[name]
}

var _ apitypes.Request = (*MockRequest)(nil)
