/*
 * Vuli API
 *
 * Vuli Authorize.net Model
 * Schema: https://api.authorize.net/xml/v1/schema/AnetApiSchema.xsd
 * API: https://developer.authorize.net/api/reference/index.html
 *
 * API version: 1
 */

package payments

import (
	AuthorizeCIM "gopkg.in/hunterlong/authorizecim.v1"
)

// Individual ..
const (
	Individual string = "individual"
	Business   string = "business"
)

// CreateCustomerProfileRequest ..
type CreateCustomerProfileRequest struct {
	Description string                   `json:"description,omitempty"`
	CC          *AuthorizeCIM.CreditCard `json:"creditCard,omitempty"`
	BillTo      *AuthorizeCIM.BillTo     `json:"billingAddress,omitempty"`
	Shipping    *AuthorizeCIM.Address    `json:"shippingAddress,omitempty"`
}

// CustomerProfileInformationRequest ..
type CustomerProfileInformationRequest struct {
	ID string `json:"id"`
}

// CustomerPaymentProfileRequest ..
type CustomerPaymentProfileRequest struct {
	ID         string                  `json:"id"`
	BillTo     *AuthorizeCIM.BillTo    `json:"billingAddress,omitempty"`
	CreditCard AuthorizeCIM.CreditCard `json:"creditCard"`
}

// CustomerPaymentDeleteRequest ..
type CustomerPaymentDeleteRequest struct {
	ID        string `json:"id"`
	PaymentID string `json:"paymentId"`
}

// CustomerPaymentUpdateRequest ..
type CustomerPaymentUpdateRequest struct {
	ID        string                   `json:"id"`
	PaymentID string                   `json:"paymentId"`
	CC        *AuthorizeCIM.CreditCard `json:"creditCard,omitempty"`
}
