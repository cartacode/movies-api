package denormalized

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
)

// Routes --
var Routes = requests.Routes{

	requests.Route{
		Name:        "DenormalizedScenes",
		Method:      "GET",
		Pattern:     "/v1/denormalized/scenes",
		HandlerFunc: Scenes,
	},
}
