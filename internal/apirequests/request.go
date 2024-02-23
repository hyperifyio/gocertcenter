// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apirequests

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

type RequestImpl struct {
	request *http.Request
}

func NewRequest(
	request *http.Request,
) *RequestImpl {
	return &RequestImpl{request}
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
