package customer

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var mongoHandler dbh.MongoDBHandler
var redisHandler dbh.RedisHandler

const collection string = "customer"

func init() {
	if err := mongoHandler.New(collection); err != nil {
		log.Fatal(err)
	}
	if err := redisHandler.New(collection); err != nil {
		log.Fatal(err)
	}
}
