/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/VuliTv/api/controllers"
	"github.com/gorilla/mux"
)

// Route --
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes --
type Routes []Route

// NewRouter --
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index --
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
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

	/*
	 * Customer Controllers
	 */

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
	/*
	 * Movie Controllers
	 */

	// swagger:operation GET /movie movie movieList
	// ---
	// summary: List all of the movies in a pagination response.
	// description: Return all movies, paginated
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
	//     "$ref": "#/responses/movieResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"MovieGet",
		strings.ToUpper("Get"),
		"/movie",
		controllers.MovieGet,
	},

	// swagger:operation DELETE /movie/{ObjectId} movie movieDeleteId
	// ---
	// summary: Delete a movie the given ObjectId.
	// description: Delete a given movie
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
		"MovieMovieIDDelete",
		strings.ToUpper("Delete"),
		"/movie/{ObjectId}",
		controllers.MovieMovieIDDelete,
	},

	// swagger:operation GET /movie/{ObjectId} movie movieGetId
	// ---
	// summary: Get a movie the given ObjectId.
	// description: Get a given movie
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB ObjectId
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/movieResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"MovieMovieIDGet",
		strings.ToUpper("Get"),
		"/movie/{ObjectId}",
		controllers.MovieMovieIDGet,
	},

	// swagger:operation GET /movie/slug/{slug} movie movieSlugGetId
	// ---
	// summary: Get a movie the given the slug.
	// description: Search for a movie by slug
	// parameters:
	// - name: slug
	//   in: path
	//   description: slug
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/movieResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"MovieSlugGet",
		strings.ToUpper("Get"),
		"/movie/slug/{slug}",
		controllers.MovieSlugGet,
	},

	// swagger:operation PATCH /movie movie moviePatch
	// ---
	// summary: Update a movie
	// description: Update a current movie
	// parameters:
	// - name: movie
	//   in: body
	//   description: New CategoryDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Category"
	// responses:
	//   "200":
	//     "$ref": "#/responses/movieResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"MovieMovieIDPatch",
		strings.ToUpper("Patch"),
		"/movie/{ObjectId}",
		controllers.MovieMovieIDPatch,
	},

	// swagger:operation POST /movie/ movie moviePost
	// ---
	// summary: Post a new movie
	// description: Return all categories, paginated
	// parameters:
	// - name: movie
	//   in: body
	//   description: New MovieDocument
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/Movie"
	// responses:
	//   "200":
	//     "$ref": "#/responses/movieResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"MoviePost",
		strings.ToUpper("Post"),
		"/movie",
		controllers.MoviePost,
	},

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

	/*
	 * Search Controllers
	 */

	// swagger:operation GET /search/movie search movieSearchList
	// ---
	// summary: List all found movies in a pagination response.
	// description: Return all movies found, paginated
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
	//     "$ref": "#/responses/movieResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"MoviesSearchGet",
		strings.ToUpper("Get"),
		"/search/movie",
		controllers.MoviesSearchGet,
	},

	// swagger:operation GET /search/scene search sceneSearchList
	// ---
	// summary: List all found scenes in a pagination response.
	// description: Return all scenes found, paginated
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
		"MoviesSearchGet",
		strings.ToUpper("Get"),
		"/search/scene",
		controllers.MoviesSearchGet,
	},
}
