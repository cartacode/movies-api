package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var studioRoutes = Routes{

	// swagger:operation GET /studio studio studioList
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
	//     "$ref": "#/responses/studioResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"StudioGet",
		strings.ToUpper("Get"),
		"/studio",
		controllers.StudioGet,
	},

	// swagger:operation POST /studio/ studio studioPost
	// ---
	// summary: Post a new studio
	// description: Return all categories, paginated
	// parameters:
	// - name: studio
	//   in: body
	//   description: New StudioDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Studio"
	// responses:
	//   "200":
	//     "$ref": "#/responses/studioResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"StudioPost",
		strings.ToUpper("Post"),
		"/studio",
		controllers.StudioPost,
	},

	// swagger:operation DELETE /studio/{ObjectId} studio studioDeleteId
	// ---
	// summary: Delete a studio the given ObjectId.
	// description: Delete a given studio
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
		"StudioStudioIDDelete",
		strings.ToUpper("Delete"),
		"/studio/{studioID}",
		controllers.StudioStudioIDDelete,
	},

	// swagger:operation GET /studio/{ObjectId} studio studioGetId
	// ---
	// summary: Get a studio the given ObjectId.
	// description: Get a given studio
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/studioResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"StudioStudioIDGet",
		strings.ToUpper("Get"),
		"/studio/{studioID}",
		controllers.StudioStudioIDGet,
	},

	// swagger:operation GET /studio/slug/{slug} studio studioSlugGetId
	// ---
	// summary: Get a studio the given the slug.
	// description: Search for a studio by slug
	// parameters:
	// - name: slug
	//   in: path
	//   description: slug
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/studioResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"StudioSlugGet",
		strings.ToUpper("Get"),
		"/studio/slug/{slug}",
		controllers.StudioSlugGet,
	},

	// swagger:operation PATCH /studio studio studioPatch
	// ---
	// summary: Update a studio
	// description: Update a current studio
	// parameters:
	// - name: studio
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/studioResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"StudioStudioIDPatch",
		strings.ToUpper("Patch"),
		"/studio/{studioID}",
		controllers.StudioStudioIDPatch,
	},
}
