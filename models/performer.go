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

// Performer --
type Performer struct {
	bongo.DocumentBase `bson:",inline"`

	Bio string `json:"bio,omitempty"`

	Name string `json:"name,omitempty"`

	Birthdate time.Time `json:"birthdate,omitempty"`

	Birthplace string `json:"birthplace,omitempty"`

	Social *PerformerSocial `json:"social,omitempty"`

	Slug string `json:"slug,omitempty"`

	Gender string `json:"gender,omitempty"`

	Size *PerformerSize `json:"size,omitempty"`

	Traits *PerformerTraits `json:"traits,omitempty"`
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
