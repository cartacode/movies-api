package operations

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
)

// Routes --
var Routes = requests.Routes{

	requests.Route{
		Name:        "UploadImage",
		Method:      "POST",
		Pattern:     "/v1/operations/upload/image/{collection}",
		HandlerFunc: UploadImage,
	},

	requests.Route{
		Name:        "Trailer",
		Method:      "POST",
		Pattern:     "/v1/operations/upload/trailer/{collection}",
		HandlerFunc: UploadTrailer,
	},

	requests.Route{
		Name:        "SignedS3Playback",
		Method:      "GET",
		Pattern:     "/v1/operations/play/{collection}/{objectID}",
		HandlerFunc: SignedS3Playback,
	},
}
