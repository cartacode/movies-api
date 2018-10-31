package models

import (
	"fmt"
	"reflect"

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
		fmt.Println(reflect.TypeOf(model))
		return model, nil

	case "series":
		model := &Series{}
		return model, nil

	case "performer":
		model := &Performer{}
		return model, nil

	case "scene":
		model := &Scene{}
		return model, nil

	case "volume":
		model := &Volume{}
		return model, nil

	case "customer":
		model := &Customer{}
		return model, nil

	case "category":
		model := &Category{}
		return model, nil

	case "studio":
		model := &Studio{}
		return model, nil

	}
	err := fmt.Errorf("No collection found")
	return "", err
}
