package auth

import "github.com/VuliTv/go-movie-api/libs/requests"

// Routes ..
// Authorization routes
var Routes = requests.Routes{
	requests.Route{
		Name:        "CustomerLogin",
		Method:      "POST",
		Pattern:     "/v1/authorize/login",
		HandlerFunc: CustomerLogin,
	},
	requests.Route{
		Name:        "CustomerSignup",
		Method:      "POST",
		Pattern:     "/v1/authorize/signup",
		HandlerFunc: CustomerSignup,
	},
	requests.Route{
		Name:        "CustomerUnlockRequest",
		Method:      "POST",
		Pattern:     "/v1/authorize/unlock",
		HandlerFunc: CustomerUnlockRequest,
	},

	requests.Route{
		Name:        "CustomerUnlock",
		Method:      "GET",
		Pattern:     "/v1/authorize/reset/{hash}",
		HandlerFunc: CustomerUnlock,
	},
}
