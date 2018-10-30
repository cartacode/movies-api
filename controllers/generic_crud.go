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

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// JSONPaginationResponse --
// HTTP status code 200 and repository model in data
// swagger:response Resp
type JSONPaginationResponse struct {
	// in: body
	Results       []interface{} `json:"results"`
	TotalResults  int           `json:"total"`
	RecordsOnPage int           `json:"recordsonpage"`
	Page          int           `json:"page"`
	TotalPages    int           `json:"totalpages"`
}

// GenericCrudGet --
func GenericCrudGet(w http.ResponseWriter, r *http.Request) {

	var model interface{}
	var retval []interface{}

	query := requests.QuerySanatizer(r.URL.Query())
	params := mux.Vars(r)
	collection := params["collection"]

	// fmt.Println(collection)
	// get out resuls
	results := connection.Collection(collection).Find(query)

	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	fmt.Println(perPage, page)
	pagination, err := results.Paginate(perPage, page)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	fmt.Println(collection)
	fmt.Println(results)

	model, err = models.ModelByCollection(collection)
	if err != nil {
		log.Fatal(err)
	}

	// Add the found results
	for results.Next(model) {
		retval = append(retval, model)

	}
	// Make our pagination response
	response := JSONPaginationResponse{
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

// GenericCrudPost --
func GenericCrudPost(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	collection := params["collection"]

	model, err := models.ModelByCollection(collection)
	if err != nil {
		log.Fatal(err)
	}
	// text := slug.Make("Hellö Wörld хелло ворлд")

	if r.Body == nil {
		requests.ReturnAPIError(w, fmt.Errorf("Please send a request body"))
		return
	}

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, err)
		return
	}
	err = connection.Collection(collection).Save(model.(bongo.Document))
	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, vErr.Errors[0])
		return
	}

	// Return the saved document
	js, err := json.Marshal(model)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// GenericCrudIDGet --
func GenericCrudIDGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	objectid := params["objectid"]
	collection := params["collection"]
	model, err := models.ModelByCollection(collection)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(objectid)
	// Check valid bson id

	// fmt.Println(hex.Enc(objectid))
	if !bson.IsObjectIdHex(objectid) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err = connection.Collection(collection).FindById(bson.ObjectIdHex(objectid), model)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Json
	js, err := json.Marshal(model)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	requests.ReturnAPIOK(w, js)
}

// GenericCrudIDDelete --
func GenericCrudIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectid := params["objectid"]
	collection := params["collection"]
	model, err := models.ModelByCollection(collection)
	if err != nil {
		log.Fatal(err)
	}

	// Check valid bson id
	if !bson.IsObjectIdHex(objectid) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}

	// Find doc
	err = connection.Collection(collection).FindById(bson.ObjectIdHex(objectid), model)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}
	// Delete the document
	err = connection.Collection(collection).DeleteDocument(model.(bongo.Document))
	log.Info(err)
	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	// Send the response
	retval := model.(bongo.Document)
	response := requests.JSONSuccessResponse{Message: "success", Identifier: retval.GetId().String()}

	js, err := json.Marshal(response)

	if err != nil {
		requests.ReturnAPIError(w, err)
		return
	}

	requests.ReturnAPIOK(w, js)
}
