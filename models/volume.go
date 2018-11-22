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

// Volume Document
//
// A volume can be associated with a Series
//
// swagger:model
type Volume struct {
	bongo.DocumentBase `bson:",inline"`

	// Media information
	Images     Images     `json:"images"`
	Extras     []Extras   `json:"extras"`
	Thumbnails Thumbnails `json:"thumbnails"`
	Trailers   Trailer    `json:"trailers"`

	// MovieInformation --
	Information MediaInformation `json:"information"`

	// Media Performance
	Performance Performance `json:"performance"`
	Category    []string    `json:"category"`

	// Unique Slug for this scene. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// Description of this scene if it has one. Not required
	Description string `json:"description"`

	// Calculated by user view. Only increases.
	Views int32 `json:"views"`

	Series string `json:"series"`

	// Unique Title for this scene
	Title string `json:"title"`

	// True/False. Has someone reviewed this scene
	Reviewed bool `json:"reviewed"`
	// Read only value. Only Admin can update. Sets the price for a the volume which supersedes the scene price
	Price float32 `json:"price"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Volume) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	volume := &Volume{}

	// Find by slug when posting new volume
	err := connection.Collection("volume").FindOne(bson.M{"slug": s.Slug}, volume)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}
	return retval
}
