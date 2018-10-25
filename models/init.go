package models

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var connection, err = dbh.NewConnection("models")
var log = logging.GetProdLog()
