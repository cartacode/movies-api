/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

// Scene Document
//
// A scene can be associated with a volume, and series
//
// swagger:model
type Scene struct {
	bongo.DocumentBase `bson:",inline"`

	Category []string `json:"category"`

	Information *VolumeInformation `json:"information"`

	// Unique Slug for this scene. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// Description of this scene if it has one. Not required
	Description string `json:"description"`

	// Calculated by user view. Only increases.
	Views int32 `json:"views"`

	// Calculated by user input. Only decreases.
	Downvotes int32 `json:"downvotes"`

	// Unique Title for this scene
	Title string `json:"title"`

	// Calculated externally and maintained here
	Rank int32 `json:"rank"`

	// True/False. Has someone reviewed this scene
	Reviewed bool `json:"reviewed"`

	Volume *Volume `json:"volume"`

	// Total scene length in seconds
	Length int32 `json:"length"`

	// Calculated by user input. Only increases.
	Upvotes int32 `json:"upvotes"`

	// Read only value. Only Admin can update. Sets the price for a scene
	Price float32 `json:"price"`

	// Free list of tag strings
	Tags []string `json:"tags"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Scene) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	scene := &Scene{}

	// Find by slug when posting new scene
	err := connection.Collection("scene").FindOne(bson.M{"slug": s.Slug}, scene)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}
	return retval
}
