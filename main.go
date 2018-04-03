//go:generate swagger generate spec
// Package main Loggenerator.
//
// the purpose of this application is to provide an interface to Generate different Log
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Julien SENON <julien.senon@gmail.com>
package main

import (
	"net/http"

	"github.com/jsenon/loggenerator/logrusgen"
	"github.com/jsenon/loggenerator/web"

	"github.com/jsenon/loggenerator/zapgen"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	// Remove CORS Header check to allow swagger and application on same host and port
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	// To be changed
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "PATCH"})

	// Web Part
	// r.HandleFunc("/index", web.Index)
	r.HandleFunc("/healthz", web.Healthz)

	// Swagger
	// r.HandleFunc("/api", web.ApiHelp)

	// API Part
	r.HandleFunc("/log/logrus", logrusgen.Generate).Methods("POST")
	r.HandleFunc("/log/zapsugar", zapgen.Generatesugar).Methods("POST")
	r.HandleFunc("/log/zaplogger", zapgen.Generatelogger).Methods("POST")

	// Static dir
	http.ListenAndServe(":9030", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
