package requests

import (
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/logging"
	"github.com/VuliTv/go-movie-api/libs/security"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

var log = logging.GetProdLog()

// GetAuthUser --
func GetAuthUser(r *http.Request) (*security.AuthUser, error) {
	var authUser security.AuthUser
	if err := mapstructure.Decode(context.Get(r, "decoded"), &authUser); err != nil {
		return nil, err
	}
	return &authUser, nil
}
