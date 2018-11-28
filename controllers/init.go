package controllers

import (
	"math/rand"
	"time"

	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var connection, dbError = dbh.NewMongoDBConnection("controllers")
var rDB, rError = dbh.NewRedisConnection()
var sesConn dbh.SeSHandler
var err error
var log = logging.GetProdLog()
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	if dbError != nil {
		panic(err)
	}

	if rError != nil {
		panic(err)
	}

	if err := sesConn.New(); err != nil {
		panic(err)
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
