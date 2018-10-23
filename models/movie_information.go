/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// MovieInformation --
type MovieInformation struct {
	Director []string `json:"director,omitempty"`

	Studio string `json:"studio,omitempty"`

	// List of Mongo ObjectID for the Performers in this movie. Embeddable
	Performers []string `json:"performers,omitempty"`
}
