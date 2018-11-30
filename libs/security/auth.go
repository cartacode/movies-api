package security

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

// JWTSecret --
// This needs to go into an ENV variable or secret file. This is for dev only
var JWTSecret = "B=dd4KgP3vPQ%3Y*Bt*7zfJngXcu=_&nwa&?ZgvdZhczqERseSe*^95zfAx8XrR939w5utP6%X^uF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSk6CU!3Gmgt+^mXW^j-AFB=dd4KgP3vPQ%3Y*Bt*7zfJngXcu=_&nwaZgvdZhczqERseSe*^95zfAx8XrR939w5utP6%X^uF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSB=dd4KgP3vPQ%3Y*Bt*7zfJngXcu=_&nwa&?ZgvdZhcuF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSk6CU!3Gmgt+^mXW^j-AF3vPQ%3Y*Bt*7zfJngXcu=_&nwa&?ZgvdZhczqERseSe*^95zfAx8XrR939w5utP6%X^uF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSk6CU!3Gmgt+^mXW^j-AF"

// AuthErrorMsg ...
// Custom error object
type AuthErrorMsg struct {
	Message string `json:"message"`
}

// AuthToken ...
// This is what is retured to the user
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthUser --
type AuthUser struct {
	Email    string `json:"email"`
	ObjectID string `json:"objectID"`
	Admin    bool   `json:"admin"`
}

// TokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type TokenClaim struct {
	*jwt.StandardClaims
	AuthUser
}

// GetAuthUser --
func GetAuthUser(r *http.Request) (*AuthUser, error) {
	var authUser AuthUser
	if err := mapstructure.Decode(context.Get(r, "decoded"), &authUser); err != nil {
		return nil, err
	}
	return &authUser, nil
}

// ValidateTokenMiddleware --
func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		authorizationHeader := req.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(JWTSecret), nil

				})
				if error != nil {
					json.NewEncoder(w).Encode(AuthErrorMsg{Message: error.Error()})
					return
				}
				if token.Valid {

					// // return
					var authUser AuthUser
					mapstructure.Decode(token.Claims, &authUser)

					val, err := redis.Get(authUser.ObjectID).Result()

					if err != nil {
						log.Error(err)
						json.NewEncoder(w).Encode(AuthErrorMsg{Message: "Invalid authorization token - Token expired"})
						return
					}

					if val != bearerToken[1] {
						log.Error(err)
						json.NewEncoder(w).Encode(AuthErrorMsg{Message: "Invalid authorization token - Does not pass validation"})
						return
					}

					// Check admin route priveleges
					if IsAdminRoute(req) {
						log.Info("checking admin route")
						if !authUser.Admin {
							log.Warnw("not an admin!", "user", authUser.Email, "admin", authUser.Admin)
							json.NewEncoder(w).Encode(AuthErrorMsg{Message: "You don't have the proper permissions"})
							return
						}
					}

					log.Debug(req.RequestURI)
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(AuthErrorMsg{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(AuthErrorMsg{Message: "Invalid authorization token"})
			}
		} else {
			if IsOpenAuthedRoute(req) {
				log.Info("found open or auth route")
				next(w, req)
				return
			}
			json.NewEncoder(w).Encode(AuthErrorMsg{Message: "An authorization header is required"})
		}
	})
}

//IsAdminRoute --
func IsAdminRoute(req *http.Request) bool {

	switch req.RequestURI {
	case "/v1/collection/studio":
		return true
	}

	return false
}

// IsOpenAuthedRoute --
func IsOpenAuthedRoute(req *http.Request) bool {

	if req.Method == "GET" {
		if strings.Contains(req.URL.Path, "/v1/data/") {
			return true
		}
	}

	return false
}
