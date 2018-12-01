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
	"fmt"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/stringops"

	"github.com/VuliTv/go-movie-api/app/customer"
	"github.com/VuliTv/go-movie-api/libs/requests"
	AuthorizeCIM "gopkg.in/hunterlong/authorizecim.v1"
	"gopkg.in/mgo.v2/bson"
)

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
	authUser, authErr := requests.GetAuthUser(r)

	if authErr != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, authErr.Error()))
		return
	}

	var userCustomerCredit CreateCustomerProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&userCustomerCredit); err != nil {
		log.Warn("Request Body parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

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

	authorizeCustomer, acErr := AuthorizeCIM.CreateProfile(data)
	if acErr != nil {
		log.Error("JSON parse error: ", acErr.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, acErr.Error())
		return
	}

	retval, rErr := json.Marshal(authorizeCustomer)
	if rErr != nil {
		log.Error("Error: ", rErr.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, rErr.Error())
		return
	}

	if authorizeCustomer.Messages.ResultCode == "Error" {
		log.Error("Error: ", authorizeCustomer.Messages.Message[0].Text)
		requests.ReturnAPIError(w, http.StatusBadRequest, authorizeCustomer.Messages.Message[0].Text)
		return
	}

	// Find doc
	user := &customer.Model{}
	if err := mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	user.Credit.InfoStored = true
	user.Credit.ProfileID = authorizeCustomer.CustomerProfileID
	user.Credit.PaymentID = authorizeCustomer.CustomerPaymentProfileIDList[0]
	if err := mongoHandler.Collection(collection).Save(user); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	requests.ReturnAPIOK(w, retval)
}

/* CustomerPaymentAdd ..
// adds a payment option to Authorize.Net user profile
func CustomerPaymentAdd(w http.ResponseWriter, r *http.Request) {
	// Get auth user information
	authUser, authErr := requests.GetAuthUser(r)

	if authErr != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, authErr.Error()))
		return
	}
	var userProfileRequest CustomerPaymentProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&userProfileRequest); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var err error
	var paymentProfileResponse *AuthorizeCIM.CustomerPaymentProfileResponse

	paymentProfile := AuthorizeCIM.CustomerPaymentProfile{
		CustomerProfileID: userProfileRequest.ID,
		PaymentProfile: AuthorizeCIM.PaymentProfile{
			BillTo: userProfileRequest.BillTo,
			Payment: &AuthorizeCIM.Payment{
				CreditCard: userProfileRequest.CreditCard,
			},
		},
	}

	if paymentProfileResponse, err = paymentProfile.Add(); err != nil {
		log.Error("Failed to add Payment profile: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Find doc
	user := &customer.Model{}
	if err := mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	user.Credit.InfoStored = true
	user.Credit.ProfileID = paymentProfileResponse.CustomerProfileId
	user.Credit.PaymentID = paymentProfileResponse.CustomerPaymentProfileID

	if err := mongoHandler.Collection(collection).Save(user); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	var retval []byte
	if retval, err = json.Marshal(paymentProfileResponse); err != nil {
		log.Error("Error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, retval)
}*/

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

	var customerPaymentRequest CustomerPaymentUpdateRequest
	// Get auth user information
	var authUser, aErr = requests.GetAuthUser(r)

	if aErr != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, aErr.Error()))
		return
	}

	// Find doc
	user := &customer.Model{}
	if err := mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&customerPaymentRequest); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if customerPaymentRequest.CC == nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, "Credit Card Information Missing")
		return
	}

	var err error
	var res *AuthorizeCIM.MessagesResponse

	authPaymentProfile := &AuthorizeCIM.PaymentProfiles{
		CustomerType: Individual,
	}
	if len(customerPaymentRequest.CC.CardNumber) != 0 {
		cc := &AuthorizeCIM.CreditCard{CardNumber: customerPaymentRequest.CC.CardNumber, ExpirationDate: customerPaymentRequest.CC.ExpirationDate, CardCode: customerPaymentRequest.CC.CardCode}
		authPaymentProfile.Payment = AuthorizeCIM.Payment{CreditCard: *cc}
	}

	profile := AuthorizeCIM.Profile{
		CustomerProfileId: user.Credit.ProfileID,
		PaymentProfileId:  user.Credit.PaymentID,
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

// CustomerPurchaseItem --
func CustomerPurchaseItem(w http.ResponseWriter, r *http.Request) {

	var customerPurchaseRequest CustomerPurchaseRequest
	// Get auth user information
	var authUser, aErr = requests.GetAuthUser(r)

	if aErr != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, aErr.Error()))
		return
	}

	// Find doc for the user
	user := &customer.Model{}
	if err := mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Decode the purchase request
	if err := json.NewDecoder(r.Body).Decode(&customerPurchaseRequest); err != nil {
		log.Error("JSON parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate
	allowedCollections := []string{"volume", "movie", "scene"}
	if !stringops.StringInSlice(customerPurchaseRequest.Collection, allowedCollections) {
		badText := fmt.Sprintf("collection must be on of [%s]", allowedCollections)
		requests.ReturnAPIError(w, http.StatusBadRequest, badText)
		return
	}
	// Find the item to be purchased
	purchase := &MediaPurchase{}
	if err := mongoHandler.Collection(customerPurchaseRequest.Collection).FindById(customerPurchaseRequest.ID, &purchase); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// See if we own it already
	badText := "item already purchased"
	// Check which collection it is
	switch customerPurchaseRequest.Collection {
	case "movie":
		if existsInCollection(purchase.Id, user.Purchased.Movies) {
			log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, badText))
			return
		}
		user.Purchased.Movies = append(user.Purchased.Movies, &purchase.Id)

	case "volume":
		if existsInCollection(purchase.Id, user.Purchased.Volumes) {
			log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, badText))
			return
		}
		user.Purchased.Volumes = append(user.Purchased.Volumes, &purchase.Id)
	case "scene":
		if existsInCollection(purchase.Id, user.Purchased.Scenes) {
			log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, badText))
			return
		}
		user.Purchased.Scenes = append(user.Purchased.Scenes, &purchase.Id)
	}

	authorizeCustomer := AuthorizeCIM.Customer{
		ID:        user.Credit.ProfileID,
		PaymentID: user.Credit.PaymentID,
	}

	newTransaction := AuthorizeCIM.NewTransaction{
		Amount: fmt.Sprintf("%f", purchase.Price),
	}

	response, err := newTransaction.ChargeProfile(authorizeCustomer)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Save the info for the user
	if err := mongoHandler.Collection(collection).Save(user); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	requests.ReturnAPIOK(w, rs)
}

func existsInCollection(id bson.ObjectId, list []*bson.ObjectId) bool {

	var slice []string
	for _, oID := range list {
		slice = append(slice, oID.Hex())
	}
	if stringops.StringInSlice(id.Hex(), slice) {
		return true
	}
	return false
}
