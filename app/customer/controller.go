/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/libs/stringops"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// WishlistUpdateRequest ..
type WishlistUpdateRequest struct {
	Collection string         `json:"collection"`
	ID         *bson.ObjectId `json:"id"`
}

// WishlistResponse ..
type WishlistResponse struct {
	// Liked Items
	Liked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*bson.ObjectId `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*bson.ObjectId `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*bson.ObjectId `json:"volumes"`

		// List of Mongo ObjectId for the stars wish list. Embeddable
		Stars []*bson.ObjectId `json:"stars"`
	} `json:"liked"`

	// Liked Items
	Disliked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*bson.ObjectId `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*bson.ObjectId `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*bson.ObjectId `json:"volumes"`
	} `json:"disliked"`

	Wishlist struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*bson.ObjectId `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*bson.ObjectId `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*bson.ObjectId `json:"volumes"`
	} `json:"wishlist"`
}

// ListAddItem --
func ListAddItem(w http.ResponseWriter, r *http.Request) {
	// Get auth user information
	params := mux.Vars(r)
	allowedLists := []string{"wishlist", "liked", "disliked"}
	if !stringops.StringInSlice(params["list"], allowedLists) {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, "only allowed methods ["+strings.Join(allowedLists, ",")+"]"))
		return
	}

	var authUser, err = requests.GetAuthUser(r)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	var wishlistRequest WishlistUpdateRequest
	if err = json.NewDecoder(r.Body).Decode(&wishlistRequest); err != nil {
		log.Warn("Request Body parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	list := fmt.Sprintf("%s.%ss", params["list"], wishlistRequest.Collection)
	if err := mongoHandler.Collection(collection).Collection().Update(bson.M{"_id": bson.ObjectIdHex(authUser.ObjectID)}, bson.M{"$push": bson.M{list: wishlistRequest.ID}}); err != nil {
		log.Errorw("unable to add item to list",
			"list", params["list"],
			"collection", wishlistRequest.Collection,
			"doc_id", wishlistRequest.ID,
			"user", authUser.Email,
			"error", err,
		)
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return

	}
	// Sending our response
	message := fmt.Sprintf("added %s to %s", wishlistRequest.ID.Hex(), params["list"])
	response := &requests.JSONSuccessResponse{Message: message, Identifier: "success"}
	js, err := json.Marshal(response)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	requests.ReturnAPIOK(w, js)
}

// WishlistDeleteItem --
func WishlistDeleteItem(w http.ResponseWriter, r *http.Request) {
	// Get auth user information
	params := mux.Vars(r)

	allowedLists := []string{"wishlist", "liked", "disliked"}
	if !stringops.StringInSlice(params["list"], allowedLists) {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, "only allowed methods ["+strings.Join(allowedLists, ",")+"]"))
		return
	}
	var authUser, err = requests.GetAuthUser(r)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	var wishlistRequest WishlistUpdateRequest
	if err = json.NewDecoder(r.Body).Decode(&wishlistRequest); err != nil {
		log.Warn("Request Body parse error: ", err.Error())
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return
	}

	list := fmt.Sprintf("%s.%ss", params["list"], wishlistRequest.Collection)
	if err := mongoHandler.Collection(collection).Collection().UpdateId(bson.ObjectIdHex(authUser.ObjectID), bson.M{"$pull": bson.M{list: wishlistRequest.ID}}); err != nil {
		log.Errorw("unable to remove item from list",
			"list", params["list"],
			"collection", wishlistRequest.Collection,
			"doc_id", wishlistRequest.ID,
			"user", authUser.Email,
			"error", err,
		)
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		return

	}
	// Sending our response
	message := fmt.Sprintf("removed %s from %s", wishlistRequest.ID.Hex(), params["list"])
	response := &requests.JSONSuccessResponse{Message: message, Identifier: "success"}
	js, err := json.Marshal(response)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	requests.ReturnAPIOK(w, js)
}

// ProfileGet --
func ProfileGet(w http.ResponseWriter, r *http.Request) {
	var authUser, err = requests.GetAuthUser(r)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	user := &Model{}
	if err := mongoHandler.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	profile := &ProfileInformationResponse{Model: *user}

	profile.Password = "**************************"
	profile.Wishlist.Movies = addDenormalizedDataFromSlice("movie", user.Wishlist.Movies)
	profile.Wishlist.Scenes = addDenormalizedDataFromSlice("scene", user.Wishlist.Scenes)
	profile.Wishlist.Volumes = addDenormalizedDataFromSlice("volume", user.Wishlist.Volumes)

	profile.Purchased.Movies = addDenormalizedDataFromSlice("movie", user.Purchased.Movies)
	profile.Purchased.Scenes = addDenormalizedDataFromSlice("scene", user.Purchased.Scenes)
	profile.Purchased.Volumes = addDenormalizedDataFromSlice("volume", user.Purchased.Volumes)

	profile.Disliked.Movies = addDenormalizedDataFromSlice("movie", user.Disliked.Movies)
	profile.Disliked.Scenes = addDenormalizedDataFromSlice("scene", user.Disliked.Scenes)
	profile.Disliked.Volumes = addDenormalizedDataFromSlice("volume", user.Disliked.Volumes)

	profile.Liked.Movies = addDenormalizedDataFromSlice("movie", user.Liked.Movies)
	profile.Liked.Scenes = addDenormalizedDataFromSlice("scene", user.Liked.Scenes)
	profile.Liked.Volumes = addDenormalizedDataFromSlice("volume", user.Liked.Volumes)

	js, err := json.Marshal(profile)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	requests.ReturnAPIOK(w, js)
}

func addDenormalizedDataFromSlice(collection string, objectIDS []*bson.ObjectId) []*ModelStub {
	if objectIDS == nil {
		return nil
	}
	retval := []*ModelStub{}
	for _, object := range objectIDS {
		objectStub := &ModelStub{}
		if err := mongoHandler.Collection(collection).FindById(*object, &objectStub); err != nil {
			log.Warnw("can't find embedded doc", "collection", collection, "objectId", object.Hex(), "error", err)
		} else {
			objectStub.ID = object
			retval = append(retval, objectStub)
		}
	}

	return retval
}

func addDenormalizedData(collection string, objectId *bson.ObjectId) *ModelStub {

	if objectId == nil {
		return nil
	}
	retval := &ModelStub{}
	if err := mongoHandler.Collection(collection).FindById(*objectId, &retval); err != nil {
		log.Warn(err)

	}

	return retval
}
