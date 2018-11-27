package router

import (
	"github.com/VuliTv/go-movie-api/controllers"
)

var crudRoutes = Routes{
	Route{
		"GenericCrudGet",
		"GET",
		"/v1/collection/{collection}",
		validateTokenMiddleware(controllers.GenericCrudGet),
	},

	Route{
		"GenericCrudIDGet",
		"GET",
		"/v1/collection/{collection}/{objectID}",
		validateTokenMiddleware(controllers.GenericCrudIDGet),
	},

	Route{
		"GenericCrudIDDelete",
		"DELETE",
		"/v1/collection/{collection}/{objectID}",
		validateTokenMiddleware(controllers.GenericCrudIDDelete),
	},

	Route{
		"GenericCrudPost",
		"POST",
		"/v1/collection/{collection}",
		validateTokenMiddleware(controllers.GenericCrudPost),
	},

	Route{
		"GenericCrudIDPatch",
		"PATCH",
		"/v1/collection/{collection}/{objectID}",
		validateTokenMiddleware(controllers.GenericCrudIDPatch),
	},
}
