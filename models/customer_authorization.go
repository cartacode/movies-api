package models

import jwt "github.com/dgrijalva/jwt-go"

// JWTSecret --
var JWTSecret = "B=dd4KgP3vPQ%3Y*Bt*7zfJngXcu=_&nwa&?ZgvdZhczqERseSe*^95zfAx8XrR939w5utP6%X^uF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSk6CU!3Gmgt+^mXW^j-AFB=dd4KgP3vPQ%3Y*Bt*7zfJngXcu=_&nwaZgvdZhczqERseSe*^95zfAx8XrR939w5utP6%X^uF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSB=dd4KgP3vPQ%3Y*Bt*7zfJngXcu=_&nwa&?ZgvdZhcuF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSk6CU!3Gmgt+^mXW^j-AF3vPQ%3Y*Bt*7zfJngXcu=_&nwa&?ZgvdZhczqERseSe*^95zfAx8XrR939w5utP6%X^uF5EC7W8gsS3Su38pwyWDuvT5ZrPnCfTSk6CU!3Gmgt+^mXW^j-AF"

// ErrorMsg ...
// Custom error object
type ErrorMsg struct {
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

// AuthTokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	AuthUser
}
