package operations

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
)

// Routes --
var Routes = requests.Routes{

	requests.Route{
		Name:        "UploadImage",
		Method:      "POST",
		Pattern:     "/v1/operations/upload/image/{collection}",
		HandlerFunc: security.ValidateTokenMiddleware(UploadImage),
	},

	requests.Route{
		Name:        "Trailer",
		Method:      "POST",
		Pattern:     "/v1/operations/upload/trailer/{collection}",
		HandlerFunc: security.ValidateTokenMiddleware(UploadTrailer),
	},

	requests.Route{
		Name:        "SignedS3Playback",
		Method:      "GET",
		Pattern:     "/v1/operations/play/{collection}/{objectID}",
		HandlerFunc: SignedS3Playback,
	},
}
