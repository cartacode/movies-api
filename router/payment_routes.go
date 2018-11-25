package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var paymentRoutes = Routes{

	Route{
		"CustomerCreateProfile",
		strings.ToUpper("Post"),
		"/v1/customer/profile",
		controllers.CustomerCreateProfile,
	},

	Route{
		"CustomerPaymentAdd",
		strings.ToUpper("Post"),
		"/v1/customer/payment",
		controllers.CustomerPaymentAdd,
	},

	Route{
		"CustomerPaymentDelete",
		strings.ToUpper("Delete"),
		"/v1/customer/payment/{ObjectId}",
		controllers.CustomerPaymentDelete,
	},

	Route{
		"CustomerPaymentUpdate",
		strings.ToUpper("Patch"),
		"/v1/customer/payment/{ObjectId}",
		controllers.CustomerPaymentUpdate,
	},
}
