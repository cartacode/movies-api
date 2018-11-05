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

// Star Document
//
// A document containing all information about a Star
//
// swagger:model
type Star struct {
	bongo.DocumentBase `bson:",inline"`

	Bio string `json:"bio"`

	Name string `json:"name"`

	Birthdate time.Time `json:"birthdate"`

	Birthplace string `json:"birthplace"`

	Rank int `json:"rank"`

	Social StarSocial `json:"social"`

	Slug string `json:"slug"`

	Gender string `json:"gender"`

	Size StarSize `json:"size"`

	Picture string `json:"picture"`

	Traits StarTraits `json:"traits"`

	Director bool `json:"director"`
}

// Validate --
func (s *Star) Validate(*bongo.Collection) []error {
	connection, err := dbh.NewConnection("models.Star")
	if err != nil {
		panic(err)
	}
	retval := make([]error, 0)
	Star := &Star{}

	// Find by slug when posting new Star
	if err := connection.Collection("Star").FindOne(bson.M{"slug": s.Slug}, Star); err != nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}

// StarTraits --
type StarTraits struct {
	Ethnicity string `json:"ethnicity"`
	HairColor string `json:"haircolor"`
	Piercings bool   `json:"piercings"`
	Tattoos   bool   `json:"tattoos"`
	StarSign  string `json:"sign"`
}

// StarSocial --
type StarSocial struct {
	Twitter string `json:"twitter"`

	Youtube string `json:"youtube"`

	Instagram string `json:"instagram"`

	Snapchat string `json:"snapchat"`
}

// StarSize --
type StarSize struct {
	Weight int32 `json:"weight"`

	Waist int32 `json:"waist"`

	Bust string `json:"bust"`

	Height int32 `json:"height"`
}
