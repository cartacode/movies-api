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

// Category --
type Category struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Slug for this category. lowercase and character stripped
	Slug string `json:"slug"`

	// Description of this category if it has one. Not required
	Description string `json:"description,omitempty"`

	// Unique Title for this category
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
