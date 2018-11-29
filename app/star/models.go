/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package star

import (
	"fmt"
	"time"

	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/VuliTv/go-movie-api/libs/stringops"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// FunctionEnum --
var FunctionEnum = []string{"performer", "director", "director-performer"}

// Model Document
//
// A document containing all information about a Model
//
// swagger:model
type Model struct {
	bongo.DocumentBase `bson:",inline"`

	Name string `json:"name"`

	Slug string `json:"slug"`

	Tagline string `json:"tagline"`

	Bio string `json:"bio"`

	Gender string `json:"gender"`

	Birthdate time.Time `json:"birthdate"`

	Birthplace string `json:"birthplace"`

	// Media Performance
	Performance media.Performance `json:"performance"`

	Studios []*bson.ObjectId `json:"studios"`

	Scenes []*bson.ObjectId `json:"scenes"`

	Movies []*bson.ObjectId `json:"movies"`

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
		ModelSign string `json:"sign"`
	} `json:"traits"`

	Function string `json:"function"`
}

// Validate --
func (s *Model) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)

	// Enum function check
	if !stringops.StringInSlice(s.Function, FunctionEnum) {
		retval = append(retval, fmt.Errorf("function must be in %s", FunctionEnum))
	}
	// star := &Model{}

	// Find by slug when posting new star
	// err := connection.Collection("star").FindOne(bson.M{"slug": s.Slug}, star)

	// if err == nil {
	// 	retval = append(retval, fmt.Errorf("this document is not unique"))
	// }

	return retval
}

func (s *Model) addScene(id *bson.ObjectId) error {

	if !id.Valid() {
		return fmt.Errorf("not a valid bson id")
	}
	s.Scenes = append(s.Scenes, id)
	return nil

}
func (s *Model) removeScene(id *bson.ObjectId) error {

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

func (s *Model) addMovie(id *bson.ObjectId) error {

	if id.Valid() {
		s.Movies = append(s.Movies, id)
		return nil
	}

	return fmt.Errorf("not a valid bson id")

}

func (s *Model) removeMovie(id *bson.ObjectId) error {

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

func (s *Model) addStudio(id *bson.ObjectId) error {

	if id.Valid() {
		s.Studios = append(s.Studios, id)
		return nil
	}

	return fmt.Errorf("not a valid bson id")

}

func (s *Model) removeStudio(id *bson.ObjectId) error {

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
