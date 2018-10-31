package requests

import (
	"strconv"
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
		case "reviewed":
			value, err = strconv.ParseBool(params[rawParam][0])
			log.Debugw("converted bool type", rawParam, value)
		case "page":
			continue
		case "perpage":
			continue
		case "perPage":
			continue
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
