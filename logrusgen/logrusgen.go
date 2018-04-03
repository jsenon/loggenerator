//Package logrusgen
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

package logrusgen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Status Answer
// swagger:response statusjson
type statusjson struct {
	// The Status message
	// in: body
	Code    int32  `json:"statuscode"`
	Message string `json:"statusmessage"`
}

//Generate func
// swagger:route POST /log/logrus log logrus
//
// Generate log with logrus
//
// This will generate log with logrus
//
//     Responses:
//       default: validationError
//       200: statusjson
func Generate(w http.ResponseWriter, req *http.Request) {
	url := "https://myexample.com"
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"url":     url,
		"attempt": 3,
		"backoff": time.Second,
	}).Info("failed to fetch URL")
	log.WithFields(log.Fields{
		"url":     url,
		"attempt": 3,
		"backoff": time.Second,
	}).Warn("Warn failed to fetch URL")
	log.WithFields(log.Fields{
		"url":     url,
		"attempt": 3,
		"backoff": time.Second,
	}).Error("Error failed to fetch URL")
	answerjson := statusjson{
		Code:    200,
		Message: "Log Generated with Logrus",
	}
	b, err := json.Marshal(answerjson)
	if err != nil {
		fmt.Println("error:", err)
	}
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
