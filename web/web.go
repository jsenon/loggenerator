// Package Web Loggenerator.
//
// the purpose of this package is to provide Api Interface
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost:9030
//     BasePath: /log
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Status Answer
// swagger:response statusjson
type statusjson struct {
	// The Status message
	// in: body
	Code    int32  `json:"statuscode"`
	Message string `json:"statusmessage"`
}

// swagger:route POST /healthz health
//
// Check health of platform
//
// This will sent information if service is up and running.
//
//     Responses:
//       default: validationError
//       200: statusjson
func Healthz(w http.ResponseWriter, req *http.Request) {
	answerjson := statusjson{
		Code:    200,
		Message: "Log Generated with ZapSuggar",
	}
	b, err := json.Marshal(answerjson)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Write(b)
}
