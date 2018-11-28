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
	auth "gopkg.in/hunterlong/authorizecim.v1"
)

var rDB dbh.RedisHandler
var mDB dbh.MongoDBHandler
var aSession, aError = dbh.NewAuthorizeNetSession()

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
	routes = append(routes, customerPaymentRoutes...)
	routes = append(routes, customerRoutes...)
	routes = append(routes, frontendDataRoutes...)

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

	if _, err := auth.IsConnected(); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
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
