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
	"strconv"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Scene Document
//
// A scene can be associated with a volume, and series
//
// swagger:model
type Scene struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Title for this movie
	Title string `json:"title"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// DynamoDBId this will go away. Used right now to set media information
	DynamoDBId string `json:"dynamoId"`
	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	// Media information
	Images     Images     `json:"images"`
	Extras     []Extras   `json:"extras"`
	Thumbnails Thumbnails `json:"thumbnails"`
	Trailer    Trailer    `json:"trailer"`

	// MovieInformation --
	Information MediaInformation `json:"information"`

	Chapters []Chapter `json:"chapters"`

	// Media Performance
	Performance Performance `json:"performance"`

	// Volume this scene is in. Not all scenes have volumes
	Volume *bson.ObjectId `json:"volume"`

	// Some scenes can have no volumes but a series (best of/star profile)
	Series *bson.ObjectId `json:"series"`

	// List of Tags
	Tags []string `json:"tags"`

	// Read only value. Only Admin can update. Sets the price for a movie
	Price float64 `json:"price"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Scene) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	scene := &Scene{}

	// Check for series
	if s.Volume == nil {
		retval = append(retval, fmt.Errorf("volume cannot be empty"))

	} else {
		if !s.Volume.Valid() {
			retval = append(retval, fmt.Errorf("not a valid volume objectId"))
		}
	}

	// Check for studio
	if s.Information.Studio == nil {
		retval = append(retval, fmt.Errorf("studio cannot be empty"))

	} else {
		if !s.Information.Studio.Valid() {
			retval = append(retval, fmt.Errorf("not a valid studio objectId"))
		}
	}

	// Check for bad IDs
	for i, e := range s.Information.Director {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("director id is not valid in position: "+strconv.Itoa(i)))
		}
	}

	// Check for bad IDs
	for i, e := range s.Information.Stars {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("star id is not valid in position: "+strconv.Itoa(i)))
		}
	}
	// Find by slug when posting new scene
	err := connection.Collection("scene").FindOne(bson.M{"slug": s.Slug}, scene)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}
	return retval
}
