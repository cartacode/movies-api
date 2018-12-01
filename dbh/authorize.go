package dbh

import (
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/go-bongo/bongo"
	AuthorizeCIM "gopkg.in/hunterlong/authorizecim.v1"
)

// AuthorizeNetHandler --
type AuthorizeNetHandler struct {
	*bongo.Connection
}

var apiID = envhelp.GetEnv("AUTHORIZE_ID", "65Vv2fYQ")
var apiKey = envhelp.GetEnv("AUTHORIZE_TRANSACTION_KEY", "2cL2W24uV35aKw4M")
var authMode = envhelp.GetEnv("AUTHORIZE_API_ENVIRONMENT", "test")

// ConnectAuthorizeNet checks for validity of Authorize.net
// credentials
func ConnectAuthorizeNet() (bool, error) {

	var err error

	if ok, _ := AuthorizeCIM.IsConnected(); ok {
		return true, err
	}
	log.Infow("new Authorize.net handler created")

	AuthorizeCIM.SetAPIInfo(apiID, apiKey, authMode)

	if ok, err := AuthorizeCIM.IsConnected(); !ok {

		return ok, err
	}
	return true, err
}
