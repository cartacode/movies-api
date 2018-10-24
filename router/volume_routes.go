package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var volumeRoutes = Routes{

	// swagger:operation GET /volume volume volumeList
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
	//     "$ref": "#/responses/volumeResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"VolumeGet",
		strings.ToUpper("Get"),
		"/volume",
		controllers.VolumeGet,
	},

	// swagger:operation POST /volume/ volume volumePost
	// ---
	// summary: Post a new volume
	// description: Return all categories, paginated
	// parameters:
	// - name: volume
	//   in: body
	//   description: New VolumeDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Volume"
	// responses:
	//   "200":
	//     "$ref": "#/responses/volumeResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"VolumePost",
		strings.ToUpper("Post"),
		"/volume",
		controllers.VolumePost,
	},

	// swagger:operation DELETE /volume/{ObjectId} volume volumeDeleteId
	// ---
	// summary: Delete a volume the given ObjectId.
	// description: Delete a given volume
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
		"VolumeVolumeIDDelete",
		strings.ToUpper("Delete"),
		"/volume/{VolumeID}",
		controllers.VolumeVolumeIDDelete,
	},

	// swagger:operation GET /volume/{ObjectId} volume volumeGetId
	// ---
	// summary: Get a volume the given ObjectId.
	// description: Get a given volume
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/volumeResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"VolumeVolumeIDGet",
		strings.ToUpper("Get"),
		"/volume/{VolumeID}",
		controllers.VolumeVolumeIDGet,
	},

	// swagger:operation PATCH /volume volume volumePatch
	// ---
	// summary: Update a volume
	// description: Update a current volume
	// parameters:
	// - name: volume
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/volumeResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"VolumeVolumeIDPatch",
		strings.ToUpper("Patch"),
		"/volume/{VolumeID}",
		controllers.VolumeVolumeIDPatch,
	},
}
