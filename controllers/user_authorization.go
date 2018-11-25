package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/security"
	"github.com/VuliTv/go-movie-api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
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

// Login --
func Login(w http.ResponseWriter, req *http.Request) {

	collection := "customer"
	var customer models.Customer
	if err = json.NewDecoder(req.Body).Decode(&customer); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Find a customer from this auth attempt
	log.Debug("looking for existing customer")
	existing := &models.Customer{}
	if err = connection.Collection(collection).FindOne(bson.M{"email": customer.Email}, &existing); err != nil {

		log.Warn(requests.ReturnAPIError(w, http.StatusUnauthorized, "no such user"))
		return
	}

	// Check password hash
	if err = bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(customer.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		log.Debugw("passwords do not match", "user", customer.Email)
		log.Warn(requests.ReturnAPIError(w, http.StatusUnauthorized, "unable to authenticate"))
		return
	}

	authUser := models.AuthUser{Email: existing.Email, ObjectID: existing.GetId().Hex(), Admin: existing.Admin}

	// Set token expire time
	expiresAt := time.Now().Add(tokenExpire).Unix()

	// extend admin time to 7 days
	if existing.Admin {
		expiresAt = time.Now().Add(tokenExpire * 7).Unix()
	}
	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &models.AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		AuthUser: authUser,
	}

	// Signing string with our secret
	tokenString, err := token.SignedString([]byte(models.JWTSecret))

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}

// Signup --
func Signup(w http.ResponseWriter, r *http.Request) {
	collection := "customer"
	// Parse and decode the request body into a new `Customer` instance
	customer := &models.Customer{}
	if err := json.NewDecoder(r.Body).Decode(customer); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	existing := &models.Customer{}
	if err = connection.Collection(collection).FindOne(bson.M{"email": customer.Email}, existing); err != nil {
		log.Debug(err.Error())
	}

	// if requests.ReturnOnError(w, err) {
	// return
	// }
	if existing.Email == customer.Email {
		log.Infow(requests.ReturnAPIError(w, http.StatusBadRequest, "user exists"), "user", customer.Email)
		return
	}

	// Check password strength
	if !validPasswordStrength(customer.Password, customer.Email) {
		log.Info(requests.ReturnAPIError(w, http.StatusBadRequest, "does not meet complexity requirements"))
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 10)

	if err != nil {
		log.Info(requests.ReturnAPIError(w, http.StatusInternalServerError, "something went wrong"), "error", err.Error())
		return

	}
	customer.Password = string(hashedPassword)

	// Do new customer setup stuff here
	customer.Active = true
	customer.Admin = false

	// Allow admin from vuli
	if strings.Contains(customer.Email, "vuli.tv") {
		customer.Admin = true
	}

	// Next, insert the username, along with the hashed password into the database
	err = connection.Collection(collection).Save(customer)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, http.StatusBadRequest, vErr.Errors[0].Error())
		return
	}
	response := requests.JSONSuccessResponse{Message: "Success", Identifier: customer.GetId().Hex()}

	js, err := json.Marshal(response)
	requests.ReturnAPIOK(w, js)
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back

}
