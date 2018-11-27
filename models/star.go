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

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
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

	Studios []*bson.ObjectId `json:"studios"`

	Scenes []*bson.ObjectId `json:"scenes"`

	Movies []*bson.ObjectId `json:"movies"`

	Rank int `json:"rank"`

	// List of Tags
	Tags []string `json:"tags"`

	Social struct {
		Twitter string `json:"twitter"`

		Youtube string `json:"youtube"`

		Instagram string `json:"instagram"`

		Snapchat string `json:"snapchat"`
	} `json:"social"`

	Size struct {
		Weight int32 `json:"weight"`

		Waist int32 `json:"waist"`

		Bust string `json:"bust"`

		Height int32 `json:"height"`
	} `json:"size"`

	Images struct {
		Portrait string `json:"portrait"`

		Landscape string `json:"landscape"`

		Small string `json:"small"`
	} `json:"images"`

	Traits struct {
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
	// star := &Star{}

	// Find by slug when posting new star
	// err := connection.Collection("star").FindOne(bson.M{"slug": s.Slug}, star)

	// if err == nil {
	// 	retval = append(retval, fmt.Errorf("this document is not unique"))
	// }

	return retval
}

func (s *Star) addScene(id *bson.ObjectId) error {

	if !id.Valid() {
		return fmt.Errorf("not a valid bson id")
	}
	s.Scenes = append(s.Scenes, id)
	return nil

}
func (s *Star) removeScene(id *bson.ObjectId) error {

	if !id.Valid() {
		return fmt.Errorf("not a valid bson id")
	}

	for i := 0; i < len(s.Scenes); i++ {
		if s.Scenes[i] == id {
			copy(s.Scenes[i:], s.Scenes[i+1:])
			s.Scenes[len(s.Scenes)-1] = nil // or the zero vs.value of T
			s.Scenes = s.Scenes[:len(s.Scenes)-1]
			return nil

		}
	}

	return fmt.Errorf("bson not in slice")
}

func (s *Star) addMovie(id *bson.ObjectId) error {

	if id.Valid() {
		s.Movies = append(s.Movies, id)
		return nil
	}

	return fmt.Errorf("not a valid bson id")

}

func (s *Star) removeMovie(id *bson.ObjectId) error {

	if !id.Valid() {
		return fmt.Errorf("not a valid bson id")
	}

	for i := 0; i < len(s.Movies); i++ {
		if s.Movies[i] == id {
			copy(s.Movies[i:], s.Movies[i+1:])
			s.Movies[len(s.Movies)-1] = nil // or the zero vs.value of T
			s.Movies = s.Movies[:len(s.Movies)-1]
			return nil

		}
	}

	return fmt.Errorf("bson not in slice")

}

func (s *Star) addStudio(id *bson.ObjectId) error {

	if id.Valid() {
		s.Studios = append(s.Studios, id)
		return nil
	}

	return fmt.Errorf("not a valid bson id")

}

func (s *Star) removeStudio(id *bson.ObjectId) error {

	if !id.Valid() {
		return fmt.Errorf("not a valid bson id")
	}

	for i := 0; i < len(s.Studios); i++ {
		if s.Studios[i] == id {
			copy(s.Studios[i:], s.Studios[i+1:])
			s.Studios[len(s.Studios)-1] = nil // or the zero vs.value of T
			s.Studios = s.Studios[:len(s.Studios)-1]
			return nil

		}
	}

	return fmt.Errorf("bson not in slice")

}
