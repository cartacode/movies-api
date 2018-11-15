package router

import (
	"github.com/VuliTv/go-movie-api/controllers"
)

var searchRoutes = Routes{
	Route{
		"GenericSearchGet",
		"GET",
		"/v1/search/{collection}/",
		controllers.GenericSearchGet,
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
}
