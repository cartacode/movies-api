package search

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var mongoHandler dbh.MongoDBHandler

var log = logging.GetProdLog()

func init() {
	if err := mongoHandler.New("search"); err != nil {
		log.Fatal(err)
	}
}
