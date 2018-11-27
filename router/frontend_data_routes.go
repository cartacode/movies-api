package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var frontendDataRoutes = Routes{

	Route{
		"DataMovieTray",
		strings.ToUpper("get"),
		"/v1/data/movie",
		validateTokenMiddleware(controllers.DataMovieTray),
	},
	Route{
		"DataSceneTray",
		strings.ToUpper("get"),
		"/v1/data/scene",
		validateTokenMiddleware(controllers.DataSceneTray),
	},
}
