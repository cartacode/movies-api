package requests

import (
	"net/http"
	"strconv"
)

// GetPaginationInfo --
func GetPaginationInfo(r *http.Request) (int, int) {

	page := 1
	perPage := 20
	var err interface{}
	// See if we are given a page number to iteratate with #?page=2
	pageQuery, ok := r.URL.Query()["page"]

	// #TODO: Add error handling
	if ok {
		page, err = strconv.Atoi(pageQuery[0])
	}

	// See if we are given a per page number to iteratate with #?perpage=25
	perQuery, ok := r.URL.Query()["perpage"]
	if ok {
		perPage, err = strconv.Atoi(perQuery[0])
	}

	if err != nil {
		log.Error(err)
		return 1, 20
	}
	return perPage, page
}
