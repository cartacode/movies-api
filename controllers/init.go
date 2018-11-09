package controllers

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var connection, dbError = dbh.NewConnection("controllers")
var err error
var log = logging.GetProdLog()

func init() {
	if dbError != nil {
		panic(err)
	}

}
