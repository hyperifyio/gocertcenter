// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apimocks

import (
	"bytes"
	"io"
	"net/url"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

// MockRequest implements the apitypes.IRequest interface
type MockRequest struct {
	IsGet            bool
	Method           string
	URL              *url.URL
	Vars             map[string]string
	QueryParams      map[string]string
	BodyContent      []byte
	BodyContentError error
}

func (m *MockRequest) Body() io.ReadCloser {
	return io.NopCloser(bytes.NewBufferString(string(m.BodyContent)))
}

func (m *MockRequest) GetBodyBytes() ([]byte, error) {
	return m.BodyContent, m.BodyContentError
}

func (m *MockRequest) IsMethodGet() bool {
	return m.IsGet
}

func (m *MockRequest) GetURL() *url.URL {
	return m.URL
}

func (m *MockRequest) GetMethod() string {
	return m.Method
}

func (m *MockRequest) GetVariable(name string) string {
	return m.Vars[name]
}

func (m *MockRequest) GetQueryParam(name string) string {
	return m.QueryParams[name]
}

var _ apitypes.IRequest = (*MockRequest)(nil)
