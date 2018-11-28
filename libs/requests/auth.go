package requests

import (
	"net/http"

	"github.com/VuliTv/go-movie-api/app/customer"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

// GetAuthUser --
func GetAuthUser(r *http.Request) (*customer.AuthUser, error) {
	var authUser customer.AuthUser
	if err := mapstructure.Decode(context.Get(r, "decoded"), &authUser); err != nil {
		return nil, err
	}
	return &authUser, nil
}
