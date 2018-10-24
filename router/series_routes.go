package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var seriesRoutes = Routes{
	// swagger:operation GET /series series seriesList
	// ---
	// summary: List all of the categories in a pagination response.
	// description: Return all categories, paginated
	// parameters:
	// - in: query
	//   name: page
	//   schema:
	//     type: integer
	//   description: The number of pages to skip before starting to collect the result set
	// - in: query
	//   name: perpage
	//   schema:
	//     type: integer
	//   description: The numbers of items to return per page
	// responses:
	//   "200":
	//     "$ref": "#/responses/seriesResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SeriesGet",
		strings.ToUpper("Get"),
		"/series",
		controllers.SeriesGet,
	},

	// swagger:operation POST /series/ series seriesPost
	// ---
	// summary: Post a new series
	// description: Return all categories, paginated
	// parameters:
	// - name: series
	//   in: body
	//   description: New SeriesDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Series"
	// responses:
	//   "200":
	//     "$ref": "#/responses/seriesResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SeriesPost",
		strings.ToUpper("Post"),
		"/series",
		controllers.SeriesPost,
	},

	// swagger:operation DELETE /series/{ObjectId} series seriesDeleteId
	// ---
	// summary: Delete a series the given ObjectId.
	// description: Delete a given series
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/ok"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SeriesSeriesIDDelete",
		strings.ToUpper("Delete"),
		"/series/{SeriesID}",
		controllers.SeriesSeriesIDDelete,
	},

	// swagger:operation GET /series/{ObjectId} series seriesGetId
	// ---
	// summary: Get a series the given ObjectId.
	// description: Get a given series
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/seriesResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SeriesSeriesIDGet",
		strings.ToUpper("Get"),
		"/series/{SeriesID}",
		controllers.SeriesSeriesIDGet,
	},

	// swagger:operation PATCH /series series seriesPatch
	// ---
	// summary: Update a series
	// description: Update a current series
	// parameters:
	// - name: series
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/seriesResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SeriesSeriesIDPatch",
		strings.ToUpper("Patch"),
		"/series/{SeriesID}",
		controllers.SeriesSeriesIDPatch,
	},
}
