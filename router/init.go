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

	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/requests"

	"github.com/gorilla/mux"
)

var rDB, rError = dbh.NewRedisConnection()
var mDB, mError = dbh.NewMongoDBConnection("router")

// Route --
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes --
type Routes []Route

// NewRouter --
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	routes = append(routes, operationsRoutes...)
	routes = append(routes, crudRoutes...)
	routes = append(routes, searchRoutes...)
	routes = append(routes, playbackRoutes...)
	routes = append(routes, authorizationRoutes...)

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
	message := requests.JSONSuccessResponse{Message: "healthy"}
	js, _ := json.Marshal(message)

	requests.ReturnAPIOK(w, js)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Status",
		"GET",
		"/v1/healthcheck",
		HealthCheck,
	},
}
