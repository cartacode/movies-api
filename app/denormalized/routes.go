package denormalized

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
)

// Routes --
var Routes = requests.Routes{

	requests.Route{
		Name:        "DenormalizedScenes",
		Method:      "GET",
		Pattern:     "/v1/denormalized/scene",
		HandlerFunc: Scenes,
	},

	requests.Route{
		Name:        "DenormalizedMovies",
		Method:      "GET",
		Pattern:     "/v1/denormalized/movie",
		HandlerFunc: Movies,
	},

	requests.Route{
		Name:        "DenormalizedStars",
		Method:      "GET",
		Pattern:     "/v1/denormalized/star",
		HandlerFunc: Stars,
	},
	requests.Route{
		Name:        "DenormalizedVolumes",
		Method:      "GET",
		Pattern:     "/v1/denormalized/volume",
		HandlerFunc: Volumes,
	},
}
