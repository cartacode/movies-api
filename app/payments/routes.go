package payments

import (
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
)

// Routes --
var Routes = requests.Routes{

	requests.Route{
		Name:        "CustomerGetProfile",
		Method:      "GET",
		Pattern:     "/v1/customer/payment",
		HandlerFunc: security.ValidateTokenMiddleware(CustomerGetPaymentProfile),
	},

	requests.Route{
		Name:        "CustomerCreateProfile",
		Method:      "POST",
		Pattern:     "/v1/customer/payment/create",
		HandlerFunc: security.ValidateTokenMiddleware(CustomerCreatePaymentProfile),
	},

	requests.Route{
		Name:        "CustomerPaymentAdd",
		Method:      "POST",
		Pattern:     "/v1/customer/payment",
		HandlerFunc: security.ValidateTokenMiddleware(CustomerPaymentAdd),
	},

	requests.Route{
		Name:        "CustomerPaymentDelete",
		Method:      "DELETE",
		Pattern:     "/v1/customer/payment",
		HandlerFunc: security.ValidateTokenMiddleware(CustomerPaymentDelete),
	},

	requests.Route{
		Name:        "CustomerPaymentUpdate",
		Method:      "PATCH",
		Pattern:     "/v1/customer/payment",
		HandlerFunc: security.ValidateTokenMiddleware(CustomerPaymentUpdate),
	},
}
