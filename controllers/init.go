package controllers

import (
	"math/rand"
	"time"

	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var connection dbh.MongoDBHandler
var rDB dbh.RedisHandler
var sesConn dbh.SeSHandler
var err error
var log = logging.GetProdLog()
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {

	if err := connection.New("controllers"); err != nil {
		log.Fatal(err)
	}
	if err := rDB.New("controllers"); err != nil {
		log.Fatal(err)
	}

	if err := sesConn.New(); err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())

}

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
