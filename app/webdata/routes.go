package webdata

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
)

// Routes ..
var Routes = requests.Routes{

	requests.Route{
		Name:        "DataMovieTray",
		Method:      "GET",
		Pattern:     "/v1/data/movie",
		HandlerFunc: security.ValidateTokenMiddleware(MovieTray),
	},
	requests.Route{
		Name:        "DataSceneTray",
		Method:      "GET",
		Pattern:     "/v1/data/scene",
		HandlerFunc: security.ValidateTokenMiddleware(SceneTray),
	},
	requests.Route{
		Name:        "DataVolumeTray",
		Method:      "GET",
		Pattern:     "/v1/data/volume",
		HandlerFunc: security.ValidateTokenMiddleware(VolumeTray),
	},
	requests.Route{
		Name:        "DataSeriesTray",
		Method:      "GET",
		Pattern:     "/v1/data/series",
		HandlerFunc: security.ValidateTokenMiddleware(SeriesTray),
	},
	requests.Route{
		Name:        "DataStarTray",
		Method:      "GET",
		Pattern:     "/v1/data/star",
		HandlerFunc: security.ValidateTokenMiddleware(StarTray),
	},
}
