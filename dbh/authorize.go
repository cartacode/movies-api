package dbh

import (
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	auth "github.com/hunterlong/authorizecim"
)

// NewAuthorizeNetSession checks for validity of Authorize.net
// credentials
func NewAuthorizeNetSession() (bool, error) {
	log.Infow("new Authorize.net handler created")

	// TODO - fetch fallback values from config
	apiName := envhelp.GetEnv("apiName", "65Vv2fYQ")
	apiKey := envhelp.GetEnv("apiKey", "6z59j6XrGA7V3TbG")

	auth.SetAPIInfo(apiName, apiKey, "test")

	var status bool
	var err error
	if status, err = auth.IsConnected(); err != nil {
		panic(err)
	}

	return status, err
}
