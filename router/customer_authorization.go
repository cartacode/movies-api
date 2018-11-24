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
	"github.com/mitchellh/mapstructure"
)

var authorizationRoutes = Routes{
	Route{
		"Login",
		"POST",
		"/v1/authorize/login",
		controllers.Login,
	},
	Route{
		"Signup",
		"POST",
		"/v1/authorize/signup",
		controllers.Signup,
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
					return []byte(models.JWTSecret), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(models.ErrorMsg{Message: error.Error()})
					return
				}
				if token.Valid {

					// // return
					var authUser models.AuthUser
					mapstructure.Decode(token.Claims, &authUser)

					val, err := rDB.Get(authUser.ObjectID).Result()

					if val != bearerToken[1] {
						log.Error(err)
						json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
						return
					}
					if err != nil {
						log.Error(err)
						json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
						return
					}
					// log.Info(token.Claims)

					// vars := mux.Vars(req)
					// name := vars["customerId"]
					// log.Info(authUser)
					// log.Debug(vars)
					// if name != authUser.Email {
					// 	json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token - Does not match UserID"})
					// 	return
					// }

					if isAdminRoute(req) && !authUser.Admin {
						json.NewEncoder(w).Encode(models.ErrorMsg{Message: "You don't have the proper permissions"})
						return
					}
					log.Debug(req.RequestURI)
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

func isAdminRoute(req *http.Request) bool {

	switch req.Method {
	case "POST":
		log.Warnw("invalid attempt to POST to admin route!")
		return true
	}
	switch req.RequestURI {
	case "/v1/collection/studio":
		log.Info("studio")
	}

	return false
}
