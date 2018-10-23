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

// Series Document
//
// A series can be associated with N volumes
//
// swagger:model
type Series struct {
	bongo.DocumentBase `bson:",inline"`

	Category []string `json:"category,omitempty"`

	Information *VolumeInformation `json:"information,omitempty"`

	// Unique Slug for this series. Made of <title><studio>-series lowercase and character stripped
	Slug string `json:"slug,omitempty"`

	// Description of this scene if it has one. Not required
	Description string `json:"description,omitempty"`

	// Calculated by user view. Only increases.
	Views int32 `json:"views,omitempty"`

	// Calculated by user input. Only decreases.
	Downvotes int32 `json:"downvotes,omitempty"`

	// Unique Title for this entire series(show)
	Title string `json:"title"`

	// Calculated externally and maintained here
	Rank int32 `json:"rank,omitempty"`

	// True/False. Has someone reviewed this scene
	Reviewed bool `json:"reviewed,omitempty"`

	// Cummulative volume lengths in seconds
	Length int32 `json:"length,omitempty"`

	// Calculated by user input. Only increases.
	Upvotes int32 `json:"upvotes,omitempty"`

	// Read only value. Only Admin can update. Sets the price for a the volume which supersedes the scene price
	Price float32 `json:"price,omitempty"`

	// Free list of tag strings
	Tags []string `json:"tags,omitempty"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published,omitempty"`
}

// Validate --
func (s *Series) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	series := &Series{}

	// Find by slug when posting new series
	err := connection.Collection("series").FindOne(bson.M{"slug": s.Slug}, series)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
