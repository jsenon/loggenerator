//Package zapgen
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

package zapgen

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// Status Answer
// swagger:response statusjson
type statusjson struct {
	// The Status message
	// in: body
	Code    int32  `json:"statuscode"`
	Message string `json:"statusmessage"`
}

//Generatesugar func
// swagger:route POST /log/zapsuggar log zap
//
// Generate Log with Sugar
//
// This will sent log with sugar format.
//
//     Responses:
//       default: validationError
//       200: statusjson
func Generatesugar(w http.ResponseWriter, req *http.Request) {
	url := "https://myexample.com"
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("error:", err)
	}
	defer logger.Sync() // nolint: errcheck
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Warnw("Warn failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Errorw("Error failed to fetch URL",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	answerjson := statusjson{
		Code:    200,
		Message: "Log Generated with ZapSuggar",
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

//Generatelogger func
// swagger:route POST /log/zaplogger log zap
//
// Generate Log with Logger
//
// This will sent log with logger format.
//
//     Responses:
//       default: validationError
//       200: statusjson
func Generatelogger(w http.ResponseWriter, req *http.Request) {
	url := "https://myexample.com"
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println("error:", err)
	}
	defer logger.Sync() // nolint: errcheck
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Warn("Warn failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Error("Error failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	answerjson := statusjson{
		Code:    200,
		Message: "Log Generated with ZapLogger",
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
