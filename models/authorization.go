package models

import jwt "github.com/dgrijalva/jwt-go"

// ErrorMsg ...
// Custom error object
type ErrorMsg struct {
	Message string `json:"message"`
}

// User ...
// Custom object which can be stored in the claims
type User struct {
	Email string `json:"username"`
	Token string `json:"token"`
}

// AuthToken ...
// This is what is retured to the user
type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// AuthTokenClaim ...
// This is the cliam object which gets parsed from the authorization header
type AuthTokenClaim struct {
	*jwt.StandardClaims
	User
}
