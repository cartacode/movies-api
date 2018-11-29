/*
 * Vuli API
 *
 * Vuli Customer Payment Methods API
 *
 * API version: 3
 */

package payments

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/app/customer"
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

	var userInfo *AuthorizeCIM.GetCustomerProfileResponse

	// Find doc
	user := &customer.Model{}
	if err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	authorizeCustomer := AuthorizeCIM.Customer{
		ID: user.Credit.ProfileID,
	}
	if userInfo, err = authorizeCustomer.Info(); err != nil {
		log.Error("Customer Information fetch error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(userInfo); err != nil {
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

	var userCustomerCredit CreateCustomerProfileRequest
	if err = json.NewDecoder(r.Body).Decode(&userCustomerCredit); err != nil {
		log.Warn("Request Body parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var authorizeCustomer *AuthorizeCIM.CustomProfileResponse

	authPaymentProfile := &AuthorizeCIM.PaymentProfiles{
		CustomerType: Individual,
	}
	if len(userCustomerCredit.CC.CardNumber) != 0 {
		cc := &AuthorizeCIM.CreditCard{CardNumber: userCustomerCredit.CC.CardNumber, ExpirationDate: userCustomerCredit.CC.ExpirationDate, CardCode: userCustomerCredit.CC.CardCode}
		authPaymentProfile.Payment = AuthorizeCIM.Payment{CreditCard: *cc}
	}
	// API not supporting
	// authPaymentProfile.BillTo = userCustomerCredit.BillTo

	data := AuthorizeCIM.Profile{
		MerchantCustomerID: authUser.Email,
		Description:        userCustomerCredit.Description,
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
	user := &customer.Model{}
	if err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	user.Credit.InfoStored = true
	user.Credit.ProfileID = authorizeCustomer.CustomerProfileID
	user.Credit.PaymentID = authorizeCustomer.CustomerPaymentProfileIDList[0]
	if err = mongoHandler.Collection(collection).Save(user); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentAdd ..
// adds a payment option to Authorize.Net user profile
func CustomerPaymentAdd(w http.ResponseWriter, r *http.Request) {

	var userProfileRequest CustomerPaymentProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&userProfileRequest); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var err error
	var user *AuthorizeCIM.CustomerPaymentProfileResponse

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	paymentProfile := AuthorizeCIM.CustomerPaymentProfile{
		CustomerProfileID: userProfileRequest.ID,
		PaymentProfile: AuthorizeCIM.PaymentProfile{
			BillTo: userProfileRequest.BillTo,
			Payment: &AuthorizeCIM.Payment{
				CreditCard: userProfileRequest.CreditCard,
			},
		},
	}

	if user, err = paymentProfile.Add(); err != nil {
		log.Error("Failed to add Payment profile: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(user); err != nil {
		log.Error("Error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}

// CustomerPaymentDelete -- deletes a payment option from
// Authorize.Net user profile
func CustomerPaymentDelete(w http.ResponseWriter, r *http.Request) {

	// Get auth user information
	var authUser, err = requests.GetAuthUser(r)

	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Find doc
	user := &customer.Model{}
	if err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	var res *AuthorizeCIM.MessagesResponse

	AuthorizeCIM.SetAPIInfo(authName, authKey, authMode)
	authorizeCustomer := AuthorizeCIM.Customer{
		ID:        user.Credit.ProfileID,
		PaymentID: user.Credit.PaymentID,
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
	user.Credit.PaymentID = ""

	if err = mongoHandler.Collection(collection).Save(user); err != nil {
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
// Authorize.Net user profile
func CustomerPaymentUpdate(w http.ResponseWriter, r *http.Request) {

	var user CustomerPaymentUpdateRequest
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
		CustomerType: Individual,
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
