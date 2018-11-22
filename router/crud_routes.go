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
		"GenericCrudPost",
		"POST",
		"/v1/collection/{collection}",
		controllers.GenericCrudPost,
	},

	Route{
		"GenericCrudIDGet",
		"GET",
		"/v1/collection/{collection}/{objectid}",
		controllers.GenericCrudIDGet,
	},

	Route{
		"GenericCrudIDDelete",
		"DELETE",
		"/v1/collection/{collection}/{objectid}",
		controllers.GenericCrudIDDelete,
	},

	Route{
		"GenericCrudIDPatch",
		"PATCH",
		"/v1/collection/{collection}/{objectid}",
		controllers.GenericCrudIDPatch,
	},
}
