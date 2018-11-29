package security

import (
	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var redis dbh.RedisHandler
var log = logging.GetProdLog()

func init() {
	if err := redis.New("security-auth"); err != nil {
		log.Fatal(err)
	}
}
