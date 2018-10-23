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

// Studio --
type Studio struct {
	bongo.DocumentBase `bson:",inline"`
	// Public description of the studio
	Description string `json:"description,omitempty"`

	// Name of the studio
	Name string `json:"name,omitempty"`

	// Unique Slug for this studio. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug,omitempty"`
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
