package auth

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var mongoHandler dbh.MongoDBHandler
var redisHandler dbh.RedisHandler
var sesHandler dbh.SeSHandler

var log = logging.GetProdLog()
var collection = "customer"

func init() {
	if err := mongoHandler.New("auth"); err != nil {
		log.Fatal(err)
	}
	if err := redisHandler.New("auth"); err != nil {
		log.Fatal(err)
	}
	if err := sesHandler.New("auth"); err != nil {
		log.Fatal(err)
	}
}
