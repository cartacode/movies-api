/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

// VolumeInformation --
type VolumeInformation struct {
	Director *Performer `json:"director"`

	Studio *Studio `json:"studio"`

	// List of Mongo ObjectID for the Performers in this scene. Embeddable
	Performers []string `json:"performers"`
}
