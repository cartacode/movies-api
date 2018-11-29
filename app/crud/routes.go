package crud

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
)

// Routes --
var Routes = requests.Routes{
	requests.Route{
		Name:        "GenericCrudGet",
		Method:      "GET",
		Pattern:     "/v1/collection/{collection}",
		HandlerFunc: security.ValidateTokenMiddleware(GenericCrudGet),
	},

	requests.Route{
		Name:        "GenericCrudIDGet",
		Method:      "GET",
		Pattern:     "/v1/collection/{collection}/{objectID}",
		HandlerFunc: security.ValidateTokenMiddleware(GenericCrudIDGet),
	},

	requests.Route{
		Name:        "GenericCrudIDDelete",
		Method:      "DELETE",
		Pattern:     "/v1/collection/{collection}/{objectID}",
		HandlerFunc: security.ValidateTokenMiddleware(GenericCrudIDDelete),
	},

	requests.Route{
		Name:        "GenericCrudPost",
		Method:      "POST",
		Pattern:     "/v1/collection/{collection}",
		HandlerFunc: security.ValidateTokenMiddleware(GenericCrudPost),
	},

	requests.Route{
		Name:        "GenericCrudIDPatch",
		Method:      "PATCH",
		Pattern:     "/v1/collection/{collection}/{objectID}",
		HandlerFunc: security.ValidateTokenMiddleware(GenericCrudIDPatch),
	},
}
