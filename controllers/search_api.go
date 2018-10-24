package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/VuliTv/api/dbh"
	"github.com/VuliTv/api/libs/http_helper"
	"github.com/VuliTv/api/models"
)

// MovieSearchGet -- Takes Movie ID for a finder
func MovieSearchGet(w http.ResponseWriter, r *http.Request) {

	// query params
	query := dbh.QuerySanatizer(r.URL.Query())

	// get out resuls
	results := connection.Collection("movie").Find(query)

	// See if we are given a page number to iteratate with #?page=2
	pageQuery, ok := r.URL.Query()["page"]

	// #TODO: Add error handling
	if ok {
		page, err = strconv.Atoi(pageQuery[0])
	}
	// See if we are given a page number to iteratate with #?page=2
	depthQuery, ok := r.URL.Query()["depth"]

	if ok {
		depth, err = strconv.Atoi(depthQuery[0])
	}

	// See if we are given a per page number to iteratate with #?perpage=25
	perQuery, ok := r.URL.Query()["perpage"]
	if ok {
		perpage, err = strconv.Atoi(perQuery[0])
	}

	// Make a list of movies to add together
	movie := &models.Movie{}

	retval := []models.Movie{}

	// Get pagination information
	pagination, err := results.Paginate(perpage, page)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(movie) {
		retval = append(retval, *movie)

	}

	// Make our pagination response
	response := JSONPaginationResponseMovie{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	httphelper.ReturnAPIOK(w, rs)
}

// SceneSearchGet -- Takes Movie ID for a finder
func SceneSearchGet(w http.ResponseWriter, r *http.Request) {

	// query params
	query := dbh.QuerySanatizer(r.URL.Query())

	// get out resuls
	results := connection.Collection("scene").Find(query)

	// See if we are given a page number to iteratate with #?page=2
	pageQuery, ok := r.URL.Query()["page"]

	// #TODO: Add error handling
	if ok {
		page, err = strconv.Atoi(pageQuery[0])
	}
	// See if we are given a page number to iteratate with #?page=2
	depthQuery, ok := r.URL.Query()["depth"]

	if ok {
		depth, err = strconv.Atoi(depthQuery[0])
	}

	// See if we are given a per page number to iteratate with #?perpage=25
	perQuery, ok := r.URL.Query()["perpage"]
	if ok {
		perpage, err = strconv.Atoi(perQuery[0])
	}

	// Make a list of scenes to add together
	scene := &models.Scene{}

	retval := []models.Scene{}

	// Get pagination information
	pagination, err := results.Paginate(perpage, page)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(scene) {
		retval = append(retval, *scene)

	}

	// Make our pagination response
	response := JSONPaginationResponseScene{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	httphelper.ReturnAPIOK(w, rs)
}
