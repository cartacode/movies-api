package requests

import (
	"net/http"

	"github.com/VuliTv/go-movie-api/models"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

// GetAuthUser --
func GetAuthUser(r *http.Request) (*models.AuthUser, error) {
	var authUser models.AuthUser
	if err := mapstructure.Decode(context.Get(r, "decoded"), &authUser); err != nil {
		return nil, err
	}
	return &authUser, nil
}
