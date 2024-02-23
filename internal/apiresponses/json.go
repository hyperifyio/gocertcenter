// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package apiresponses

import (
	"encoding/json"
	"github.com/hyperifyio/gocertcenter/internal/dtos"
	"net/http"
)

type JSONResponse struct {
	writer http.ResponseWriter
}

func NewJSONResponse(w http.ResponseWriter) *JSONResponse {
	return &JSONResponse{writer: w}
}

func (sender *JSONResponse) Send(statusCode int, data interface{}) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		// FIXME: Add test
		http.Error(sender.writer, "Error creating JSON writer", http.StatusInternalServerError)
		return
	}

	jsonData = append(jsonData, '\n')
	sender.writer.Header().Set("Content-Type", "application/json")
	sender.writer.WriteHeader(statusCode)
	_, err = sender.writer.Write(jsonData)

	if err != nil {
		// FIXME: Add test
		http.Error(sender.writer, "Error writing JSON writer", http.StatusInternalServerError)
		return
	}

}

func (sender *JSONResponse) SendError(statusCode int, error string) {
	if error == "" {
		error = http.StatusText(statusCode)
	}
	data := dtos.NewErrorDTO(statusCode, error)
	sender.Send(statusCode, data)
}

func (sender *JSONResponse) SendMethodNotSupportedError() {
	sender.SendError(http.StatusMethodNotAllowed, "")
}

func (sender *JSONResponse) SendNotFoundError() {
	sender.SendError(http.StatusNotFound, "")
}

func (sender *JSONResponse) SendConflictError(error string) {
	sender.SendError(http.StatusConflict, error)
}

func (sender *JSONResponse) SendInternalServerError(error string) {
	sender.SendError(http.StatusInternalServerError, error)
}
