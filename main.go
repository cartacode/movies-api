// Vuli Platform API
//
// This is the core service for vuli content and member managment
//
//     Schemes: https
//     Host: api-stage.vuli.tv
//     Version: 0.0.1
//     License: Proprietary
//     Contact: Richard<richard@vuli.tv>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta

package main

import (
	"net/http"

	"github.com/VuliTv/api/libs/logging"
	"github.com/VuliTv/api/router"
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

	r := router.NewRouter()

	log.Fatal(http.ListenAndServe(":3001", r))
}
