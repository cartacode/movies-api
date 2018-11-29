/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package crud

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/models"
	"github.com/VuliTv/go-movie-api/libs/requests"

	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// GenericCrudGet --
func GenericCrudGet(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}
	query := requests.QuerySanatizer(r.URL.Query())
	log.Debugw("query running", "Q", query)
	params := mux.Vars(r)
	collection := params["collection"]

	// get out resuls
	results := mongoHandler.Collection(collection).Find(query)

	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	pagination, err := results.Paginate(perPage, page)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	model, err := models.ModelByCollection(collection)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Warnw("collection find error", "collection", collection, "error", err)
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
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
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
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}
	// text := slug.Make("Hellö Wörld хелло ворлд")

	if r.Body == nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, "Please send a request body")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = mongoHandler.Collection(collection).Save(model.(bongo.Document))
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	if vErr, ok := err.(*bongo.ValidationError); ok {
		requests.ReturnAPIError(w, http.StatusBadRequest, vErr.Errors[0].Error())
		return
	}

	// Return the saved document
	js, err := json.Marshal(model)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	requests.ReturnAPIOK(w, js)
}

// GenericCrudIDGet --
func GenericCrudIDGet(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	objectID := params["objectID"]
	collection := params["collection"]
	model, err := models.ModelByCollection(collection)

	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	// Check valid bson id

	if !bson.IsObjectIdHex(objectID) {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id"))
		return
	}

	// Find doc
	if err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(objectID), &model); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Json
	js, error := json.Marshal(model)

	if error != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
	}

	requests.ReturnAPIOK(w, js)
}

// GenericCrudIDDelete --
func GenericCrudIDDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectID := params["objectID"]
	collection := params["collection"]
	model, err := models.ModelByCollection(collection)
	if err != nil {
		log.Fatal(err)
	}

	// Check valid bson id
	if !bson.IsObjectIdHex(objectID) {
		requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id")
		return
	}

	// Find doc
	err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(objectID), model)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}
	// Delete the document
	err = mongoHandler.Collection(collection).DeleteDocument(model.(bongo.Document))
	log.Info(err)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Send the response
	retval := model.(bongo.Document)
	response := requests.JSONSuccessResponse{Message: "success", Identifier: retval.GetId().String()}

	js, err := json.Marshal(response)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	requests.ReturnAPIOK(w, js)
}

// GenericCrudIDPatch --
func GenericCrudIDPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	objectID := params["objectID"]
	collection := params["collection"]
	// model, err := models.ModelByCollection(collection)
	// if err != nil {
	// log.Fatal(err)
	// }

	// Check valid bson id
	if !bson.IsObjectIdHex(objectID) {
		requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id")
		return
	}

	var patchBody interface{}

	if r.Body == nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, "Please send a request body")
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&patchBody); err != nil {
		log.Error(err)
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Update the document
	if err := mongoHandler.Collection(collection).Collection().Update(
		bson.M{"_id": bson.ObjectIdHex(objectID)}, bson.M{"$set": patchBody}); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := requests.JSONSuccessResponse{Message: "success", Identifier: objectID, Extra: patchBody}

	if js, err := json.Marshal(response); err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
	} else {
		requests.ReturnAPIOK(w, js)
	}

}
