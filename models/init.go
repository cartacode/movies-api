package models

import (
	"github.com/VuliTv/api/dbh"
	"github.com/VuliTv/api/libs/logging"
)

var connection, err = dbh.NewConnection("models")
var log = logging.GetProdLog()
