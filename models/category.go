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
	Description string `json:"description"`

	// Unique Name for this category
	//
	// required: true
	Name string `json:"name"`
}

// Validate --
func (s *Category) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	category := &Category{}

	if s.Name == "" || s.Slug == "" {
		retval = append(retval, fmt.Errorf("Name cannot be empty"))
		return retval
	}
	// Find by slug when posting new category
	err := connection.Collection("category").FindOne(bson.M{"slug": s.Slug}, category)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
