/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// CustomerWishlist --
type CustomerWishlist struct {

	// List of Mongo ObjectID for the movies wish list. Embeddable
	Movies []string `json:"movies,omitempty"`

	// List of Mongo ObjectID for the scenes wish list. Embeddable
	Scenes []string `json:"scenes,omitempty"`

	// List of Mongo ObjectID for the volumes wish list. Embeddable
	Volumes []string `json:"volumes,omitempty"`
}
