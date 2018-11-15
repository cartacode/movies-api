package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GenericSearchGet --
func GenericSearchGet(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}
	query := requests.QuerySanatizer(r.URL.Query())
	log.Debugw("query running", "Q", query)
	params := mux.Vars(r)
	collection := params["collection"]

	log.Debug(params)
	if _, ok := query["title"]; !ok {

		requests.ReturnAPIError(w, fmt.Errorf("title query missing"))
		return
	}

	results := connection.Collection(collection).Find(bson.M{"title": bson.RegEx{Pattern: query["title"].(string), Options: "i"}})

	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	pagination, err := results.Paginate(perPage, page)

	if requests.ReturnOnError(w, err) {
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	model, err := models.ModelByCollection(collection)

	if requests.ReturnOnError(w, err) {
		return
	}
	// Add the found results
	for results.Next(&model) {
		retval = append(retval, model)

	}
	// log.Debug(retval)
	// Make our pagination response
	response := requests.JSONPaginationResponse{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}
	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)

	if requests.ReturnOnError(w, err) {
		return
	}

	requests.ReturnAPIOK(w, rs)
}
