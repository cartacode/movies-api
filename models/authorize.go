/*
 * Vuli API
 *
 * Vuli Authorize.net Model
 * Schema: https://api.authorize.net/xml/v1/schema/AnetApiSchema.xsd
 * API: https://developer.authorize.net/api/reference/index.html
 *
 * API version: 1
 */

package models

import (
	AuthorizeCIM "github.com/hunterlong/authorizecim"
)

const (
	Individual string = "individual"
	Business   string = "business"
)

type CreateCustomerProfileRequest struct {
	ID          string                   `json:"id"`
	Email       string                   `json:"email"`
	Description string                   `json:"description,omitempty"`
	CC          *AuthorizeCIM.CreditCard `json:"creditCard,omitempty"`
}

type CustomerProfileInformationRequest struct {
	ID string `json:"id"`
}

type CustomerPaymentProfileRequest struct {
	ID string `json:"id"`
}

type CustomerPaymentDeleteRequest struct {
	ID        string
	PaymentID string
}

type CustomerPaymentUpdateRequest struct {
	ID        string                   `json:"id"`
	PaymentID string                   `json:"paymentId"`
	CC        *AuthorizeCIM.CreditCard `json:"creditCard,omitempty"`
}
