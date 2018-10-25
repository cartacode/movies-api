// Copyright 2017 Emir Ribic. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Golang SwaggerUI example
//
// This documentation describes example APIs found under https://github.com/ribice/golang-swaggerui-example
//
//     Schemes: https
//     BasePath: /v1
//     Version: 1.0.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Emir Ribic <ribice@gmail.com> https://ribice.ba
//     Host: ribice.ba/goswagg
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - bearer
//
//     SecurityDefinitions:
//     bearer:
//          type: apiKey
//          name: Authorization
//          in: header
//
// swagger:meta

package main

import (
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/logging"
	"github.com/VuliTv/go-movie-api/router"
	"github.com/gorilla/handlers"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
)

var log = logging.GetProdLog()

func main() {
	logging.SetLevel(logging.DEBUG)
	log.Info("Server started")

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r)))
}
