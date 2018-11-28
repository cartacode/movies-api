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
	connection.New(collection)
	if err := rDB.New(collection); err != nil {
		log.Fatal(err)
	}
}
