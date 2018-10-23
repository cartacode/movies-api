/*
 * Vuli API
 *
 * Vuli Studio Delivery API
 *
 * API version: 3

 */

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VuliTv/api/libs/http_helper"
	"github.com/VuliTv/api/models"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// JSONPaginationResponseStudio --
type JSONPaginationResponseStudio struct {
	Results       []models.Studio `json:"results"`
	TotalResults  int             `json:"total"`
	RecordsOnPage int             `json:"recordsonpage"`
	Page          int             `json:"page"`
	TotalPages    int             `json:"totalpages"`
}

// StudioStudioIDDelete --
func StudioStudioIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studioID := params["studioID"]
	studio := &models.Studio{}

	log.Info(studioID)
	// Check valid bson id
	if !bson.IsObjectIdHex(studioID) {
		httphelper.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("studio").FindById(bson.ObjectIdHex(studioID), studio)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("studio").DeleteDocument(studio)
	log.Info(err)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := httphelper.JSONSuccessResponse{Message: "success", Identifier: studio.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}

// StudioStudioIDGet -- Takes Studio ID for a finder
func StudioStudioIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	studioID := params["studioID"]
	studio := &models.Studio{}

	// Check valid bson id
	if !bson.IsObjectIdHex(studioID) {
		httphelper.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("studio").FindById(bson.ObjectIdHex(studioID), studio)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(studio)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)

}

// StudioSlugGet -- Takes Studio ID for a finder
func StudioSlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	studio := &models.Studio{}
	err := connection.Collection("studio").FindOne(bson.M{"slug": slug}, studio)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(studio)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}

// StudioStudioIDPatch --
func StudioStudioIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// StudioGet --
func StudioGet(w http.ResponseWriter, r *http.Request) {

	// Get all studios
	results := connection.Collection("studio").Find(bson.M{})

	// See if we are given a page number to iteratate with #?page=2
	pageQuery, ok := r.URL.Query()["page"]

	// #TODO: Add error handling
	if ok {
		page, err = strconv.Atoi(pageQuery[0])
	}

	// See if we are given a per page number to iteratate with #?perpage=25
	perQuery, ok := r.URL.Query()["perpage"]
	if ok {
		perpage, err = strconv.Atoi(perQuery[0])
	}

	// Make a list of studios to add together
	studio := &models.Studio{}

	retval := []models.Studio{}

	// Get pagination information
	pagination, err := results.Paginate(perpage, page)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(studio) {
		retval = append(retval, *studio)

	}

	// Make our pagination response
	response := JSONPaginationResponseStudio{
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

// StudioPost --
func StudioPost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	studio := &models.Studio{}
	if r.Body == nil {
		httphelper.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&studio)
	if err != nil {
		log.Error(err)
		httphelper.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("studio").Save(studio)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		httphelper.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(studio)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}
