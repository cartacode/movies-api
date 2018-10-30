package models

import (
	"fmt"

	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var connection, err = dbh.NewConnection("models")
var log = logging.GetProdLog()

// ModelByCollection --
func ModelByCollection(collection string) (interface{}, error) {
	switch collection {
	case "movie":
		model := &Movie{}
		return model, nil

	}
	switch collection {
	case "series":
		model := &Series{}
		return model, nil

	}
	switch collection {
	case "performer":
		model := &Performer{}
		return model, nil

	}
	switch collection {
	case "scene":
		model := &Scene{}
		return model, nil

	}
	switch collection {
	case "volume":
		model := &Volume{}
		return model, nil

	}
	switch collection {
	case "customer":
		model := &Customer{}
		return model, nil

	}

	err := fmt.Errorf("No collection found")
	return "", err
}
