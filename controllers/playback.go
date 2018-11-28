package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"

	"github.com/VuliTv/go-movie-api/libs/models"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// RQ --
type RQ struct {
	HLS  string `json:"hls"`
	Dash string `json:"dash"`
}

// SignedS3Playback --
// SignedS3Playback --
func SignedS3Playback(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	objectID := params["objectID"]
	collection := params["collection"]
	model, err := models.ModelByCollection(collection)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	// Check valid bson id

	if !bson.IsObjectIdHex(objectID) {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id"))
		return

	}

	// Find doc
	if err = connection.Collection(collection).FindById(bson.ObjectIdHex(objectID), model); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	var DynamoID string

	switch collection {
	case "scene":
		log.Infow("found scene model")
		scene := *model.(*scene.Model)
		DynamoID = scene.DynamoDBId
	case "movie":
		log.Infow("found movie model")
		movie := *model.(*movie.Model)
		DynamoID = movie.DynamoDBId

	default:
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, "No such model"))
		return
	}

	log.Info(DynamoID)
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("vuli-media-pipeline"),
		Key: map[string]*dynamodb.AttributeValue{
			"guid": {
				S: aws.String(DynamoID),
			},
		},
	})

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	item := media.DynamoRecord{}

	if err = dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	// Json
	a := &RQ{HLS: item.HlsURL, Dash: item.DashURL}

	log.Info(a)
	js, err := json.Marshal(a)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	requests.ReturnAPIOK(w, js)
}
