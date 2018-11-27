package models

import (
	"fmt"

	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var connection, dbError = dbh.NewMongoDBConnection("models")
var err error
var log = logging.GetProdLog()

func init() {
	if dbError != nil {
		panic(err)
	}

}

// ModelByCollection --
func ModelByCollection(collection string) (interface{}, error) {

	switch collection {
	case "movie":
		model := &Movie{}
		return model, nil

	case "series":
		model := &Series{}
		return model, nil

	case "star":
		model := &Star{}
		return model, nil

	case "scene":
		model := &Scene{}
		return model, nil

	case "volume":
		model := &Volume{}
		return model, nil

	case "studio":
		model := &Studio{}
		return model, nil

	}
	err := fmt.Errorf("No collection found")
	return "", err
}
