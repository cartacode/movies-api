package payments

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var mongoHandler dbh.MongoDBHandler

var collection = "customer"

func init() {
	dbh.NewAuthorizeNetSession()
	if err := mongoHandler.New(collection); err != nil {
		log.Fatal(err)
	}
}
