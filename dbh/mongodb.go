package dbh

import (
	"github.com/VuliTv/go-movie-api/libs/envhelp"

	"github.com/go-bongo/bongo"
)

var config = &bongo.Config{
	ConnectionString: envhelp.GetEnv("MONGO_HOST", "localhost"),
	Database:         envhelp.GetEnv("MONGO_DATABASE", "vuliapi"),
}

// NewConnection --
func NewConnection(caller string) (*bongo.Connection, error) {
	log.Infow("new database handler created",
		"connection_string", config.ConnectionString,
		"database", config.Database,
		"caller", caller,
	)

	connection, err := bongo.Connect(config)

	log.Infow("testing connection")

	_, err = connection.Session.BuildInfo()
	if err != nil {
		panic(err)
	}
	err = connection.Session.Ping()
	if err != nil {
		panic(err)
	}

	return connection, err
}
