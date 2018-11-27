package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var paymentRoutes = Routes{

	Route{
		"CustomerGetProfile",
		strings.ToUpper("Get"),
		"/v1/customer/payment",
		validateTokenMiddleware(controllers.CustomerGetPaymentProfile),
	},

	Route{
		"CustomerCreateProfile",
		strings.ToUpper("Post"),
		"/v1/customer/payment/create",
		validateTokenMiddleware(controllers.CustomerCreatePaymentProfile),
	},

	Route{
		"CustomerPaymentAdd",
		strings.ToUpper("Post"),
		"/v1/customer/payment",
		validateTokenMiddleware(controllers.CustomerPaymentAdd),
	},

	Route{
		"CustomerPaymentDelete",
		strings.ToUpper("Delete"),
		"/v1/customer/payment",
		validateTokenMiddleware(controllers.CustomerPaymentDelete),
	},

	Route{
		"CustomerPaymentUpdate",
		strings.ToUpper("Patch"),
		"/v1/customer/payment",
		validateTokenMiddleware(controllers.CustomerPaymentUpdate),
	},
}
