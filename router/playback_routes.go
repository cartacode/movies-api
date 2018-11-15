package router

import (
	"github.com/VuliTv/go-movie-api/controllers"
)

var playbackRoutes = Routes{
	Route{
		"SignedS3Playback",
		"GET",
		"/v1/play/{collection}/{objectid}",
		controllers.SignedS3Playback,
	},

	// swagger:operation POST /v1/collection/{collection}/ crud crudPaginate
	// ---
	// summary: Get any model type. Searchable
	// description: Return all playback results of a model, paginated
	// parameters:
	// - name: collection
	//   in: path
	//   description: Collection Name to playback
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
