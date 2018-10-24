package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var sceneRoutes = Routes{
	// swagger:operation GET /scene scene sceneList
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
	//     "$ref": "#/responses/sceneResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SceneGet",
		strings.ToUpper("Get"),
		"/scene",
		controllers.SceneGet,
	},

	// swagger:operation POST /scene/ scene scenePost
	// ---
	// summary: Post a new scene
	// description: Return all categories, paginated
	// parameters:
	// - name: scene
	//   in: body
	//   description: New SceneDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Scene"
	// responses:
	//   "200":
	//     "$ref": "#/responses/sceneResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"ScenePost",
		strings.ToUpper("Post"),
		"/scene",
		controllers.ScenePost,
	},

	// swagger:operation DELETE /scene/{ObjectId} scene sceneDeleteId
	// ---
	// summary: Delete a scene the given ObjectId.
	// description: Delete a given scene
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
		"SceneSceneIDDelete",
		strings.ToUpper("Delete"),
		"/scene/{SceneID}",
		controllers.SceneSceneIDDelete,
	},

	// swagger:operation GET /scene/{ObjectId} scene sceneGetId
	// ---
	// summary: Get a scene the given ObjectId.
	// description: Get a given scene
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/sceneResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SceneSceneIDGet",
		strings.ToUpper("Get"),
		"/scene/{SceneID}",
		controllers.SceneSceneIDGet,
	},

	// swagger:operation PATCH /scene scene scenePatch
	// ---
	// summary: Update a scene
	// description: Update a current scene
	// parameters:
	// - name: scene
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/sceneResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"SceneSceneIDPatch",
		strings.ToUpper("Patch"),
		"/scene/{SceneID}",
		controllers.SceneSceneIDPatch,
	},
}
