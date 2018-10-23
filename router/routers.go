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

	// swagger:operation DELETE /category/{id} category categoryDeleteId
	// ---
	// summary: Delete a category the given id.
	// description: Delete a given category
	// parameters:
	// - name: category
	//   in: path
	//   description: MongoDB ObjectId
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/categoryResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CategoryCategoryIdDelete",
		strings.ToUpper("Delete"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDDelete,
	},

	// swagger:operation GET /category/{id} category categoryGetId
	// ---
	// summary: Get a category the given id.
	// description: Get a given category
	// parameters:
	// - name: category
	//   in: path
	//   description: MongoDB ObjectId
	//   type: string
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

	// swagger:operation GET /category/slug/{id} category categorySlugGetId
	// ---
	// summary: Get a category the given the slug.
	// description: Search for a category by slug
	// parameters:
	// - name: category
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

	// swagger:operation PATCH /category/ category categoryPatch
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

	// swagger:operation GET /category/ category categoryList
	// ---
	// summary: List all of the categories in a pagination response.
	// description: Return all categories, paginated
	// parameters:
	// - name: category
	//   in: path
	//   description: MongoDB ObjectId
	//   type: string
	//   required: true
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

	// swagger:operation DELETE /customer/{id} customer customerDeleteId
	// ---
	// summary: Delete a customer the given id.
	// description: Delete a given customer
	// parameters:
	// - name: customer
	//   in: path
	//   description: MongoDB ObjectId
	//   type: string
	//   required: true
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerCustomerIDDelete",
		strings.ToUpper("Delete"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDDelete,
	},

	// swagger:operation GET /customer/{id} customer customerGetId
	// ---
	// summary: Get a customer the given id.
	// description: Get a given customer
	// parameters:
	// - name: customer
	//   in: path
	//   description: MongoDB ObjectId
	//   type: string
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

	// swagger:operation PATCH /customer/ customer customerPatch
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

	// swagger:operation GET /customer/ customer customerList
	// ---
	// summary: List all of the categories in a pagination response.
	// description: Return all categories, paginated
	// parameters:
	// - name: customer
	//   in: path
	//   description: MongoDB ObjectId
	//   type: string
	//   required: true
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
	//   description: New CategoryDocument
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

	Route{
		"MovieGet",
		strings.ToUpper("Get"),
		"/movie",
		controllers.MovieGet,
	},

	Route{
		"MovieMovieIDDelete",
		strings.ToUpper("Delete"),
		"/movie/{movieID}",
		controllers.MovieMovieIDDelete,
	},

	Route{
		"MovieMovieIDGet",
		strings.ToUpper("Get"),
		"/movie/{movieID}",
		controllers.MovieMovieIDGet,
	},

	// swagger:operation GET /movie/slug/{id} movie movieSlugGetId
	// ---
	// summary: Get a movie the given the slug.
	// description: Search for a movie by slug
	// parameters:
	// - name: movie
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

	Route{
		"MovieMovieIDPatch",
		strings.ToUpper("Patch"),
		"/movie/{movieID}",
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

	Route{
		"PerformerGet",
		strings.ToUpper("Get"),
		"/performer",
		controllers.PerformerGet,
	},

	Route{
		"PerformerPerformerIDDelete",
		strings.ToUpper("Delete"),
		"/performer/{performerID}",
		controllers.PerformerPerformerIDDelete,
	},

	Route{
		"PerformerPerformerIDGet",
		strings.ToUpper("Get"),
		"/performer/{performerID}",
		controllers.PerformerPerformerIDGet,
	},

	Route{
		"PerformerSlugGet",
		strings.ToUpper("Get"),
		"/performer/slug/{slug}",
		controllers.PerformerSlugGet,
	},

	Route{
		"PerformerPerformerIDPatch",
		strings.ToUpper("Patch"),
		"/performer/{performerID}",
		controllers.PerformerPerformerIDPatch,
	},

	Route{
		"PerformerPost",
		strings.ToUpper("Post"),
		"/performer",
		controllers.PerformerPost,
	},

	Route{
		"SceneGet",
		strings.ToUpper("Get"),
		"/scene",
		controllers.SceneGet,
	},

	Route{
		"ScenePost",
		strings.ToUpper("Post"),
		"/scene",
		controllers.ScenePost,
	},

	Route{
		"SceneSceneIDDelete",
		strings.ToUpper("Delete"),
		"/scene/{SceneID}",
		controllers.SceneSceneIDDelete,
	},

	Route{
		"SceneSceneIDGet",
		strings.ToUpper("Get"),
		"/scene/{SceneID}",
		controllers.SceneSceneIDGet,
	},

	Route{
		"SceneSceneIDPatch",
		strings.ToUpper("Patch"),
		"/scene/{SceneID}",
		controllers.SceneSceneIDPatch,
	},

	Route{
		"SeriesGet",
		strings.ToUpper("Get"),
		"/series",
		controllers.SeriesGet,
	},

	Route{
		"SeriesPost",
		strings.ToUpper("Post"),
		"/series",
		controllers.SeriesPost,
	},

	Route{
		"SeriesSeriesIDDelete",
		strings.ToUpper("Delete"),
		"/series/{SeriesID}",
		controllers.SeriesSeriesIDDelete,
	},

	Route{
		"SeriesSeriesIDGet",
		strings.ToUpper("Get"),
		"/series/{SeriesID}",
		controllers.SeriesSeriesIDGet,
	},

	Route{
		"SeriesSeriesIDPatch",
		strings.ToUpper("Patch"),
		"/series/{SeriesID}",
		controllers.SeriesSeriesIDPatch,
	},

	Route{
		"StudioGet",
		strings.ToUpper("Get"),
		"/studio",
		controllers.StudioGet,
	},

	Route{
		"StudioPost",
		strings.ToUpper("Post"),
		"/studio",
		controllers.StudioPost,
	},

	Route{
		"StudioStudioIDDelete",
		strings.ToUpper("Delete"),
		"/studio/{studioID}",
		controllers.StudioStudioIDDelete,
	},

	Route{
		"StudioStudioIDGet",
		strings.ToUpper("Get"),
		"/studio/{studioID}",
		controllers.StudioStudioIDGet,
	},

	Route{
		"StudioSlugGet",
		strings.ToUpper("Get"),
		"/studio/slug/{slug}",
		controllers.StudioSlugGet,
	},

	Route{
		"StudioStudioIDPatch",
		strings.ToUpper("Patch"),
		"/studio/{studioID}",
		controllers.StudioStudioIDPatch,
	},

	Route{
		"VolumeGet",
		strings.ToUpper("Get"),
		"/volume",
		controllers.VolumeGet,
	},

	Route{
		"VolumePost",
		strings.ToUpper("Post"),
		"/volume",
		controllers.VolumePost,
	},

	Route{
		"VolumeVolumeIDDelete",
		strings.ToUpper("Delete"),
		"/volume/{VolumeID}",
		controllers.VolumeVolumeIDDelete,
	},

	Route{
		"VolumeVolumeIDGet",
		strings.ToUpper("Get"),
		"/volume/{VolumeID}",
		controllers.VolumeVolumeIDGet,
	},

	Route{
		"VolumeVolumeIDPatch",
		strings.ToUpper("Patch"),
		"/volume/{VolumeID}",
		controllers.VolumeVolumeIDPatch,
	},
}
