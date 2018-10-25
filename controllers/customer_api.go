/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"gopkg.in/mgo.v2/bson"
)

// JSONPaginationResponseCustomer --
// HTTP status code 200 and repository model in data
// swagger:response customerResp
type JSONPaginationResponseCustomer struct {
	// in: body
	Results       []models.Customer `json:"results"`
	TotalResults  int               `json:"total"`
	RecordsOnPage int               `json:"recordsonpage"`
	Page          int               `json:"page"`
	TotalPages    int               `json:"totalpages"`
}

// CustomerCustomerIDDelete --
func CustomerCustomerIDDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerCustomerIDGet --
func CustomerCustomerIDGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerCustomerIDPatch --
func CustomerCustomerIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerGet --
func CustomerGet(w http.ResponseWriter, r *http.Request) {

	// Get all seriess
	results := connection.Collection("series").Find(bson.M{})

	// Make a list of series to add together
	series := &models.Series{}

	retval := []models.Series{}

	// Get pagination information
	pagination, err := results.Paginate(requests.GetPaginationInfo(r))

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(series) {
		retval = append(retval, *series)

	}

	// Make our pagination response
	response := JSONPaginationResponseSeries{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	requests.ReturnAPIOK(w, rs)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// CustomerPost --
func CustomerPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
