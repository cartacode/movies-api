package denormalized

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var mongoHandler dbh.MongoDBHandler

const controller string = "denormalized"

func init() {
	if err := mongoHandler.New(controller); err != nil {
		log.Fatal(err)
	}
}
