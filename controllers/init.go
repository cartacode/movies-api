package controllers

import (
	"github.com/VuliTv/api/dbh"
	"github.com/VuliTv/api/libs/logging"
)

var connection, err = dbh.NewConnection("controllers")
var log = logging.GetProdLog()
var page = 0
var perpage = 20
var depth = 0
