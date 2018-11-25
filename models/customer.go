/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

import "github.com/go-bongo/bongo"

// Customer Document
//
// A customer on the site and all of their preferences and profile
//
// swagger:model
type Customer struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique email for this customer, read only, cognito controlled
	// read only: true
	Email string `json:"email"`

	// Stored as a bcrypt hash
	Password string

	// True/False. Is the user active
	Active bool `json:"active"`

	// Are they an admin or now. Allows for POST/PUT/PATCH to CRUD routes
	Admin bool `json:"admin"`

	// Liked Items
	Liked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`

		// List of Mongo ObjectId for the stars wish list. Embeddable
		Stars []string `json:"stars"`
	} `json:"liked"`

	// Liked Items
	Disliked struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`

		// List of Mongo ObjectId for the stars wish list. Embeddable
		Stars []string `json:"stars"`
	} `json:"disliked"`

	// Credit Information
	Credit struct {

		// Has the user stored credit information
		InfoStored bool `json:"info_stored"`

		// Key for 3 leg transactions to provider bank
		Key string `json:"key"`
	} `json:"credit"`

	// Purchased Items
	Purchased struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`
	} `json:"purchased"`

	// User wishlist
	Wishlist struct {

		// List of Mongo ObjectId for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectId for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectId for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`
	} `json:"wishlist"`

	Preferences []struct {
		Tag    string  `json:"tag"`
		Weight float64 `json:"weight"`
	} `json:"preferences"`
}
