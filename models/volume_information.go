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
	Director *Star `json:"director"`

	Studio *Studio `json:"studio"`

	// List of Mongo ObjectID for the Stars in this scene. Embeddable
	Stars []string `json:"Stars"`
}
