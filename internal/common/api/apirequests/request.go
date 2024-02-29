// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apirequests

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

type HttpRequest struct {
	request *http.Request
}

func (r *HttpRequest) GetHeader(name string) string {
	return r.request.Header.Get(name)
}

func (r *HttpRequest) Body() io.ReadCloser {
	return r.request.Body
}

func (r *HttpRequest) GetBodyBytes() ([]byte, error) {
	bytes, err := io.ReadAll(r.request.Body)
	if err != nil {
		return nil, fmt.Errorf("GetBodyBytes: failed: %w", err)
	}
	_ = r.request.Body.Close()
	return bytes, nil
}

func (r *HttpRequest) IsMethodGet() bool {
	return r.request.Method == http.MethodGet
}

func (r *HttpRequest) GetURL() *url.URL {
	return r.request.URL
}

func (r *HttpRequest) GetMethod() string {
	return r.request.Method
}

func (r *HttpRequest) GetVariable(name string) string {
	return mux.Vars(r.request)[name]
}

func (r *HttpRequest) GetQueryParam(name string) string {
	queryParams := r.request.URL.Query()
	return queryParams.Get(name)
}

func NewRequest(
	request *http.Request,
) *HttpRequest {
	return &HttpRequest{request}
}

var _ apitypes.Request = (*HttpRequest)(nil)
