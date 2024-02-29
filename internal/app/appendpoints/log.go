// Copyright (c) 2024. Heusala Group Oy <info@heusalagroup.fi>. All rights reserved.

package appendpoints

import (
	"fmt"
	"log"
	"strings"

	"github.com/hyperifyio/gocertcenter/internal/common/api/apitypes"
)

func (c *HttpApiController) logf(request apitypes.Request, format string, args ...interface{}) {

	// Prepend the request method and URL to the format
	fullFormat := "[%s %s]: " + format

	// Prepend the request method and URL to the args
	fullArgs := append([]interface{}{request.GetMethod(), request.GetURL()}, args...)

	// Use log.Printf to log the formatted string
	log.Printf(fullFormat, fullArgs...)

}

func (c *HttpApiController) log(request apitypes.Request, args ...interface{}) {

	// Convert each argument in args to a string using fmt.Sprint
	stringArgs := make([]string, len(args))
	for i, arg := range args {
		stringArgs[i] = fmt.Sprint(arg)
	}

	// Join the string representations of all arguments with spaces
	argsStr := strings.Join(stringArgs, " ")

	// Log the message with the request method and URL, and the joined arguments string
	log.Printf("[%s %s]: %s", request.GetMethod(), request.GetURL(), argsStr)

}
