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
	"fmt"
	"net/http"
	"strconv"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// JSONPaginationResponseMovie --
// HTTP status code 200 and repository model in data
// swagger:response movieResp
type JSONPaginationResponseMovie struct {
	// in: body
	Results       []models.Movie `json:"results"`
	TotalResults  int            `json:"total"`
	RecordsOnPage int            `json:"recordsonpage"`
	Page          int            `json:"page"`
	TotalPages    int            `json:"totalpages"`
}

// MovieMovieIDDelete --
func MovieMovieIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["movieID"]
	movie := &models.Movie{}

	log.Info(movieID)
	// Check valid bson id
	if !bson.IsObjectIdHex(movieID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("movie").FindById(bson.ObjectIdHex(movieID), movie)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("movie").DeleteDocument(movie)
	log.Info(err)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := requests.JSONSuccessResponse{Message: "success", Identifier: movie.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// MovieMovieIDGet -- Takes Movie ID for a finder
func MovieMovieIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID := params["movieID"]
	movie := &models.Movie{}

	// Check valid bson id
	if !bson.IsObjectIdHex(movieID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("movie").FindById(bson.ObjectIdHex(movieID), movie)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(movie)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)

}

// MovieSlugGet -- Takes Movie ID for a finder
func MovieSlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	movie := &models.Movie{}
	err := connection.Collection("movie").FindOne(bson.M{"slug": slug}, movie)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(movie)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// MovieMovieIDPatch --
func MovieMovieIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// MovieGet --
func MovieGet(w http.ResponseWriter, r *http.Request) {

	// Get all movies
	results := connection.Collection("movie").Find(bson.M{})

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
		requests.ReturnAPIError(w, err)
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
		requests.ReturnAPIError(w, err)
		return
	}

	requests.ReturnAPIOK(w, rs)

}

// MoviePost --
func MoviePost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	movie := &models.Movie{}
	if r.Body == nil {
		requests.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("movie").Save(movie)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(movie)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}
