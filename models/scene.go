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

// Scene Document
//
// A scene can be associated with a volume, and series
//
// swagger:model
type Scene struct {
	bongo.DocumentBase `bson:",inline"`

	// Media information
	MediaContent MediaContent `bson:",inline" json:"media"`

	// MovieInformation --
	Information MediaInformation `json:"information"`

	// Media Performance
	Performance Performance `json:"performance"`

	// Unique Title for this movie
	Title string `json:"title"`

	// DynamoDBId
	DynamoDBId string `json:"dynamoId"`

	// List of Categories
	Category []string `json:"category"`

	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	Volume string `json:"volume"`
	// Read only value. Only Admin can update. Sets the price for a movie
	Price float32 `json:"price"`

	// True/False. Has someone reviewed this movie
	Reviewed bool `json:"reviewed"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	Series string `json:"series"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Scene) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	scene := &Scene{}

	// Find by slug when posting new scene
	err := connection.Collection("scene").FindOne(bson.M{"slug": s.Slug}, scene)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}
	return retval
}
