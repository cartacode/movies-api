/*
 * Vuli API
 *
 * Vuli Authorize.net Model
 *
 * API version: 3
 */

package models

type CustomerType string
type ValidationMode string
type ResponseCode string

const (
	Individual CustomerType = "individual"
	Business   CustomerType = "business"
)
const (
	ValidationModeNone    ValidationMode = "none"
	ValidationModeTest    ValidationMode = "testMode"
	ValidationModeLive    ValidationMode = "liveMode"
	ValidationModeOldLive ValidationMode = "oldLiveMode"
)
const (
	ResponseOK    = "OK"
	ResponseError = "Error"
)

type AuthorizeMerchant struct {
	Name           string `json:"name"`
	TransactionKey string `json:"transactionKey"`
}

// Customer Document
// A profile for customer on Authorize.net
type AuthorizeCustomer struct {
	MerchantCustomerId string `json:"merchantCustomerId"`

	// Profile description for this customer
	Description string `json:"description,omitempty"`

	// Unique email for this customer
	Email string `json:"email"`

	// Unique email for this customer
	PaymentProfiles AuthorizePaymentProfile `json:"paymentProfiles"`
}

type AuthorizePaymentProfile struct {
	CustomerType CustomerType      `json:"customerType"`
	Payment      *AuthorizePayment `json:"payment,omitempty"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}

type AuthorizePayment struct {
	CreditCard CreditCard `json:"creditCard,omitempty"`
}

type AuthorizeProfileRequest struct {
	MerchantAuthentication AuthorizeMerchant `json:"merchantAuthentication"`
	Profile                AuthorizeCustomer `json:"profile"`
	ValidationMode         ValidationMode    `json:"validationMode" default:"testMode"`
}

type CreateCustomerProfile struct {
	CreateCustomerProfileRequest AuthorizeProfileRequest `json:"createCustomerProfileRequest"`
}

// CreateCustomerProfileRequest - request to create user profile
// e.g. from front-end
type CreateCustomerProfileRequest struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Description string     `json:"description,omitempty"`
	CC          CreditCard `json:"creditCard"`
}

// CreateCustomerProfileRequest - parse response from Authorize.net
type CreateCustomerProfileResponse struct {
	CustomerProfileId             string        `json:"customerProfileId"`
	CustomerPaymentProfileIdList  []interface{} `json:"customerPaymentProfileIdList"`
	CustomerShippingAddressIdList []interface{} `json:"customerShippingAddressIdList"`
	ValidationDirectResponseList  []string      `json:"validationDirectResponseList"`
	Messages                      struct {
		ResultCode string `json:"resultCode"`
		Message    []struct {
			Code string `json:"code"`
			Text string `json:"text"`
		} `json:"message"`
	} `json:"messages"`
}
