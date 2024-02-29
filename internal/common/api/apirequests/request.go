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

type RequestImpl struct {
	request *http.Request
}

func (r *RequestImpl) GetHeader(name string) string {
	return r.request.Header.Get(name)
}

func (r *RequestImpl) Body() io.ReadCloser {
	return r.request.Body
}

func (r *RequestImpl) GetBodyBytes() ([]byte, error) {
	bytes, err := io.ReadAll(r.request.Body)
	if err != nil {
		return nil, fmt.Errorf("GetBodyBytes: failed: %w", err)
	}
	defer r.request.Body.Close()
	return bytes, nil
}

func (r *RequestImpl) IsMethodGet() bool {
	return r.request.Method == http.MethodGet
}

func (r *RequestImpl) GetURL() *url.URL {
	return r.request.URL
}

func (r *RequestImpl) GetMethod() string {
	// FIXME: Add tests
	return r.request.Method
}

func (r *RequestImpl) GetVariable(name string) string {
	return mux.Vars(r.request)[name]
}

func (r *RequestImpl) GetQueryParam(name string) string {
	queryParams := r.request.URL.Query()
	return queryParams.Get(name)
}

func NewRequest(
	request *http.Request,
) *RequestImpl {
	return &RequestImpl{request}
}

var _ apitypes.IRequest = (*RequestImpl)(nil)
