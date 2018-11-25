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

type messages struct {
	ResultCode string `json:"resultCode"`
	Message    []struct {
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"message"`
}

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
	CustomerType             CustomerType      `json:"customerType,omitempty"`
	Payment                  *AuthorizePayment `json:"payment,omitempty"`
	CustomerPaymentProfileId string            `json:"customerPaymentProfileId,omitempty"`
	BillTo                   *BillTo           `json:"billTo,omitempty"`
}

type CreditCard struct {
	CardNumber     string `json:"cardNumber,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	CardCode       string `json:"cardCode,omitempty"`
}

type AuthorizePayment struct {
	CreditCard CreditCard `json:"creditCard,omitempty"`
}

type BillTo struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
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
	Messages                      messages      `json:"messages"`
}

/******************************************** Customer Profile Fetch */
type CustomerProfileInformationRequest struct {
	ID string `json:"id"`
}

type GetCustomerProfileRequest struct {
	MerchantAuthentication AuthorizeMerchant `json:"merchantAuthentication"`
	CustomerProfileId      string            `json:"customerProfileId"`
	IncludeIssuerInfo      bool              `json:"includeIssuerInfo" default:"true"`
}

type GetCustomerProfile struct {
	GetCustomerProfileRequest GetCustomerProfileRequest `json:"getCustomerProfileRequest"`
}

// CustomerProfileInformationResponse - parse customer profile informatoin
// response from Authorize.net
type CustomerProfileInformationResponse struct {
	Profile struct {
		MerchantCustomerId string `json:"merchantCustomerId"`

		// Profile description for this customer
		Description string `json:"description,omitempty"`

		// Unique email for this customer
		Email string `json:"email"`

		// Customer profile type - regular / guest
		ProfileType string `json:"profileType"`

		// Unique email for this customer
		PaymentProfiles []AuthorizePaymentProfile `json:"paymentProfiles"`
	} `json:"profile"`
	Messages messages `json:"messages"`
}
