package denormalized

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var mongoHandler dbh.MongoDBHandler
var redisHandler dbh.RedisHandler

const controller string = "denormalized"

func init() {
	if err := mongoHandler.New(controller); err != nil {
		log.Fatal(err)
	}
	if err := redisHandler.New(controller); err != nil {
		log.Fatal(err)
	}
}
