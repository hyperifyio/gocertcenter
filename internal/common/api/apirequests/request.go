// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apirequests

import (
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

type RequestImpl struct {
	request *http.Request
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

func (request *RequestImpl) GetVars() map[string]string {
	// FIXME: Add tests
	return mux.Vars(request.request)
}

func NewRequest(
	request *http.Request,
) *RequestImpl {
	return &RequestImpl{request}
}

var _ apitypes.IRequest = (*RequestImpl)(nil)
