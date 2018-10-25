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
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/go-bongo/bongo"
)

// Performer Document
//
// A document containing all information about a performer
//
// swagger:model
type Performer struct {
	bongo.DocumentBase `bson:",inline"`

	Bio string `json:"bio"`

	Name string `json:"name"`

	Birthdate time.Time `json:"birthdate"`

	Birthplace string `json:"birthplace"`

	Social *PerformerSocial `json:"social"`

	Slug string `json:"slug"`

	Gender string `json:"gender"`

	Size *PerformerSize `json:"size"`

	Traits *PerformerTraits `json:"traits"`
}

// Validate --
func (s *Performer) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	performer := &Performer{}

	// Find by slug when posting new performer
	err := connection.Collection("performer").FindOne(bson.M{"slug": s.Slug}, performer)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
