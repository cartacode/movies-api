package dbh

import (
	"github.com/VuliTv/api/libs/envhelp"
	"github.com/go-bongo/bongo"
)

var config = &bongo.Config{
	ConnectionString: envhelp.GetEnv("MONGO_HOST", "mongodb"),
	Database:         envhelp.GetEnv("MONGO_DB", "vuliapi"),
}

// NewConnection --
func NewConnection(caller string) (*bongo.Connection, error) {
	log.Infow("new database handler created",
		"connection_string", config.ConnectionString,
		"database", config.Database,
		"caller", caller,
	)

	connection, err := bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}

	return connection, err
}
