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

// Studio Document
//
// A studio can be associated with a Series, Volume, or Movie
//
// swagger:model
type Studio struct {
	bongo.DocumentBase `bson:",inline"`
	// Public description of the studio
	Description string `json:"description"`

	// Name of the studio
	Name string `json:"name"`

	// Unique Slug for this studio. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`
}

// Validate --
func (s *Studio) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	studio := &Studio{}

	// Find by slug when posting new studio
	err := connection.Collection("studio").FindOne(bson.M{"slug": s.Slug}, studio)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
