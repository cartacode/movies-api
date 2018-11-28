/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

import (
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Series Document
//
// A series can be associated with N volumes
//
// swagger:model
type Series struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Title for this movie
	Title string `json:"title"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// List of Tags
	Tags []string `json:"tags"`

	// Volumes this series contains
	Volumes []*bson.ObjectId `json:"volumes"`

	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Series) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	// series := &Series{}

	// Find by slug when posting new series
	// err := connection.Collection("series").FindOne(bson.M{"slug": s.Slug}, series)

	// if err == nil {
	// 	retval = append(retval, fmt.Errorf("This document is not unique"))
	// }

	return retval
}
