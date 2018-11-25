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

// Star Document
//
// A document containing all information about a Star
//
// swagger:model
type Star struct {
	bongo.DocumentBase `bson:",inline"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Tagline string `json:"tagline"`

	Bio string `json:"bio"`

	Gender string `json:"gender"`

	Birthdate time.Time `json:"birthdate"`

	Birthplace string `json:"birthplace"`

	Favorites int `json:"favorites"`

	Likes int32 `json:"likes"`

	Dislikes int32 `json:"dislikes"`

	Studios []string `json:"studios"`

	Scenes []string `json:"scenes"`

	Movies []string `json:"movies"`

	Rank int `json:"rank"`

	// List of Tags
	Tags []string `json:"tags"`

	Social struct {
		Twitter string `json:"twitter"`

		Youtube string `json:"youtube"`

		Instagram string `json:"instagram"`

		Snapchat string `json:"snapchat"`
	} `json:"social"`

	StarSize struct {
		Weight int32 `json:"weight"`

		Waist int32 `json:"waist"`

		Bust string `json:"bust"`

		Height int32 `json:"height"`
	} `json:"size"`

	Images struct {
		Portrait string `json:"portrait"`

		Landscape string `json:"landscape"`
	} `json:"images"`

	StarTraits struct {
		Ethnicity string `json:"ethnicity"`
		HairColor string `json:"haircolor"`
		Piercings bool   `json:"piercings"`
		Tattoos   bool   `json:"tattoos"`
		StarSign  string `json:"sign"`
	} `json:"traits"`

	Director bool `json:"director"`
}

// Validate --
func (s *Star) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	star := &Star{}

	// Find by slug when posting new star
	err := connection.Collection("star").FindOne(bson.M{"slug": s.Slug}, star)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
