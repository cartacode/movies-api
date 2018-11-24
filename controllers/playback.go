package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
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

	if requests.ReturnOnError(w, err) {
		return
	}

	// Check valid bson id

	if !bson.IsObjectIdHex(objectID) {
		if requests.ReturnOnError(w, fmt.Errorf("Not a valid bson Id")) {
			return
		}

	}

	// Find doc
	err = connection.Collection(collection).FindById(bson.ObjectIdHex(objectID), model)

	if requests.ReturnOnError(w, err) {
		return
	}

	var DynamoID string

	switch collection {
	case "scene":
		log.Infow("found scene model")
		scene := *model.(*models.Scene)
		DynamoID = scene.DynamoDBId
	case "movie":
		log.Infow("found movie model")
		movie := *model.(*models.Movie)
		DynamoID = movie.DynamoDBId

	default:
		requests.ReturnOnError(w, fmt.Errorf("No such model"))
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

	if requests.ReturnOnError(w, err) {
		return
	}

	item := models.MediaDynamoRecord{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if requests.ReturnOnError(w, err) {
		return
	}

	// Json
	a := &RQ{HLS: item.HlsURL, Dash: item.DashURL}

	log.Info(a)
	js, err := json.Marshal(a)

	if requests.ReturnOnError(w, err) {
		return
	}

	requests.ReturnAPIOK(w, js)
}
