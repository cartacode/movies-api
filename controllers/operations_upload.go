package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"
	"github.com/VuliTv/go-movie-api/app/volume"
	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/VuliTv/go-movie-api/libs/models"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-bongo/bongo"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// s3region
var s3Region = envhelp.GetEnv("AWS_DEFAULT_REGION", "us-east-1")
var bucket = "vuli-public-assets"

// OperationsUploadImage --
func OperationsUploadImage(w http.ResponseWriter, r *http.Request) {

	var path string

	if r.Body == nil {
		http.Error(w, "Body must be set", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	query := requests.QuerySanatizer(r.URL.Query())
	field := query["key"].(string)
	objectID := query["id"].(string)
	collection := params["collection"]

	// Check for a hexId
	if !bson.IsObjectIdHex(objectID) {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id"))
		return
	}
	// Read the image
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Create a new session to AWS
	s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Quick verification step
	_, err = s.Config.Credentials.Get()

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Upload to our bucket
	path, err = requests.AddFileToS3(s, bucket, "media/"+objectID+"/images/"+field, content)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	// Patch the collection document with the new image path
	patch := make(map[string]string)
	patch["images."+field] = path
	if err = connection.Collection(collection).Collection().Update(bson.M{"_id": bson.ObjectIdHex(objectID)}, bson.M{"$set": patch}); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// Sending our response
	response := &requests.JSONSuccessResponse{Message: path, Identifier: "success"}
	js, err := json.Marshal(response)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	requests.ReturnAPIOK(w, js)
}

// OperationsUploadTrailer --
func OperationsUploadTrailer(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Body must be set", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	query := requests.QuerySanatizer(r.URL.Query())
	field := query["key"].(string)
	objectID := query["id"].(string)
	collection := params["collection"]

	// Check for a hexId
	if !bson.IsObjectIdHex(objectID) {
		if err != nil {
			log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id"))
			return
		}

	}
	log.Debugw("looking for collection")
	model, err := models.ModelByCollection(collection)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	// Find the document
	err = connection.Collection(collection).FindById(bson.ObjectIdHex(objectID), model)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	content, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	// Create a new session to AWS
	s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	// Quick verification step
	_, err = s.Config.Credentials.Get()

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}
	log.Infow("logging into AWS")

	// Upload to our bucket
	path, err := requests.AddFileToS3(s, bucket, "media/"+objectID+"/trailers/"+field, content)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	log.Infow("upload successful")
	patch := media.Trailer{Title: field}

	log.Debugw("creating new patch", "object", patch)

	switch collection {
	case "scene":
		log.Infow("found scene model")
		scene := *model.(*scene.Model)

		scene.Trailer = patch
		err = connection.Collection(collection).Save(&scene)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, vErr.Errors[0].Error()))
			return
		}
	case "movie":
		log.Infow("found movie model")
		movie := *model.(*movie.Model)

		movie.Trailer = patch
		err = connection.Collection(collection).Save(&movie)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, vErr.Errors[0].Error()))
			return
		}

	case "volume":
		log.Infow("found volume model")
		volume := *model.(*volume.Model)

		volume.Trailer = patch
		err = connection.Collection(collection).Save(&volume)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, vErr.Errors[0].Error()))
			return
		}
	default:
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, "No such model"))
		return
	}

	// // Sending our response
	response := &requests.JSONSuccessResponse{Message: path, Identifier: "success", Extra: patch}
	js, err := json.Marshal(response)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
	}

	requests.ReturnAPIOK(w, js)
}
