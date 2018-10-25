/*
 * Vuli API
 *
 * Vuli Series Delivery API
 *
 * API version: 3

 */

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// JSONPaginationResponseSeries --
// HTTP status code 200 and repository model in data
// swagger:response seriesResp
type JSONPaginationResponseSeries struct {
	// in: body
	Results       []models.Series `json:"results"`
	TotalResults  int             `json:"total"`
	RecordsOnPage int             `json:"recordsonpage"`
	Page          int             `json:"page"`
	TotalPages    int             `json:"totalpages"`
}

// SeriesSeriesIDDelete --
func SeriesSeriesIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	seriesID := params["seriesID"]
	series := &models.Series{}

	log.Info(seriesID)
	// Check valid bson id
	if !bson.IsObjectIdHex(seriesID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("series").FindById(bson.ObjectIdHex(seriesID), series)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("series").DeleteDocument(series)
	log.Info(err)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := requests.JSONSuccessResponse{Message: "success", Identifier: series.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// SeriesSeriesIDGet -- Takes Series ID for a finder
func SeriesSeriesIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	seriesID := params["seriesID"]
	series := &models.Series{}

	// Check valid bson id
	if !bson.IsObjectIdHex(seriesID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("series").FindById(bson.ObjectIdHex(seriesID), series)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(series)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)

}

// SeriesSlugGet -- Takes Series ID for a finder
func SeriesSlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	series := &models.Series{}
	err := connection.Collection("series").FindOne(bson.M{"slug": slug}, series)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(series)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// SeriesSeriesIDPatch --
func SeriesSeriesIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// SeriesGet --
func SeriesGet(w http.ResponseWriter, r *http.Request) {

	// Get all seriess
	results := connection.Collection("series").Find(bson.M{})

	// Make a list of seriess to add together
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

}

// SeriesPost --
func SeriesPost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	series := &models.Series{}
	if r.Body == nil {
		requests.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&series)
	if err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("series").Save(series)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(series)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}
