package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var customerRoutes = Routes{

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
		"/v1/customer/wishlist",
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
