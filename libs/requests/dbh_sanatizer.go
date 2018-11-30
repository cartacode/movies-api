package requests

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

// QuerySanatizer --
func QuerySanatizer(params map[string][]string) (map[string]interface{}, error) {

	query := make(map[string]interface{})

	for rawParam := range params {

		// default value for switch
		var value interface{}
		var err error
		// fmt.Println(reflect.TypeOf(params[rawParam][0]))
		switch rawParam {
		case "star__contains":
			value, err = strconv.ParseBool(params[rawParam][0])
			if err != nil {
				return nil, fmt.Errorf("reviewed is true/false")
			}
			rawParam = "star"
		case "_id":
			if bson.IsObjectIdHex(params[rawParam][0]) {
				value = bson.ObjectIdHex(params[rawParam][0])

			} else {
				return nil, fmt.Errorf("_id must be a valid bson id")
			}

		case "reviewed":
			value, err = strconv.ParseBool(params[rawParam][0])
			if err != nil {
				return nil, fmt.Errorf("reviewed is true/false")
			}
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

		query[rawParam] = value
	}

	log.Debugw("full query", "string", query)
	return query, nil
}
