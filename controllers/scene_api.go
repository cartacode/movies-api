/*
 * Vuli API
 *
 * Vuli Scene Delivery API
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

// JSONPaginationResponseScene --
// HTTP status code 200 and repository model in data
// swagger:response sceneResp
type JSONPaginationResponseScene struct {
	// in: body
	Results       []models.Scene `json:"results"`
	TotalResults  int            `json:"total"`
	RecordsOnPage int            `json:"recordsonpage"`
	Page          int            `json:"page"`
	TotalPages    int            `json:"totalpages"`
}

// SceneSceneIDDelete --
func SceneSceneIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sceneID := params["sceneID"]
	scene := &models.Scene{}

	log.Info(sceneID)
	// Check valid bson id
	if !bson.IsObjectIdHex(sceneID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("scene").FindById(bson.ObjectIdHex(sceneID), scene)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("scene").DeleteDocument(scene)
	log.Info(err)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := requests.JSONSuccessResponse{Message: "success", Identifier: scene.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// SceneSceneIDGet -- Takes Scene ID for a finder
func SceneSceneIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sceneID := params["sceneID"]
	scene := &models.Scene{}

	// Check valid bson id
	if !bson.IsObjectIdHex(sceneID) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("scene").FindById(bson.ObjectIdHex(sceneID), scene)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(scene)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)

}

// SceneSlugGet -- Takes Scene ID for a finder
func SceneSlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	scene := &models.Scene{}
	err := connection.Collection("scene").FindOne(bson.M{"slug": slug}, scene)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(scene)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// SceneSceneIDPatch --
func SceneSceneIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// SceneGet --
func SceneGet(w http.ResponseWriter, r *http.Request) {

	// Get all scenes
	results := connection.Collection("scene").Find(bson.M{})

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

	// Make a list of scenes to add together
	scene := &models.Scene{}

	retval := []models.Scene{}

	// Get pagination information
	pagination, err := results.Paginate(perpage, page)

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

// ScenePost --
func ScenePost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	scene := &models.Scene{}
	if r.Body == nil {
		requests.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&scene)
	if err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("scene").Save(scene)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(scene)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}
