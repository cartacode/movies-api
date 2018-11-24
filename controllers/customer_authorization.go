package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/go-bongo/bongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// Login --
func Login(w http.ResponseWriter, req *http.Request) {

	collection := "customer"
	var customer models.Customer
	_ = json.NewDecoder(req.Body).Decode(&customer)

	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	// Find doc
	existing := &models.Customer{}
	err = connection.Collection(collection).FindOne(bson.M{"email": customer.Email}, &existing)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		requests.ReturnAPIError(w, fmt.Errorf("no such user"))
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(existing.Password), []byte(customer.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authUser := models.AuthUser{Email: existing.Email, ObjectID: existing.GetId().Hex(), Admin: existing.Admin}

	token.Claims = &models.AuthTokenClaim{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		AuthUser: authUser,
	}

	tokenString, error := token.SignedString([]byte(models.JWTSecret))
	if error != nil {
		fmt.Println(error)
	}
	err := rDB.Set(existing.GetId().Hex(), tokenString, 0).Err()
	if err != nil {
		panic(err)
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
	err := json.NewDecoder(r.Body).Decode(customer)
	if requests.ReturnOnError(w, err) {
		return
	}

	existing := &models.Customer{}
	_ = connection.Collection(collection).FindOne(bson.M{"email": customer.Email}, existing)

	// if requests.ReturnOnError(w, err) {
	// return
	// }
	if existing.Email == customer.Email {
		w.WriteHeader(http.StatusUnauthorized)
		requests.ReturnAPIError(w, fmt.Errorf("user exists with this email"))
		return
	}
	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 8)

	customer.Password = string(hashedPassword)

	// Do new customer setup stuff here
	customer.Active = true

	// Next, insert the username, along with the hashed password into the database
	err = connection.Collection(collection).Save(customer)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, vErr.Errors[0])
		return
	}
	response := requests.JSONSuccessResponse{Message: "Success", Identifier: customer.GetId().Hex()}

	js, err := json.Marshal(response)
	requests.ReturnAPIOK(w, js)
	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back

}
