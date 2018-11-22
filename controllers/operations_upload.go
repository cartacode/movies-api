package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
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
	objectid := query["id"].(string)
	collection := params["collection"]

	// Check for a hexId
	if !bson.IsObjectIdHex(objectid) {
		requests.ReturnAPIError(w, fmt.Errorf("Not a valid bson Id"))
		return
	}
	// Read the image
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new session to AWS
	s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if requests.ReturnOnError(w, err) {
		return
	}

	// Quick verification step
	_, err = s.Config.Credentials.Get()

	if requests.ReturnOnError(w, err) {
		return
	}

	// if field == "image" {
	// 	// Upload to our bucket
	// 	path, err := requests.AddFileToS3(s, bucket, "media/image/available/"+objectid+"/"+field, content)
	// 	if requests.ReturnOnError(w, err) {
	// 		return
	// 	}

	// 	// Patch the collection document with the new image path
	// 	patch := make(map[string]string)
	// 	patch[field] = path
	// 	err = connection.Collection(collection).Collection().Update(bson.M{"_id": bson.ObjectIdHex(objectid)}, bson.M{"$set": patch})

	// 	if requests.ReturnOnError(w, err) {
	// 		return
	// 	}
	// } else {

	// Upload to our bucket
	path, err = requests.AddFileToS3(s, bucket, "media/"+objectid+"/images/"+field, content)
	if requests.ReturnOnError(w, err) {
		return
	}
	// Patch the collection document with the new image path
	patch := make(map[string]string)
	patch["images."+field] = path
	err = connection.Collection(collection).Collection().Update(bson.M{"_id": bson.ObjectIdHex(objectid)}, bson.M{"$set": patch})

	if requests.ReturnOnError(w, err) {
		return
	}

	// Sending our response
	response := &requests.JSONSuccessResponse{Message: path, Identifier: "success"}
	js, err := json.Marshal(response)

	if requests.ReturnOnError(w, err) {
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
	objectid := query["id"].(string)
	collection := params["collection"]

	// Check for a hexId
	if !bson.IsObjectIdHex(objectid) {
		if requests.ReturnOnError(w, fmt.Errorf("Not a valid bson Id")) {
			return
		}
	}
	log.Debugw("looking for collection")
	model, err := models.ModelByCollection(collection)

	if requests.ReturnOnError(w, err) {
		return
	}

	// Find the document
	err = connection.Collection(collection).FindById(bson.ObjectIdHex(objectid), model)
	if requests.ReturnOnError(w, err) {
		return
	}

	content, err := ioutil.ReadAll(r.Body)

	if requests.ReturnOnError(w, err) {
		return
	}

	// Create a new session to AWS
	s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if requests.ReturnOnError(w, err) {
		return
	}

	// Quick verification step
	_, err = s.Config.Credentials.Get()

	if requests.ReturnOnError(w, err) {
		return
	}
	log.Infow("logging into AWS")

	// Upload to our bucket
	path, err := requests.AddFileToS3(s, bucket, "media/"+objectid+"/trailers/"+field, content)
	if requests.ReturnOnError(w, err) {
		return
	}

	log.Infow("upload successful")
	patch := models.Trailer{Title: field}

	log.Debugw("creating new patch", "object", patch)

	switch collection {
	case "scene":
		log.Infow("found scene model")
		scene := *model.(*models.Scene)

		scene.Trailer = patch
		err = connection.Collection(collection).Save(&scene)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			requests.ReturnAPIError(w, vErr.Errors[0])
			return
		}
	case "movie":
		log.Infow("found movie model")
		movie := *model.(*models.Movie)

		movie.Trailer = patch
		err = connection.Collection(collection).Save(&movie)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			requests.ReturnAPIError(w, vErr.Errors[0])
			return
		}

	case "volume":
		log.Infow("found volume model")
		volume := *model.(*models.Volume)

		volume.Trailer = patch
		err = connection.Collection(collection).Save(&volume)
		if vErr, ok := err.(*bongo.ValidationError); ok {
			requests.ReturnAPIError(w, vErr.Errors[0])
			return
		}
	default:
		requests.ReturnOnError(w, fmt.Errorf("No such model"))
		return
	}

	// // Sending our response
	response := &requests.JSONSuccessResponse{Message: path, Identifier: "success", Extra: patch}
	js, err := json.Marshal(response)

	if requests.ReturnOnError(w, err) {
		return
	}

	requests.ReturnAPIOK(w, js)
}
