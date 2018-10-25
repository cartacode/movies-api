package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var customerRoutes = Routes{
	// swagger:operation DELETE /customer/{ObjectId} customer customerDeleteId
	// ---
	// summary: Delete a customer the given ObjectId.
	// description: Delete a given customer
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
		"CustomerCustomerIDDelete",
		strings.ToUpper("Delete"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDDelete,
	},

	// swagger:operation GET /customer/{ObjectId} customer customerGetId
	// ---
	// summary: Get a customer the given ObjectId.
	// description: Get a given customer
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerCustomerIDGet",
		strings.ToUpper("Get"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDGet,
	},

	// swagger:operation PATCH /customer customer customerPatch
	// ---
	// summary: Update a customer
	// description: Update a current customer
	// parameters:
	// - name: customer
	//   in: body
	//   description: New CustomerDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Customer"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerCustomerIDPatch",
		strings.ToUpper("Patch"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDPatch,
	},

	// swagger:operation GET /customer customer customerList
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
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerGet",
		strings.ToUpper("Get"),
		"/customer",
		controllers.CustomerGet,
	},

	// swagger:operation POST /customer/ customer customerPost
	// ---
	// summary: Post a new customer
	// description: Return all categories, paginated
	// parameters:
	// - name: customer
	//   in: body
	//   description: New CustomerDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Customer"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerPost",
		strings.ToUpper("Post"),
		"/customer",
		controllers.CustomerPost,
	},

	// swagger:operation POST /customer/wishlist/scene/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a scene
	// description: POST sceneId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistAddSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/scene/{ObjectId}",
		controllers.CustomerWishlistAddSceneItem,
	},

	// swagger:operation Delete /customer/wishlist/scene/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a scene
	// description: Delete ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistDeleteSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/scene/{ObjectId}",
		controllers.CustomerWishlistDeleteSceneItem,
	},

	// swagger:operation POST /customer/wishlist/movie/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a movie
	// description: POST ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistAddSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/movie/{ObjectId}",
		controllers.CustomerWishlistAddSceneItem,
	},

	// swagger:operation Delete /customer/wishlist/movie/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a movie
	// description: Delete ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistDeleteSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/movie/{ObjectId}",
		controllers.CustomerWishlistDeleteSceneItem,
	},

	// swagger:operation POST /customer/wishlist/volume/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a volume
	// description: POST ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistAddSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/volume/{ObjectId}",
		controllers.CustomerWishlistAddSceneItem,
	},

	// swagger:operation Delete /customer/wishlist/volume/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a volume
	// description: Delete ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistDeleteSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/volume/{ObjectId}",
		controllers.CustomerWishlistDeleteSceneItem,
	},

	// swagger:operation POST /customer/wishlist/series/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a series
	// description: POST ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistAddSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/series/{ObjectId}",
		controllers.CustomerWishlistAddSceneItem,
	},

	// swagger:operation Delete /customer/wishlist/series/{ObjectId} customer customerWishSceneAdd
	// ---
	// summary: Post a new customer wishlist for a series
	// description: Delete ObjectId to add item to wishlist
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerWishlistDeleteSceneItem",
		strings.ToUpper("Post"),
		"/customer/wishlist/series/{ObjectId}",
		controllers.CustomerWishlistDeleteSceneItem,
	},
}
