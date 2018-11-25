/*
 * Vuli API
 *
 * Vuli Customer Payment Methods API
 *
 * API version: 3
 */

package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/api"
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/gorilla/mux"
)

var authName = envhelp.GetEnv("AUTHORIZE_ID", "65Vv2fYQ")
var authKey = envhelp.GetEnv("AUTHORIZE_TRANSACTION_KEY", "4Ld7LUr432q6e7Uz")

var (
	auth models.AuthorizeMerchant = models.AuthorizeMerchant{Name: authName, TransactionKey: authKey}
)

// CustomerCreateProfile -- creates a user profile with Authorize.net
func CustomerCreateProfile(w http.ResponseWriter, r *http.Request) {

	var retval models.CreateCustomerProfileResponse

	var user models.CreateCustomerProfileRequest
	if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Error("Request Body parse error: ", err.Error())
		api.Respond(w, r, retval, err)
		return
	}

	authPaymentProfile := &models.AuthorizePaymentProfile{
		CustomerType: models.Individual,
	}
	if len(user.CC.CardNumber) != 0 {
		cc := &models.CreditCard{CardNumber: user.CC.CardNumber, ExpirationDate: user.CC.ExpirationDate, CardCode: user.CC.CardCode}
		authPaymentProfile.Payment = &models.AuthorizePayment{CreditCard: *cc}
	}

	data := models.CreateCustomerProfile{
		CreateCustomerProfileRequest: models.AuthorizeProfileRequest{
			Profile: models.AuthorizeCustomer{
				MerchantCustomerId: user.ID,
				Description:        user.Description,
				Email:              user.Email,
				PaymentProfiles:    *authPaymentProfile,
			},
			MerchantAuthentication: auth,
			ValidationMode:         models.ValidationModeTest,
		},
	}

	var req []byte
	if req, err = json.Marshal(data); err != nil {
		log.Error("JSON Parse Error: ", err.Error())
	}
	// TODO - endpoint from config
	res, _ := api.Post("https://apitest.authorize.net/xml/v1/request.api", req)
	res = bytes.TrimPrefix(res, []byte("\xef\xbb\xbf"))

	//****** TODO - we could also not parse the response and send it as is
	if err := json.Unmarshal(res, &retval); err != nil {
		log.Error("JSON Parse Error: ", err.Error())
	}
	if retval.Messages.ResultCode != models.ResponseOK {
		err = errors.New(retval.Messages.Message[0].Text)
	}
	//****** TODO - we could also not parse the response and send it as is

	api.Respond(w, r, retval, err)
	// requests.ReturnAPIOK(w, res)
}

// GetCustomerProfile -- fetches a user profile from Authorize.net
func GetCustomerProfile(w http.ResponseWriter, r *http.Request) {

	var retval models.CustomerProfileInformationResponse

	params := mux.Vars(r)
	objectID := params["userID"]

	var user models.CustomerProfileInformationRequest
	user.ID = objectID

	data := models.GetCustomerProfile{
		GetCustomerProfileRequest: models.GetCustomerProfileRequest{
			MerchantAuthentication: auth,
			CustomerProfileId:      user.ID,
			IncludeIssuerInfo:      true,
		},
	}

	var req []byte
	if req, err = json.Marshal(data); err != nil {
		log.Error("JSON Parse Error: ", err.Error())
	}
	// TODO - endpoint from config
	res, _ := api.Post("https://apitest.authorize.net/xml/v1/request.api", req)
	res = bytes.TrimPrefix(res, []byte("\xef\xbb\xbf"))

	if err := json.Unmarshal(res, &retval); err != nil {
		log.Error("JSON Parse Error: ", err.Error())
	}
	if retval.Messages.ResultCode != models.ResponseOK {
		err = errors.New(retval.Messages.Message[0].Text)
	}

	api.Respond(w, r, retval, err)
}

// CustomerPaymentAdd --
func CustomerPaymentAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerPaymentDelete --
func CustomerPaymentDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerPaymentUpdate --
func CustomerPaymentUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
