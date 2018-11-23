package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VuliTv/go-movie-api/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Authorization --
func Authorization(w http.ResponseWriter, req *http.Request) {
	var user models.User
	_ = json.NewDecoder(req.Body).Decode(&user)

	expiresAt := time.Now().Add(time.Hour * 4).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &models.AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		models.User{Email: user.Email},
	}

	log.Debug(user)
	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.AuthToken{
		Token:     tokenString,
		TokenType: "Bearer",
		ExpiresIn: expiresAt,
	})
}
