package requests

import (
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// QuerySanatizer --
func QuerySanatizer(params map[string][]string) map[string]interface{} {

	query := make(map[string]interface{})

	for rawParam := range params {

		// default value for switch
		var value interface{}
		var err interface{}

		// fmt.Println(reflect.TypeOf(params[rawParam][0]))
		switch rawParam {
		case "star__contains":
			value, err = strconv.ParseBool(params[rawParam][0])
			rawParam = "star"
		case "_id":
			if bson.IsObjectIdHex(params[rawParam][0]) {
				value = bson.ObjectIdHex(params[rawParam][0])

			} else {
				continue
			}

		case "reviewed":
			value, err = strconv.ParseBool(params[rawParam][0])
		case "page":
			continue
		case "perpage":
			continue
		case "perPage":
			continue
		default:

			varType := reflect.TypeOf(params[rawParam][0]).Kind()

			// See if it's a string
			if len(params[rawParam]) == 1 && varType == reflect.String {
				if bson.IsObjectIdHex(params[rawParam][0]) {
					value = bson.ObjectIdHex(params[rawParam][0])

				} else {
					switch strings.ToLower(params[rawParam][0]) {
					case "true":
						value = true
						log.Debugw("converted bool type", rawParam, value)
					case "false":
						value = false
						log.Debugw("converted bool type", rawParam, value)
					default:
						value = params[rawParam][0]
					}
				}

			} else {
				value = params[rawParam][0]

			}
		}

		if err != nil {
			log.Error(err)
			return nil
		}
		query[rawParam] = value
	}

	log.Debugw("full query", "string", query)
	return query
}
