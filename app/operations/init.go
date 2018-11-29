package operations

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()
var mongoHandler dbh.MongoDBHandler

func init() {
	if err := mongoHandler.New("operations"); err != nil {
		log.Fatal(err)
	}
}
