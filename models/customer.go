/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// Customer Document
//
// A customer on the site and all of their preferences and profile
//
// swagger:model
type Customer struct {

	// Unique email for this customer, read only, cognito controlled
	// read only: true
	Email string `json:"email"`

	// True/False. Is the user active
	Active bool `json:"active"`

	// Liked Items
	Liked struct {

		// List of Mongo ObjectID for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectID for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectID for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`

		// List of Mongo ObjectID for the stars wish list. Embeddable
		Stars []string `json:"stars"`
	} `json:"liked"`

	// Credit Information
	Credit struct {

		// Has the user stored credit information
		InfoStored bool `json:"info_stored"`

		// Key for 3 leg transactions to provider bank
		Key string `json:"key"`
	} `json:"credit"`

	// Purchased Items
	Purchased struct {

		// List of Mongo ObjectID for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectID for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectID for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`
	} `json:"purchased"`

	// User wishlist
	Wishlist struct {

		// List of Mongo ObjectID for the movies wish list. Embeddable
		Movies []string `json:"movies"`

		// List of Mongo ObjectID for the scenes wish list. Embeddable
		Scenes []string `json:"scenes"`

		// List of Mongo ObjectID for the volumes wish list. Embeddable
		Volumes []string `json:"volumes"`
	} `json:"wishlist"`
}
