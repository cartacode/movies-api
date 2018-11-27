package dbh

import (
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	auth "gopkg.in/hunterlong/authorizecim.v1"
)

// NewAuthorizeNetSession checks for validity of Authorize.net
// credentials
func NewAuthorizeNetSession() (bool, error) {
	log.Infow("new Authorize.net handler created")

	// TODO - fetch fallback values from config
	apiID := envhelp.GetEnv("AUTHORIZE_API_ID", "65Vv2fYQ")
	apiKey := envhelp.GetEnv("AUTHORIZE_API_KEY", "422uVB78H7dn3BcH")

	auth.SetAPIInfo(apiID, apiKey, envhelp.GetEnv("AUTHORIZE_API_ENVIRONMENT", "test"))

	var err error
	if _, err := auth.IsConnected(); err != nil {
		return false, err
	}
	return true, err
}
