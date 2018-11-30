/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VuliTv/go-movie-api/app/auth"
	"github.com/VuliTv/go-movie-api/app/crud"
	"github.com/VuliTv/go-movie-api/app/customer"
	"github.com/VuliTv/go-movie-api/app/denormalized"
	"github.com/VuliTv/go-movie-api/app/operations"
	"github.com/VuliTv/go-movie-api/app/payments"
	"github.com/VuliTv/go-movie-api/app/search"
	"github.com/VuliTv/go-movie-api/app/webdata"
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/gorilla/mux"
	authorize_net "gopkg.in/hunterlong/authorizecim.v1"
)

var rDB dbh.RedisHandler
var mDB dbh.MongoDBHandler
var aSession, aError = dbh.NewAuthorizeNetSession()
var log = logging.GetProdLog()

func init() {

	if err := rDB.New("router"); err != nil {
		log.Fatalw("redis connection failure. exiting", "error", err)
	}
	if err := mDB.New("router"); err != nil {
		log.Fatalw("mongodb connection failure. exiting", "error", err)
	}

	if aError != nil {
		log.Fatalw("authorize.net connection failure. exiting", "error", aError)
	}
	log.Info("authorize.net connected", aSession, aError)
	if aSession == false {
		log.Fatalw("authorize.net session failure. exiting", "error", aError)
	}
	log.Info("authorize.net session connected")
}

// NewRouter --
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	routes = append(routes, auth.Routes...)
	routes = append(routes, crud.Routes...)
	routes = append(routes, search.Routes...)
	routes = append(routes, operations.Routes...)
	routes = append(routes, payments.Routes...)
	routes = append(routes, customer.Routes...)
	routes = append(routes, webdata.Routes...)
	routes = append(routes, denormalized.Routes...)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = requests.Routes{
	requests.Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	requests.Route{
		Name:        "Status",
		Method:      "GET",
		Pattern:     "/v1/healthcheck",
		HandlerFunc: HealthCheck,
	},
}

// Index --
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "vuli")
}

// HealthCheck --
func HealthCheck(w http.ResponseWriter, r *http.Request) {

	if err := mDB.Session.Ping(); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	if ok := rDB.Ping(); ok == nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, ok.String()))
		return
	}

	if _, err := authorize_net.IsConnected(); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	message := requests.JSONSuccessResponse{Message: "healthy"}
	js, _ := json.Marshal(message)

	requests.ReturnAPIOK(w, js)
}

// Logger --
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()
		inner.ServeHTTP(w, r)

		if r.RequestURI == "/v1/healthcheck" {
			return
		}

		log.Infow("api_call",
			"type", w.Header()["Content-Type"],
			"request", r.Method,
			"uri", r.RequestURI,
			"method", name,
			"response", time.Since(start),
		)
	})
}
