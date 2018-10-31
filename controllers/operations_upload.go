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
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// s3region
var s3Region = envhelp.GetEnv("AWS_DEFAULT_REGION", "us-east-1")
var bucket = "vuli-public-assets"

// OperationsUploadCoverImage --
func OperationsUploadCoverImage(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Body must be set", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	objectid := params["objectid"]
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

	// Upload to our bucket
	path, err := requests.AddFileToS3(s, bucket, "media/"+objectid+"/cover-image", content)
	if requests.ReturnOnError(w, err) {
		return
	}

	// Patch the collection document with the new image path
	patch := make(map[string]string)
	patch["coverimage"] = path
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
	objectid := params["objectid"]
	collection := params["collection"]
	slug := params["slug"]

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

	// Upload to our bucket
	path, err := requests.AddFileToS3(s, bucket, "media/"+objectid+"/trailers/"+slug, content)
	if requests.ReturnOnError(w, err) {
		return
	}

	// Patch the collection document with the new image path
	patch := models.Trailer{Title: slug}

	patch.Title = slug
	patch.Path = path

	fmt.Println(patch)
	err = connection.Collection(collection).Collection().Update(bson.M{"_id": bson.ObjectIdHex(objectid)}, bson.M{"$addToSet": bson.M{"trailers": patch}})

	if requests.ReturnOnError(w, err) {
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
