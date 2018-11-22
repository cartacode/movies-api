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

	// List of Tags
	Tags []string `json:"tags"`

	// Volumes
	Volumes []string `json:"volumes"`

	Performance Performance `json:"performance"`

	// Unique Title for this movie
	Title string `json:"title"`

	// Description of this movie if it has one. Not required
	Description string `json:"description"`


	// True/False. Has someone reviewed this movie
	Reviewed bool `json:"reviewed"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
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
