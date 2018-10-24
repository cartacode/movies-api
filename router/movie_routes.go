package router

import (
	"strings"

	"github.com/VuliTv/api/controllers"
)

var movieRoutes = Routes{
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
}
