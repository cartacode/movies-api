package router

import (
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
)

var paymentRoutes = Routes{

	// swagger:operation POST /v1/customer/payment/{ObjectId} customer CustomerPaymentAdd
	// ---
	// summary: Add a new payment method for a customer
	// description:
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerPaymentAdd",
		strings.ToUpper("Post"),
		"/v1/customer/payment",
		controllers.CustomerPaymentAdd,
	},

	// swagger:operation DELETE /v1/customer/payment/{ObjectId} customer CustomerPaymentDelete
	// ---
	// summary: Delete a payment method for a customer
	// description:
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerPaymentDelete",
		strings.ToUpper("Delete"),
		"/v1/customer/payment/{ObjectId}",
		controllers.CustomerPaymentDelete,
	},

	// swagger:operation PATCH /v1/customer/payment/{ObjectId} customer CustomerPaymentUpdate
	// ---
	// summary: Update a payment method for a customer
	// description:
	// parameters:
	// - name: ObjectId
	//   in: path
	//   description: MongoDB Scene Document ID
	//   required: true
	//   schema:
	//     "$ref": "#/definitions/ObjectId"
	// responses:
	//   "200":
	//     "$ref": "#/responses/customerResp"
	//   "404":
	//     "$ref": "#/responses/genericJsonError"
	Route{
		"CustomerPaymentUpdate",
		strings.ToUpper("Patch"),
		"/v1/customer/payment/{ObjectId}",
		controllers.CustomerPaymentUpdate,
	},
}
