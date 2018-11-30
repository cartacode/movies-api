package operations

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"

	"github.com/VuliTv/go-movie-api/libs/envhelp"
	"github.com/VuliTv/go-movie-api/libs/models"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/stringops"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// s3region
var s3Region = envhelp.GetEnv("AWS_DEFAULT_REGION", "us-east-1")
var bucket = "vuli-public-assets"

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
	if err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(objectID), model); err != nil {
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

// UploadImage --
func UploadImage(w http.ResponseWriter, r *http.Request) {

	var path string

	if r.Body == nil {
		http.Error(w, "Body must be set", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	query, err := requests.QuerySanatizer(r.URL.Query())
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}
	// Check for params
	keys := make([]string, len(query))
	for key := range query {
		keys = append(keys, key)
	}
	if !stringops.StringInSlice("key", keys) {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, "requires key and _id params"))
		return
	}
	if !stringops.StringInSlice("id", keys) {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, "requires key and _id params"))
		return
	}
	field := query["key"].(string)
	objectID := query["id"].(bson.ObjectId)
	collection := params["collection"]

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
	path, err = requests.AddFileToS3(s, bucket, "media/"+objectID.Hex()+"/images/"+field, content)
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	// Patch the collection document with the new image path
	patch := make(map[string]string)
	patch["images."+field] = path
	if err = mongoHandler.Collection(collection).Collection().Update(bson.M{"_id": objectID}, bson.M{"$set": patch}); err != nil {
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

// UploadTrailer --
func UploadTrailer(w http.ResponseWriter, r *http.Request) {

	if r.Body == nil {
		http.Error(w, "Body must be set", http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	query, err := requests.QuerySanatizer(r.URL.Query())
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}
	field := query["key"].(string)
	objectID := query["id"].(string)
	collection := params["collection"]

	// Check for a hexId
	if !bson.IsObjectIdHex(objectID) {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, "Not a valid bson Id"))
		return

	}
	log.Debugw("looking for collection")
	model, err := models.ModelByCollection(collection)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}

	// Find the document
	err = mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(objectID), model)
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

	// switch collection {
	// case "scene":
	// 	log.Infow("found scene model")
	// 	scene := *model.(*scene.Model)

	// 	scene.Trailer = patch
	// 	err = mongoHandler.Collection(collection).Save(&scene)
	// 	if vErr, ok := err.(*bongo.ValidationError); ok {
	// 		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, vErr.Errors[0].Error()))
	// 		return
	// 	}
	// case "movie":
	// 	log.Infow("found movie model")
	// 	movie := *model.(*movie.Model)

	// 	movie.Trailer = patch
	// 	err = mongoHandler.Collection(collection).Save(&movie)
	// 	if vErr, ok := err.(*bongo.ValidationError); ok {
	// 		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, vErr.Errors[0].Error()))
	// 		return
	// 	}

	// case "volume":
	// 	log.Infow("found volume model")
	// 	volume := *model.(*volume.Model)

	// 	volume.Trailer = patch
	// 	err = mongoHandler.Collection(collection).Save(&volume)
	// 	if vErr, ok := err.(*bongo.ValidationError); ok {
	// 		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, vErr.Errors[0].Error()))
	// 		return
	// 	}
	// default:
	// 	log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, "No such model"))
	// 	return
	// }

	// // Sending our response
	response := &requests.JSONSuccessResponse{Message: path, Identifier: "success", Extra: patch}
	js, err := json.Marshal(response)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
	}

	requests.ReturnAPIOK(w, js)
}
