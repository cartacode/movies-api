package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var searchRoutes = Routes{
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
		controllers.MovieSearchGet,
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
		"SceneSearchGet",
		strings.ToUpper("Get"),
		"/search/scene",
		controllers.SceneSearchGet,
	},

	// swagger:operation GET /search/volume search volumeSearchList
	// ---
	// summary: List all found volumes in a pagination response.
	// description: Return all volumes found, paginated
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
		"SearchGet",
		strings.ToUpper("Get"),
		"/search/collection/{collection}",
		controllers.SearchGet,
	},
}
