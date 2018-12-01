package payments

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var mongoHandler dbh.MongoDBHandler

var collection = "customer"

func init() {
	if ok, err := dbh.ConnectAuthorizeNet(); !ok {
		log.Fatalw("unable to connect to authorize.net", "error", err)
	}
	if err := mongoHandler.New(collection); err != nil {
		log.Fatalw("unable to connect to mongoDB", "error", err)
	}

}
