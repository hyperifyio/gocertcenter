// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *HttpApiController) badRequest(response apitypes.Response, request apitypes.Request, publicMessage string, err error) error {
	msg := fmt.Sprintf("[%s %s]: %s", request.Method(), request.URL(), publicMessage)
	if err != nil {
		log.Printf("%s: %v", msg, err)
	}
	response.SendError(400, msg)
	return nil
}

func (c *HttpApiController) notFound(response apitypes.Response, request apitypes.Request, err error) error {
	publicMsg := fmt.Sprintf("[%s %s]: Not Found", request.Method(), request.URL())
	if err != nil {
		log.Printf("%s: %v", publicMsg, err)
	}
	response.SendError(404, publicMsg)
	return nil
}

func (c *HttpApiController) internalServerError(response apitypes.Response, request apitypes.Request, err error) error {
	publicMsg := fmt.Sprintf("[%s %s]: Internal Server Error", request.Method(), request.URL())
	if err != nil {
		log.Printf("%s: %v", publicMsg, err)
	}
	response.SendError(404, publicMsg)
	return nil
}

func (c *HttpApiController) conflict(response apitypes.Response, request apitypes.Request, err error, publicMessage string) error {
	publicMsg := fmt.Sprintf("[%s %s]: Conflict: %s", request.Method(), request.URL(), publicMessage)
	if err != nil {
		log.Printf("%s: %v", publicMsg, err)
	}
	response.SendError(http.StatusConflict, publicMsg)
	return nil
}

func (c *HttpApiController) ok(response apitypes.Response, data interface{}) error {
	response.Send(http.StatusOK, data)
	return nil
}
