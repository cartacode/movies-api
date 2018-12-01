package customer

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
)

// Routes ..
var Routes = requests.Routes{

	requests.Route{
		Name:        "CustomerListAddItem",
		Method:      "POST",
		Pattern:     "/v1/customer/preferences/{list}",
		HandlerFunc: security.ValidateTokenMiddleware(ListAddItem),
	},

	requests.Route{
		Name:        "WishlistDeleteItem",
		Method:      "DELETE",
		Pattern:     "/v1/customer/preferences/{list}",
		HandlerFunc: security.ValidateTokenMiddleware(DeleteItem),
	},

	requests.Route{
		Name:        "ProfileGet",
		Method:      "GET",
		Pattern:     "/v1/customer/profile",
		HandlerFunc: security.ValidateTokenMiddleware(ProfileGet),
	},

	requests.Route{
		Name:        "GetRecentlyWatched",
		Method:      "GET",
		Pattern:     "/v1/customer/recent",
		HandlerFunc: security.ValidateTokenMiddleware(GetRecentlyWatched),
	},

	requests.Route{
		Name:        "SetRecentlyWatched",
		Method:      "POST",
		Pattern:     "/v1/customer/recent",
		HandlerFunc: security.ValidateTokenMiddleware(SetRecentlyWatched),
	},
}
