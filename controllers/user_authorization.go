package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/VuliTv/go-movie-api/app/customer"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var tokenExpire = time.Hour * 24

func validPasswordStrength(password string, email string) bool {
	pass := security.NewPasswordChecker(password)
	pass.ProcessPassword()

	log.Debugw("password rating", "user", email, "rating", pass.ComplexityRating())
	return true
}

// CustomerLogin --
func CustomerLogin(w http.ResponseWriter, req *http.Request) {

	collection := "customer"
	var user customer.Model
	if err = json.NewDecoder(req.Body).Decode(&user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Find a customer from this auth attempt
	log.Debug("looking for existing customer")
	existing := &customer.Model{}
	if err = connection.Collection(collection).FindOne(bson.M{"email": user.Email}, existing); err != nil {

		log.Warn(requests.ReturnAPIError(w, http.StatusUnauthorized, "no such user"))
		return
	}

	// Check for Lockout
	if existing.AuthLocked() {
		log.Warn(requests.ReturnAPIError(w, http.StatusUnauthorized, "account locked"))
		return
	}

	// Check password hash
	if err = bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(user.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		log.Debugw("passwords do not match", "user", user.Email)
		log.Warn(requests.ReturnAPIError(w, http.StatusUnauthorized, "unable to authenticate"))

		// Log the bad attempt
		existing.AuthBadAttempt()

		return
	}

	authUser := customer.AuthUser{Email: existing.Email, ObjectID: existing.GetId().Hex(), Admin: existing.Admin}

	// Set token expire time
	expiresAt := time.Now().Add(tokenExpire).Unix()

	// extend admin time to 7 days
	if existing.Admin {
		expiresAt = time.Now().Add(tokenExpire * 3650).Unix()
	}
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &customer.AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		AuthUser: authUser,
	}

	// Signing string with our secret
	tokenString, err := token.SignedString([]byte(customer.JWTSecret))

	if err != nil {
		log.Debug(err)
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, "unable to tokenize"))
		return
	}

	log.Debugw("setting redis token",
		"key", existing.GetId().Hex(),
		"value", tokenString,
		"expire", tokenExpire,
	)
	if err = rDB.Set(existing.GetId().Hex(), tokenString, tokenExpire).Err(); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
	}

	// reset auth on good attempt
	existing.AuthReset()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}

// CustomerSignup --
func CustomerSignup(w http.ResponseWriter, r *http.Request) {
	collection := "customer"
	// Parse and decode the request body into a new `Customer` instance
	user := &customer.Model{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	existing := &customer.Model{}
	if err = connection.Collection(collection).FindOne(bson.M{"email": user.Email}, existing); err != nil {
		log.Debug(err.Error())
	}

	// if requests.ReturnOnError(w, err) {
	// return
	// }
	if existing.Email == user.Email {
		log.Infow(requests.ReturnAPIError(w, http.StatusBadRequest, "user exists"), "user", user.Email)
		return
	}

	// Check password strength
	if !validPasswordStrength(user.Password, user.Email) {
		log.Info(requests.ReturnAPIError(w, http.StatusBadRequest, "does not meet complexity requirements"))
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		log.Info(requests.ReturnAPIError(w, http.StatusInternalServerError, "something went wrong"), "error", err.Error())
		return

	}
	user.Password = string(hashedPassword)

	// Do new customer setup stuff here
	user.Active = true
	user.Admin = false

	// Allow admin from vuli
	if strings.Contains(user.Email, "vuli.tv") {
		user.Admin = true
	}

	// Next, insert the username, along with the hashed password into the database
	if err = connection.Collection(collection).Save(user); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	if vErr, ok := err.(*bongo.ValidationError); ok {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, vErr.Errors[0].Error()))
		return
	}
	response := requests.JSONSuccessResponse{Message: "Success", Identifier: user.GetId().Hex()}

	js, err := json.Marshal(response)
	requests.ReturnAPIOK(w, js)
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back

}

// CustomerUnlockRequest --
func CustomerUnlockRequest(w http.ResponseWriter, r *http.Request) {
	customerUnlockReq := &customer.ModelUnlockRequest{}
	if err := json.NewDecoder(r.Body).Decode(customerUnlockReq); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	customer := &customer.Model{}
	query := make(map[string]interface{})
	query["email"] = customerUnlockReq.Email
	if err := connection.Collection("customer").FindOne(query, &customer); err != nil {
		log.Error(err)
	}

	hash := RandStringRunes(32)
	if err = rDB.Set(hash, customer.GetId().Hex(), time.Hour*1).Err(); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
	}

	body := fmt.Sprintf("<html><body><a href='https://api-stage.vuli.tv/v1/authorize/reset/%s'>Click here to unlock your account</a></body></html>", string(hash[:]))
	if err := sesConn.GenerateAndSendEmail("dev@vuli.tv", customer.Email, "Vuli: Unlock Your Acccount", body); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	log.Infow("account reset password sent", "email", customer.Email)

	response := requests.JSONSuccessResponse{Message: "success", Identifier: customer.GetId().Hex(), Extra: customer.Email}

	js, err := json.Marshal(response)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}

// CustomerUnlock --
func CustomerUnlock(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hash := params["hash"]
	val, err := rDB.Get(hash).Result()

	if err != nil {
		log.Error(err)
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	customer := &customer.Model{}
	connection.Collection("customer").FindById(bson.ObjectIdHex(val), &customer)

	customer.AuthReset()

	response := requests.JSONSuccessResponse{Message: "success", Identifier: customer.GetId().Hex(), Extra: "auth reset"}

	js, err := json.Marshal(response)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}
