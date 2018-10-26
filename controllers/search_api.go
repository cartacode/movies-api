package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/gorilla/mux"
)

// MovieSearchGet -- Takes Movie ID for a finder
func MovieSearchGet(w http.ResponseWriter, r *http.Request) {

	// query params
	query := requests.QuerySanatizer(r.URL.Query())

	// get out resuls
	results := connection.Collection("movie").Find(query)

	// Make a list of movies to add together
	movie := &models.Movie{}

	retval := []models.Movie{}

	// Get pagination information
	pagination, err := results.Paginate(requests.GetPaginationInfo(r))

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

// SceneSearchGet -- Takes Movie ID for a finder
func SceneSearchGet(w http.ResponseWriter, r *http.Request) {

	// query params
	query := requests.QuerySanatizer(r.URL.Query())

	// get out resuls
	results := connection.Collection("scene").Find(query)

	// Make a list of scenes to add together
	scene := &models.Scene{}

	retval := []models.Scene{}

	// Get pagination information
	pagination, err := results.Paginate(requests.GetPaginationInfo(r))

	if err != nil {
		requests.ReturnAPIError(w, err)
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
		requests.ReturnAPIError(w, err)
		return
	}

	requests.ReturnAPIOK(w, rs)
}

// VolumeSearchGet -- Takes Movie ID for a finder
func VolumeSearchGet(w http.ResponseWriter, r *http.Request) {

	// query params
	query := requests.QuerySanatizer(r.URL.Query())

	// get out resuls
	results := connection.Collection("volume").Find(query)

	// Make a list of volumes to add together
	volume := &models.Volume{}

	retval := []models.Volume{}

	// Get pagination information
	pagination, err := results.Paginate(requests.GetPaginationInfo(r))

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(volume) {
		retval = append(retval, *volume)

	}

	// Make our pagination response
	response := JSONPaginationResponseVolume{
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

// SearchGet -- Takes Movie ID for a finder
func SearchGet(w http.ResponseWriter, r *http.Request) {

	// query params
	query := requests.QuerySanatizer(r.URL.Query())

	params := mux.Vars(r)
	collection := params["collection"]

	// fmt.Println(collection)
	// get out resuls
	results := connection.Collection(collection).Find(query)

	// Get pagination information
	pagination, err := results.Paginate(requests.GetPaginationInfo(r))

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Make a list of volumes to add together
	// model := &models.Volume{}
	// retval := []models.Volume{}
	switch collection {
	case "volume":
		// model := &models.Movie{}
		// retval := &[]models.Movie{}
		// for results.Next(model) {
		// 	retval = append(retval, model)

		// }
		// // Add the found results
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)
	// Make our pagination response
	response := requests.JSONPaginationResponse{
		// Results:       retval,
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
