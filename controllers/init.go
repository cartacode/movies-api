package controllers

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
	"github.com/VuliTv/go-movie-api/models"
)

var connection, err = dbh.NewConnection("controllers")
var log = logging.GetProdLog()
var page = 0
var perpage = 20
var depth = 0
var collections map[string]interface{}

func init() {
	collections = make(map[string]interface{})

	collections["customer"] = &models.Customer{}
	collections["category"] = &models.Category{}
}

// ReturnModels --
func ReturnModels(modelName string) interface{} {

	return collections[modelName]
}
