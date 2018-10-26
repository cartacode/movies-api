/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	routes := append(routes, categoryRoutes...)
	routes = append(routes, customerRoutes...)
	routes = append(routes, movieRoutes...)
	routes = append(routes, performerRoutes...)
	routes = append(routes, studioRoutes...)
	routes = append(routes, sceneRoutes...)
	routes = append(routes, seriesRoutes...)
	routes = append(routes, volumeRoutes...)
	routes = append(routes, searchRoutes...)
	routes = append(routes, operationsRoutes...)

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
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
