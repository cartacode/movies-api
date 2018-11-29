package search

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
)

// Routes --
var Routes = requests.Routes{
	requests.Route{
		Name:        "GenericSearchGet",
		Method:      "GET",
		Pattern:     "/v1/search/{collection}/",
		HandlerFunc: GenericSearchGet,
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
