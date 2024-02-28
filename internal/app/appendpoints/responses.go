// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *ApiController) sendBadRequest(response apitypes.IResponse, request apitypes.IRequest, publicMessage string, err error) error {
	msg := fmt.Sprintf("[%s %s]: %s", request.GetMethod(), request.GetURL(), publicMessage)
	if err != nil {
		log.Printf("%s: %v", msg, err)
	}
	response.SendError(400, msg)
	return nil
}

func (c *ApiController) sendNotFound(response apitypes.IResponse, request apitypes.IRequest, err error) error {
	publicMsg := fmt.Sprintf("[%s %s]: Not Found", request.GetMethod(), request.GetURL())
	if err != nil {
		log.Printf("%s: %v", publicMsg, err)
	}
	response.SendError(404, publicMsg)
	return nil
}

func (c *ApiController) sendInternalServerError(response apitypes.IResponse, request apitypes.IRequest, err error) error {
	publicMsg := fmt.Sprintf("[%s %s]: Internal Server Error", request.GetMethod(), request.GetURL())
	if err != nil {
		log.Printf("%s: %v", publicMsg, err)
	}
	response.SendError(404, publicMsg)
	return nil
}

func (c *ApiController) sendConflict(response apitypes.IResponse, request apitypes.IRequest, err error, publicMessage string) error {
	publicMsg := fmt.Sprintf("[%s %s]: Conflict: %s", request.GetMethod(), request.GetURL(), publicMessage)
	if err != nil {
		log.Printf("%s: %v", publicMsg, err)
	}
	response.SendError(http.StatusConflict, publicMsg)
	return nil
}

func (c *ApiController) sendOK(response apitypes.IResponse, data interface{}) error {
	response.Send(http.StatusOK, data)
	return nil
}
