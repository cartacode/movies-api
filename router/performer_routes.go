package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var performerRoutes = Routes{
	// swagger:operation GET /performer performer performerList
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
	//     "$ref": "#/responses/performerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"PerformerGet",
		strings.ToUpper("Get"),
		"/performer",
		controllers.PerformerGet,
	},

	// swagger:operation DELETE /performer/{ObjectId} performer performerDeleteId
	// ---
	// summary: Delete a performer the given ObjectId.
	// description: Delete a given performer
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
		"PerformerPerformerIDDelete",
		strings.ToUpper("Delete"),
		"/performer/{ObjectId}",
		controllers.PerformerPerformerIDDelete,
	},

	// swagger:operation GET /performer/{ObjectId} performer performerGetId
	// ---
	// summary: Get a performer the given ObjectId.
	// description: Get a given performer
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/performerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"PerformerPerformerIDGet",
		strings.ToUpper("Get"),
		"/performer/{performerID}",
		controllers.PerformerPerformerIDGet,
	},

	// swagger:operation GET /performer/slug/{slug} performer performerSlugGetId
	// ---
	// summary: Get a performer the given the slug.
	// description: Search for a performer by slug
	// parameters:
	// - name: slug
	//   in: path
	//   description: slug
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/performerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"PerformerSlugGet",
		strings.ToUpper("Get"),
		"/performer/slug/{slug}",
		controllers.PerformerSlugGet,
	},

	// swagger:operation PATCH /performer performer performerPatch
	// ---
	// summary: Update a performer
	// description: Update a current performer
	// parameters:
	// - name: performer
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/performerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"PerformerPerformerIDPatch",
		strings.ToUpper("Patch"),
		"/performer/{performerID}",
		controllers.PerformerPerformerIDPatch,
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
	//     "$ref": "#/responses/performerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"PerformerPost",
		strings.ToUpper("Post"),
		"/performer",
		controllers.PerformerPost,
	},
}
