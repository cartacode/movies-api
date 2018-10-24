package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var categoryRoutes = Routes{
	// swagger:operation GET /category/{ObjectId} category categoryGetId
	// ---
	// summary: Get a category the given ObjectId.
	// description: Get a given category
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/categoryResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CategoryCategoryIDGet",
		strings.ToUpper("Get"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDGet,
	},

	// swagger:operation GET /category/slug/{slug} category categorySlugGetByNAme
	// ---
	// summary: Get a category the given the slug.
	// description: Search for a category by slug
	// parameters:
	// - name: slug
	//   in: path
	//   description: slug
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/categoryResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CategorySlugGet",
		strings.ToUpper("Get"),
		"/category/slug/{slug}",
		controllers.CategorySlugGet,
	},

	// swagger:operation PATCH /category category categoryPatch
	// ---
	// summary: Update a category
	// description: Update a current category
	// parameters:
	// - name: category
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/categoryResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CategoryCategoryIDPatch",
		strings.ToUpper("Patch"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDPatch,
	},

	// swagger:operation GET /category category categoryList
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
	//     "$ref": "#/responses/categoryResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CategoryGet",
		strings.ToUpper("Get"),
		"/category",
		controllers.CategoryGet,
	},

	// swagger:operation POST /category/ category categoryPost
	// ---
	// summary: Post a new category
	// description: Return all categories, paginated
	// parameters:
	// - name: category
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/categoryResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CategoryPost",
		strings.ToUpper("Post"),
		"/category",
		controllers.CategoryPost,
	},

	// swagger:operation DELETE /category/{ObjectId} category categoryDeleteId
	// ---
	// summary: Delete a category the given ObjectId.
	// description: Delete a given category
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
		"CategoryCategoryIdDelete",
		strings.ToUpper("Delete"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDDelete,
	},
}
