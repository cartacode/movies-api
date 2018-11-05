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

	"github.com/VuliTv/go-movie-api/dbh"
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

	Rank int `json:"rank"`

	Social PerformerSocial `json:"social"`

	Slug string `json:"slug"`

	Gender string `json:"gender"`

	Size PerformerSize `json:"size"`

	Picture string `json:"picture"`

	Traits PerformerTraits `json:"traits"`

	Director bool `json:"director"`
}

// Validate --
func (s *Performer) Validate(*bongo.Collection) []error {
	connection, err := dbh.NewConnection("models.performer")
	if err != nil {
		panic(err)
	}
	retval := make([]error, 0)
	performer := &Performer{}

	// Find by slug when posting new performer
	if err := connection.Collection("performer").FindOne(bson.M{"slug": s.Slug}, performer); err != nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
