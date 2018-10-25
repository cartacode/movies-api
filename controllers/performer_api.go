/*
 * Vuli API
 *
 * Vuli Performer Delivery API
 *
 * API version: 3

 */

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VuliTv/api/libs/requests"
	"github.com/VuliTv/api/models"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// JSONPaginationResponsePerformer --
// HTTP status code 200 and repository model in data
// swagger:response performerResp
type JSONPaginationResponsePerformer struct {
	// in:body
	Results       []models.Performer `json:"results"`
	TotalResults  int                `json:"total"`
	RecordsOnPage int                `json:"recordsonpage"`
	Page          int                `json:"page"`
	TotalPages    int                `json:"totalpages"`
}

// PerformerPerformerIDDelete --
func PerformerPerformerIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	performerID := params["performerID"]
	performer := &models.Performer{}

	log.Info(performerID)
	// Check valid bson id
	if !bson.IsObjectIdHex(performerID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("performer").FindById(bson.ObjectIdHex(performerID), performer)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("performer").DeleteDocument(performer)
	log.Info(err)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := requests.JSONSuccessResponse{Message: "success", Identifier: performer.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// PerformerPerformerIDGet -- Takes Performer ID for a finder
func PerformerPerformerIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	performerID := params["performerID"]
	performer := &models.Performer{}

	// Check valid bson id
	if !bson.IsObjectIdHex(performerID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("performer").FindById(bson.ObjectIdHex(performerID), performer)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(performer)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)

}

// PerformerSlugGet -- Takes Performer ID for a finder
func PerformerSlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	performer := &models.Performer{}
	err := connection.Collection("performer").FindOne(bson.M{"slug": slug}, performer)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(performer)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// PerformerPerformerIDPatch --
func PerformerPerformerIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// PerformerGet --
func PerformerGet(w http.ResponseWriter, r *http.Request) {

	// Get all performers
	results := connection.Collection("performer").Find(bson.M{})

	// Make a list of performers to add together
	performer := &models.Performer{}

	retval := []models.Performer{}

	// Get pagination information
	pagination, err := results.Paginate(requests.GetPaginationInfo(r))

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(performer) {
		retval = append(retval, *performer)

	}

	// Make our pagination response
	response := JSONPaginationResponsePerformer{
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

// PerformerPost --
func PerformerPost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	performer := &models.Performer{}
	if r.Body == nil {
		requests.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&performer)
	if err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("performer").Save(performer)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(performer)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}
