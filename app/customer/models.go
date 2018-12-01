/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package customer

import (
	"fmt"
	"strconv"

	"github.com/VuliTv/go-movie-api/app/media"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Model Document
//
// A customer on the site and all of their preferences and profile
type Model struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique email for this customer, read only, cognito controlled
	// read only: true
	Email string `json:"email"`

	// Stored as a bcrypt hash
	Password string `json:"password"`

	// True/False. Is the user active
	Active bool `json:"active"`

	// Are they an admin or now. Allows for POST/PUT/PATCH to CRUD routes
	Admin bool `json:"admin"`

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

		Studios []*bson.ObjectId `json:"studios"`
	} `json:"liked"`

	// Liked Items
	Disliked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*bson.ObjectId `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*bson.ObjectId `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*bson.ObjectId `json:"volumes"`

		// List of Mongo ObjectId for the stars wish list. Embeddable
		Stars []*bson.ObjectId `json:"stars"`

		Studios []*bson.ObjectId `json:"studios"`
	} `json:"disliked"`

	// Credit Information
	Credit struct {

		// Has the user stored credit information
		InfoStored bool `json:"info_stored"`

		// Key for 3 leg transactions to provider bank
		ProfileID string `json:"profileId"`

		// Key for 3 leg transactions to provider bank
		PaymentID string `json:"paymentId"`
	} `json:"credit"`

	// Purchased Items
	Purchased struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*bson.ObjectId `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*bson.ObjectId `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*bson.ObjectId `json:"volumes"`
	} `json:"purchased"`

	// User wishlist
	Wishlist struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*bson.ObjectId `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*bson.ObjectId `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*bson.ObjectId `json:"volumes"`
	} `json:"wishlist"`

	Preferences []struct {
		Tag    string  `json:"tag"`
		Weight float64 `json:"weight"`
	} `json:"preferences"`
}

// AuthBadAttempt --
func (c *Model) AuthBadAttempt() {
	log.Debugw("adding bad auth attempt", "email", c.Email, "id", c.Id.Hex())
	hash := fmt.Sprintf("auth-%s", c.Id.Hex())
	val, err := redisHandler.Get(hash).Result()
	if err != nil {
		if sErr := redisHandler.Set(hash, 0, 0).Err(); sErr != nil {
			log.Error(sErr)
		}
		log.Error(err)
		return
	}

	log.Debugw("found auth val", "val", val)
	attempts, err := strconv.Atoi(val)
	if err != nil {
		log.Error(err)
	}
	attempts++
	if err = redisHandler.Set(hash, attempts, 0).Err(); err != nil {
		log.Error(err)
	}

}

// AuthReset --
func (c *Model) AuthReset() {
	hash := fmt.Sprintf("auth-%s", c.Id.Hex())
	if err := redisHandler.Set(hash, 0, 0).Err(); err != nil {
		log.Error(err)
	}

}

// AuthLocked --
// We always default to true on failures.
// Keeps us from attack issues if we have backend issues
func (c *Model) AuthLocked() bool {
	hash := fmt.Sprintf("auth-%s", c.Id.Hex())
	val, err := redisHandler.Get(hash).Result()
	if err != nil {
		if err = redisHandler.Set(hash, 0, 0).Err(); err != nil {
			log.Error(err)
			return true
		}
		return false

	}

	attempts, err := strconv.Atoi(val)
	if err != nil {
		log.Error(err)
		return true
	}

	if attempts >= 5 {
		return true
	}
	return false
}

// ModelUnlockRequest --
type ModelUnlockRequest struct {
	Email string `json:"email"`
}

// ModelStub ==
type ModelStub struct {
	ID         *bson.ObjectId `json:"_id"`
	Title      string         `json:"title"`
	Slug       string         `json:"slug"`
	Images     media.Images   `json:"images"`
	Completion int32          `json:"completion"`
	Original   bool           `json:"vuliOriginal"`
}

// ProfileInformationResponse --
type ProfileInformationResponse struct {
	Model `json:",inline"`
	// Purchased Items
	Purchased struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*ModelStub `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*ModelStub `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*ModelStub `json:"volumes"`
	} `json:"purchased"`

	// User wishlist
	Wishlist struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*ModelStub `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*ModelStub `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*ModelStub `json:"volumes"`
	} `json:"wishlist"`

	// Liked Items
	Liked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*ModelStub `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*ModelStub `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*ModelStub `json:"volumes"`

		// List of Mongo ObjectId for the stars wish list. Embeddable
		Stars []*ModelStub `json:"stars"`

		Studios []*ModelStub `json:"studios"`
	} `json:"liked"`

	// Liked Items
	Disliked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []*ModelStub `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []*ModelStub `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []*ModelStub `json:"volumes"`

		// List of Mongo ObjectId for the stars wish list. Embeddable
		Stars []*ModelStub `json:"stars"`

		Studios []*ModelStub `json:"studios"`
	} `json:"disliked"`
}
