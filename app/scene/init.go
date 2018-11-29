package scene

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var connection dbh.MongoDBHandler

func init() {
	connection.New("scene")
}
