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
	if err := connection.Collection(collection).Collection().Update(bson.M{"_id": bson.ObjectIdHex(authUser.ObjectID)}, bson.M{"$push": bson.M{list: wishlistRequest.ID}}); err != nil {
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
	if err := connection.Collection(collection).Collection().UpdateId(bson.ObjectIdHex(authUser.ObjectID), bson.M{"$pull": bson.M{list: wishlistRequest.ID}}); err != nil {
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

	user := Model{}
	if err := connection.Collection(collection).FindById(bson.ObjectIdHex(authUser.ObjectID), &user); err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	// for _, e := range user.Wishlist.Scenes {
	// response.Wishlist.Scenes = append(response.Wishlist.Scenes, *e)
	// }

	// Don't give real passwords
	user.Password = "**************************"
	js, err := json.Marshal(user)

	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}

	requests.ReturnAPIOK(w, js)
}
