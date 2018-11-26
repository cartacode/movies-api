/*
 * Vuli API
 *
 * Vuli Customer Payment Methods API
 *
 * API version: 3
 */

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/gorilla/mux"
	AuthorizeCIM "gopkg.in/hunterlong/authorizecim.v1"
)

var authName = envhelp.GetEnv("AUTHORIZE_ID", "65Vv2fYQ")
var authKey = envhelp.GetEnv("AUTHORIZE_TRANSACTION_KEY", "4Ld7LUr432q6e7Uz")
var authMode = "test"

// GetCustomerProfile -- fetches a customer profile from Authorize.net
func GetCustomerProfile(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	user := models.CustomerProfileInformationRequest{ID: params["userID"]}

	var err error
	var customerInfo *AuthorizeCIM.GetCustomerProfileResponse

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	customer := AuthorizeCIM.Customer{
		ID: user.ID,
	}
	if customerInfo, err = customer.Info(); err != nil {
		log.Error("Customer Information fetch error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(customerInfo); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerCreateProfile -- creates a user profile with Authorize.net
func CustomerCreateProfile(w http.ResponseWriter, r *http.Request) {

	var user models.CreateCustomerProfileRequest
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error("Request Body parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var err error
	var customer *AuthorizeCIM.CustomProfileResponse

	authPaymentProfile := &AuthorizeCIM.PaymentProfiles{
		CustomerType: models.Individual,
	}
	if len(user.CC.CardNumber) != 0 {
		cc := &AuthorizeCIM.CreditCard{CardNumber: user.CC.CardNumber, ExpirationDate: user.CC.ExpirationDate, CardCode: user.CC.CardCode}
		authPaymentProfile.Payment = AuthorizeCIM.Payment{CreditCard: *cc}
	}

	data := AuthorizeCIM.Profile{
		MerchantCustomerID: user.ID,
		Description:        user.Description,
		Email:              user.Email,
		PaymentProfiles:    authPaymentProfile,
		// Shipping           *Address         `json:"address,omitempty"`
		// PaymentProfile     *PaymentProfile  `json:"paymentProfile,omitempty"`
	}

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	if customer, err = AuthorizeCIM.CreateProfile(data); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(customer); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentAdd -- adds a payment option to
// Authorize.Net customer profile
func CustomerPaymentAdd(w http.ResponseWriter, r *http.Request) {

	var user models.CustomerPaymentProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var err error
	var customer *AuthorizeCIM.CustomerPaymentProfileResponse

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	paymentProfile := AuthorizeCIM.CustomerPaymentProfile{
		CustomerProfileID: user.ID,
		PaymentProfile: AuthorizeCIM.PaymentProfile{
			BillTo: &AuthorizeCIM.BillTo{
				FirstName:   "Amit",
				LastName:    "Mangal",
				Address:     "789 - Sector 2",
				City:        "Bellevue",
				State:       "OH",
				Zip:         "43212",
				Country:     "USA",
				PhoneNumber: "347-111-2222",
			},
			Payment: &AuthorizeCIM.Payment{
				CreditCard: AuthorizeCIM.CreditCard{
					CardNumber:     "4111111111111111",
					ExpirationDate: "2025-12",
				},
			},
		},
	}

	if customer, err = paymentProfile.Add(); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(customer); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentDelete -- deletes a payment option from
// Authorize.Net customer profile
func CustomerPaymentDelete(w http.ResponseWriter, r *http.Request) {

	var user models.CustomerPaymentDeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var err error
	var res *AuthorizeCIM.MessagesResponse

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	customer := AuthorizeCIM.Customer{
		ID:        user.ID,
		PaymentID: user.PaymentID,
	}
	if res, err = customer.DeletePaymentProfile(); err != nil {
		log.Error("Customer Information fetch error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(res); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentUpdate -- updates a payment option in
// Authorize.Net customer profile
func CustomerPaymentUpdate(w http.ResponseWriter, r *http.Request) {

	var user models.CustomerPaymentUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if user.CC == nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, "Credit Card Information Missing")
		return
	}

	var err error
	var res *AuthorizeCIM.MessagesResponse

	authPaymentProfile := &AuthorizeCIM.PaymentProfiles{
		CustomerType: models.Individual,
	}
	if len(user.CC.CardNumber) != 0 {
		cc := &AuthorizeCIM.CreditCard{CardNumber: user.CC.CardNumber, ExpirationDate: user.CC.ExpirationDate, CardCode: user.CC.CardCode}
		authPaymentProfile.Payment = AuthorizeCIM.Payment{CreditCard: *cc}
	}

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	profile := AuthorizeCIM.Profile{
		CustomerProfileId: user.ID,
		PaymentProfileId:  user.PaymentID,
		PaymentProfiles:   authPaymentProfile,
	}

	if res, err = profile.UpdatePaymentProfile(); err != nil {
		log.Error("Customer Information fetch error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(res); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}
