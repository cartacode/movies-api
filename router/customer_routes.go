package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var customerRoutes = Routes{

	Route{
		"CustomerListAddItem",
		strings.ToUpper("post"),
		"/v1/customer/preferences/{list}",
		validateTokenMiddleware(controllers.CustomerListAddItem),
	},

	Route{
		"CustomerWishlistDeleteItem",
		strings.ToUpper("delete"),
		"/v1/customer/preferences/{list}",
		validateTokenMiddleware(controllers.CustomerWishlistDeleteItem),
	},

	Route{
		"CustomerProfileGet",
		strings.ToUpper("get"),
		"/v1/customer/profile",
		validateTokenMiddleware(controllers.CustomerProfileGet),
	},
}
