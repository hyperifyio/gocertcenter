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

func (request *RequestImpl) Body() io.ReadCloser {
	return request.request.Body
}

func (request *RequestImpl) GetBodyBytes() ([]byte, error) {
	bytes, err := io.ReadAll(request.request.Body)
	if err != nil {
		return nil, fmt.Errorf("GetBodyBytes: failed: %w", err)
	}
	defer request.request.Body.Close()
	return bytes, nil
}

func (request *RequestImpl) IsMethodGet() bool {
	return request.request.Method == http.MethodGet
}

func (request *RequestImpl) GetURL() *url.URL {
	// FIXME: Add tests
	return request.request.URL
}

func (request *RequestImpl) GetMethod() string {
	// FIXME: Add tests
	return request.request.Method
}

func (request *RequestImpl) GetVariable(name string) string {
	return mux.Vars(request.request)[name]
}

func NewRequest(
	request *http.Request,
) *RequestImpl {
	return &RequestImpl{request}
}

var _ apitypes.IRequest = (*RequestImpl)(nil)
