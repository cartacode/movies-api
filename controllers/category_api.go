/*
 * Vuli API
 *
 * Vuli Category Delivery API
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

// JSONPaginationResponseCategory --
// HTTP status code 200 and repository model in data
// swagger:response categoryResp
type JSONPaginationResponseCategory struct {
	// in: body
	Results       []models.Category `json:"results"`
	TotalResults  int               `json:"total"`
	RecordsOnPage int               `json:"recordsonpage"`
	Page          int               `json:"page"`
	TotalPages    int               `json:"totalpages"`
}

// CategoryCategoryIDDelete --
func CategoryCategoryIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID := params["categoryID"]
	category := &models.Category{}

	log.Info(categoryID)
	// Check valid bson id
	if !bson.IsObjectIdHex(categoryID) {
		httphelper.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("category").FindById(bson.ObjectIdHex(categoryID), category)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection("category").DeleteDocument(category)
	log.Info(err)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	// Send the response
	response := httphelper.JSONSuccessResponse{Message: "success", Identifier: category.Id.String()}

	js, err := json.Marshal(response)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}

// CategoryCategoryIDGet -- Takes Category ID for a finder
func CategoryCategoryIDGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	categoryID := params["categoryID"]
	category := &models.Category{}

	// Check valid bson id
	if !bson.IsObjectIdHex(categoryID) {
		httphelper.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err := connection.Collection("category").FindById(bson.ObjectIdHex(categoryID), category)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(category)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)

}

// CategorySlugGet -- Takes Category ID for a finder
func CategorySlugGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	slug := params["slug"]
	category := &models.Category{}
	err := connection.Collection("category").FindOne(bson.M{"slug": slug}, category)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}

	js, err := json.Marshal(category)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}

// CategoryCategoryIDPatch --
func CategoryCategoryIDPatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(http.StatusOK)
}

// CategoryGet --
func CategoryGet(w http.ResponseWriter, r *http.Request) {

	// Get all categorys
	results := connection.Collection("category").Find(bson.M{})

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

	// Make a list of categorys to add together
	category := &models.Category{}

	retval := []models.Category{}

	// Get pagination information
	pagination, err := results.Paginate(perpage, page)

	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	// Get which page we are on to skip
	results.Query.Skip(page * perpage)

	// Add the found results
	for results.Next(category) {
		retval = append(retval, *category)

	}

	// Make our pagination response
	response := JSONPaginationResponseCategory{
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

// CategoryPost --
func CategoryPost(w http.ResponseWriter, r *http.Request) {
	// text := slug.Make("Hellö Wörld хелло ворлд")
	category := &models.Category{}
	if r.Body == nil {
		httphelper.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Error(err)
		httphelper.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection("category").Save(category)
	if vErr, ok := err.(*bongo.ValidationError); ok {
		httphelper.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(category)
	if err != nil {
		httphelper.ReturnAPIError(w, err)
		return
	}
	httphelper.ReturnAPIOK(w, js)
}
