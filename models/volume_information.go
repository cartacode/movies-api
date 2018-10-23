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
	Director *Performer `json:"director,omitempty"`

	Studio *Studio `json:"studio,omitempty"`

	// List of Mongo ObjectID for the Performers in this scene. Embeddable
	Performers []string `json:"performers,omitempty"`
}
