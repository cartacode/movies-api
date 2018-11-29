package customer

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var connection dbh.MongoDBHandler
var rDB dbh.RedisHandler

var collection = "customer"

func init() {
	if err := connection.New(collection); err != nil {
		log.Fatal(err)
	}

	if err := rDB.New(collection); err != nil {
		log.Fatal(err)
	}
}
