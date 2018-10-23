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

	Route{
		"CategoryCategoryIdDelete",
		strings.ToUpper("Delete"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDDelete,
	},

	Route{
		"CategoryCategoryIDGet",
		strings.ToUpper("Get"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDGet,
	},

	Route{
		"CategorySlugGet",
		strings.ToUpper("Get"),
		"/category/slug/{slug}",
		controllers.CategorySlugGet,
	},

	Route{
		"CategoryCategoryIDPatch",
		strings.ToUpper("Patch"),
		"/category/{categoryID}",
		controllers.CategoryCategoryIDPatch,
	},

	Route{
		"CategoryGet",
		strings.ToUpper("Get"),
		"/category",
		controllers.CategoryGet,
	},

	Route{
		"CategoryPost",
		strings.ToUpper("Post"),
		"/category",
		controllers.CategoryPost,
	},

	Route{
		"CustomerCustomerIDDelete",
		strings.ToUpper("Delete"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDDelete,
	},

	Route{
		"CustomerCustomerIDGet",
		strings.ToUpper("Get"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDGet,
	},

	Route{
		"CustomerCustomerIDPatch",
		strings.ToUpper("Patch"),
		"/customer/{CustomerID}",
		controllers.CustomerCustomerIDPatch,
	},

	Route{
		"CustomerGet",
		strings.ToUpper("Get"),
		"/customer",
		controllers.CustomerGet,
	},

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
