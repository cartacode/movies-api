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

					if err != nil {
						log.Error(err)
						json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token - Token expired"})
						return
					}

					if val != bearerToken[1] {
						log.Error(err)
						json.NewEncoder(w).Encode(models.ErrorMsg{Message: "Invalid authorization token - Does not pass validation"})
						return
					}

					// Check admin route priveleges
					if isAdminRoute(req) {
						log.Info("checking admin route")
						if !authUser.Admin {
							log.Warnw("not an admin!", "user", authUser.Email, "admin", authUser.Admin)
							json.NewEncoder(w).Encode(models.ErrorMsg{Message: "You don't have the proper permissions"})
							return
						}
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
			if isOpenAuthedRoute(req) {
				log.Info("found open or auth route")
				next(w, req)
				return
			}
			json.NewEncoder(w).Encode(models.ErrorMsg{Message: "An authorization header is required"})
		}
	})
}

func isAdminRoute(req *http.Request) bool {

	switch req.RequestURI {
	case "/v1/collection/studio":
		switch req.Method {
		case "POST":
			return true
		}
	}

	return false
}

func isOpenAuthedRoute(req *http.Request) bool {

	switch req.RequestURI {
	case "/v1/data/movie":
		switch req.Method {
		case "GET":
			return true
		}
	case "/v1/data/scene":
		switch req.Method {
		case "GET":
			return true
		}
	case "/v1/data/volume":
		switch req.Method {
		case "GET":
			return true
		}
	case "/v1/data/series":
		switch req.Method {
		case "GET":
			return true
		}
	case "/v1/data/star":
		switch req.Method {
		case "GET":
			return true
		}
	}

	return false
}
