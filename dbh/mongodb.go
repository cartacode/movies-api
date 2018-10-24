package dbh

import (
	"strconv"

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

// QuerySanatizer --
func QuerySanatizer(params map[string][]string) map[string]interface{} {

	query := make(map[string]interface{})

	for rawParam := range params {

		// default value for switch
		var value interface{}
		var err interface{}

		// fmt.Println(reflect.TypeOf(params[rawParam][0]))
		switch rawParam {
		case "reviewed":
			value, err = strconv.ParseBool(params[rawParam][0])
			log.Debugw("converted bool type", rawParam, value)
		default:
			value = params[rawParam][0]
		}

		if err != nil {
			log.Fatal(err)
		}
		query[rawParam] = value
	}

	return query
}
