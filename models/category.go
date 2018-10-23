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

// Category Document
//
// A category can be associated with a scene, movie, volume, and series
//
// swagger:model category categoryDocument
type Category struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Slug for this category. lowercase and character stripped
	//
	// required: true
	Slug string `json:"slug"`

	// Description of this category if it has one. Not required
	// required: false
	Description string `json:"description,omitempty"`

	// Unique Title for this category
	//
	// required: true
	Title string `json:"title"`
}

// Validate --
func (s *Category) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	category := &Category{}

	// Find by slug when posting new category
	err := connection.Collection("category").FindOne(bson.M{"slug": s.Slug}, category)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
