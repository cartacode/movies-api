package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/VuliTv/go-movie-api/controllers"
	"github.com/VuliTv/go-movie-api/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

var authorizationRoutes = Routes{
	Route{
		"Authorization",
		"POST",
		"/v1/authorize",
		controllers.Authorization,
	},
}

func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret"), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(models.ErrorMsg{Message: error.Error()})
					return
				}
				if token.Valid {

					// return
					var user models.User
					mapstructure.Decode(token.Claims, &user)

					vars := mux.Vars(req)
					name := vars["userId"]
					if name != user.Email {
						json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
						return
					}

					log.Info(user)
					log.Debug(user)
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(models.ErrorMsg{Message: "An authorization header is required"})
		}
	})
}
