package controllers

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
	"github.com/go-bongo/bongo"
)

var connection *bongo.Connection
var err error
var log = logging.GetProdLog()

func init() {
	connection, err = dbh.NewConnection("controllers")
	if err != nil {
		panic(err)
	}

}
