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

	"github.com/VuliTv/go-movie-api/app/customer"
	"github.com/VuliTv/go-movie-api/app/payments"
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/VuliTv/go-movie-api/libs/requests"
	AuthorizeCIM "gopkg.in/hunterlong/authorizecim.v1"
	"gopkg.in/mgo.v2/bson"
)

var authName = envhelp.GetEnv("AUTHORIZE_ID", "65Vv2fYQ")
var authKey = envhelp.GetEnv("AUTHORIZE_TRANSACTION_KEY", "4Ld7LUr432q6e7Uz")
var authMode = "test"

// CustomerGetPaymentProfile ..
// fetches a customer profile from Authorize.net
func CustomerGetPaymentProfile(w http.ResponseWriter, r *http.Request) {

	// Get auth user information
	var authUser, err = requests.GetAuthUser(r)

	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	var customerInfo *AuthorizeCIM.GetCustomerProfileResponse

	// Find doc
	customer := &customer.Model{}
	if err = connection.Collection("customer").FindById(bson.ObjectIdHex(authUser.ObjectID), &customer); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	authorizeCustomer := AuthorizeCIM.Customer{
		ID: customer.Credit.ProfileID,
	}
	if customerInfo, err = authorizeCustomer.Info(); err != nil {
		log.Error("Customer Information fetch error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(customerInfo); err != nil {
		log.Error("Error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerCreatePaymentProfile ..
// creates a user profile with Authorize.net
func CustomerCreatePaymentProfile(w http.ResponseWriter, r *http.Request) {
	// Get auth user information
	var authUser, err = requests.GetAuthUser(r)

	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	var user payments.CreateCustomerProfileRequest
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Warn("Request Body parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var authorizeCustomer *AuthorizeCIM.CustomProfileResponse

	authPaymentProfile := &AuthorizeCIM.PaymentProfiles{
		CustomerType: payments.Individual,
	}
	if len(user.CC.CardNumber) != 0 {
		cc := &AuthorizeCIM.CreditCard{CardNumber: user.CC.CardNumber, ExpirationDate: user.CC.ExpirationDate, CardCode: user.CC.CardCode}
		authPaymentProfile.Payment = AuthorizeCIM.Payment{CreditCard: *cc}
	}
	// API not supporting
	// authPaymentProfile.BillTo = user.BillTo

	data := AuthorizeCIM.Profile{
		MerchantCustomerID: authUser.Email,
		Description:        user.Description,
		Email:              authUser.Email,
		PaymentProfiles:    authPaymentProfile,
	}

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	if authorizeCustomer, err = AuthorizeCIM.CreateProfile(data); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(authorizeCustomer); err != nil {
		log.Error("Error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if authorizeCustomer.Messages.ResultCode == "Error" {
		log.Error("Error: ", authorizeCustomer.Messages.Message[0].Text)
		requests.ReturnAPIError(w, http.StatusBadRequest, authorizeCustomer.Messages.Message[0].Text)
		return
	}

	// Find doc
	customer := &customer.Model{}
	if err = connection.Collection("customer").FindById(bson.ObjectIdHex(authUser.ObjectID), &customer); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	customer.Credit.InfoStored = true
	customer.Credit.ProfileID = authorizeCustomer.CustomerProfileID
	customer.Credit.PaymentID = authorizeCustomer.CustomerPaymentProfileIDList[0]
	if err = connection.Collection("customer").Save(customer); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentAdd ..
// adds a payment option to Authorize.Net customer profile
func CustomerPaymentAdd(w http.ResponseWriter, r *http.Request) {

	var user payments.CustomerPaymentProfileRequest
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
			BillTo: user.BillTo,
			Payment: &AuthorizeCIM.Payment{
				CreditCard: user.CreditCard,
			},
		},
	}

	if customer, err = paymentProfile.Add(); err != nil {
		log.Error("Failed to add Payment profile: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(customer); err != nil {
		log.Error("Error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentDelete -- deletes a payment option from
// Authorize.Net customer profile
func CustomerPaymentDelete(w http.ResponseWriter, r *http.Request) {

	// Get auth user information
	var authUser, err = requests.GetAuthUser(r)

	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Find doc
	customer := &customer.Model{}
	if err = connection.Collection("customer").FindById(bson.ObjectIdHex(authUser.ObjectID), &customer); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	var res *AuthorizeCIM.MessagesResponse

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	authorizeCustomer := AuthorizeCIM.Customer{
		ID:        customer.Credit.ProfileID,
		PaymentID: customer.Credit.PaymentID,
	}
	if res, err = authorizeCustomer.DeletePaymentProfile(); err != nil {
		log.Error("Customer Payment delete error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if res.Messages.ResultCode == "Error" {
		log.Error("Customer Payment delete error: ", res.Messages.Message[0].Text)
		requests.ReturnAPIError(w, http.StatusBadRequest, res.Messages.Message[0].Text)
		return
	}
	customer.Credit.PaymentID = ""

	if err = connection.Collection("customer").Save(customer); err != nil {
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

	var user payments.CustomerPaymentUpdateRequest
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
		CustomerType: payments.Individual,
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
		log.Error("Customer Information update error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(res); err != nil {
		log.Error("Error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}
