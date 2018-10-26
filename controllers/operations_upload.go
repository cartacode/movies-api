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

	s, err := session.NewSession(&aws.Config{Region: aws.String(s3Region)})
	if err != nil {
		log.Fatal(err)
	}
	_, err = s.Config.Credentials.Get()
	if err != nil {
		log.Fatal(err)
	}

	// Upload
	path, err := requests.AddFileToS3(s, bucket, "cover-art/"+objectid, content)
	if err != nil {
		log.Fatal(err)
	}
	response := &requests.JSONSuccessResponse{Message: path, Identifier: "success"}
	js, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	switch collection {
	case "movie":

		movie := &models.Movie{}
		// Find doc
		err := connection.Collection(collection).FindById(bson.ObjectIdHex(objectid), movie)

		// fmt.Println(dnfError)
		if err != nil {
			requests.ReturnAPIError(w, err)
			return
		}

		movie.MediaContent.CoverImage = path
		err = connection.Collection(collection).Save(movie)
		if err != nil {
			requests.ReturnAPIError(w, err)
			return
		}
	}
	requests.ReturnAPIOK(w, js)
}
