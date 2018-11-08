package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var operationsRoutes = Routes{
	// swagger:operation POST /performer/ performer performerPost
	// ---
	// summary: Post a new performer
	// description: Return all categories, paginated
	// parameters:
	// - name: performer
	//   in: body
	//   description: New PerformerDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Performer"
	// responses:
	//   "200":
	//     "$ref": "#/responses/JSONPaginationResponse"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"OperationsUploadImage",
		strings.ToUpper("Post"),
		"/operations/upload/{field}/{collection}/{objectid}",
		controllers.OperationsUploadImage,
	},
	// swagger:operation POST /performer/ performer performerPost
	// ---
	// summary: Post a new performer
	// description: Return all categories, paginated
	// parameters:
	// - name: performer
	//   in: body
	//   description: New PerformerDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Performer"
	// responses:
	//   "200":
	//     "$ref": "#/responses/JSONPaginationResponse"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"OperationsUploadTrailer",
		strings.ToUpper("Post"),
		"/operations/upload/trailer/{collection}/{objectid}/{slug}",
		controllers.OperationsUploadTrailer,
	},
}
