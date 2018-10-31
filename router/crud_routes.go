package router

import (
	"github.com/VuliTv/go-movie-api/controllers"
)

var crudRoutes = Routes{
	Route{
		"GenericCrudGet",
		"GET",
		"/v1/collection/{collection}",
		controllers.GenericCrudGet,
	},

	// swagger:operation POST /v1/collection/{collection}/ crud crudPaginate
	// ---
	// summary: Get any model type. Searchable
	// description: Return all search results of a model, paginated
	// parameters:
	// - name: collection
	//   in: path
	//   description: Collection Name to search
	//   required: true
	//   schema:
	//    type: string
	//    enum: [movie, scene, volume]
	// responses:
	//   "200":
	//     "$ref": "#/responses/jsonPaginationResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"

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
}
