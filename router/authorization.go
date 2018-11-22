package router

import (
	"github.com/VuliTv/go-movie-api/controllers"
)

var authorizationRoutes = Routes{
	Route{
		"Authorization",
		"GET",
		"/v1/authorize",
		controllers.Authorization,
	},
}
