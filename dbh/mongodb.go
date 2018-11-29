package dbh

import (
	"github.com/VuliTv/go-movie-api/libs/envhelp"

	"github.com/go-bongo/bongo"
)

// MongoDBHandler --
type MongoDBHandler struct {
	*bongo.Connection
}

var mongoConfig = &bongo.Config{
	ConnectionString: envhelp.GetEnv("MONGO_HOST", "localhost"),
	Database:         envhelp.GetEnv("MONGO_DATABASE", "vuliapi"),
}

// New --
func (m *MongoDBHandler) New(caller string) error {
	log.Infow("new database handler created",
		"connection_string", mongoConfig.ConnectionString,
		"database", mongoConfig.Database,
		"caller", caller,
	)

	connection, err := bongo.Connect(mongoConfig)

	log.Infow("testing connection")

	_, err = connection.Session.BuildInfo()
	if err != nil {
		return err
	}
	err = connection.Session.Ping()

	if err != nil {
		return err
	}

	m.Connection = connection
	return err
}
