/*
 * Vuli API
 *
 * Vuli Volume Delivery API
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

// JSONPaginationResponseVolume --
// HTTP status code 200 and repository model in data
// swagger:response volumeResp
type JSONPaginationResponseVolume struct {
	// in: body
	Results       []models.Volume `json:"results"`
	TotalResults  int             `json:"total"`
	RecordsOnPage int             `json:"recordsonpage"`
	Page          int             `json:"page"`
	TotalPages    int             `json:"totalpages"`
}

// VolumeVolumeIDDelete --
func VolumeVolumeIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	volumeID := params["volumeID"]
	volume := &models.Volume{}

	log.Info(volumeID)
	// Check valid bson id
	if !bson.IsObjectIdHex(volumeID) {
		httphelper.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("volume").FindById(bson.ObjectIdHex(volumeID), volume)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("volume").DeleteDocument(volume)
	log.Info(err)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := httphelper.JSONSuccessResponse{Message: "success", Identifier: volume.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}

// VolumeVolumeIDGet -- Takes Volume ID for a finder
func VolumeVolumeIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	volumeID := params["volumeID"]
	volume := &models.Volume{}

	// Check valid bson id
	if !bson.IsObjectIdHex(volumeID) {
		httphelper.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("volume").FindById(bson.ObjectIdHex(volumeID), volume)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(volume)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)

}

// VolumeSlugGet -- Takes Volume ID for a finder
func VolumeSlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	volume := &models.Volume{}
	err := connection.Collection("volume").FindOne(bson.M{"slug": slug}, volume)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(volume)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}

// VolumeVolumeIDPatch --
func VolumeVolumeIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// VolumeGet --
func VolumeGet(w http.ResponseWriter, r *http.Request) {

	// Get all volumes
	results := connection.Collection("volume").Find(bson.M{})

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

	// Make a list of volumes to add together
	volume := &models.Volume{}

	retval := []models.Volume{}

	// Get pagination information
	pagination, err := results.Paginate(perpage, page)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
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
		httphelper.ReturnAPIError(w, err)
		return
	}

	httphelper.ReturnAPIOK(w, rs)

}

// VolumePost --
func VolumePost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	volume := &models.Volume{}
	if r.Body == nil {
		httphelper.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&volume)
	if err != nil {
		log.Error(err)
		httphelper.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("volume").Save(volume)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		httphelper.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(volume)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}
