/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

import (
	"github.com/go-bongo/bongo"
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

	// Media Performance
	Performance Performance `json:"performance"`

	// Volume this scene is in. Not all scenes have volumes
	Volume string `json:"volume"`

	// Some scenes can have no volumes but a series (best of/star profile)
	Series string `json:"series"`

	// List of Tags
	Tags []string `json:"tags"`

	// Read only value. Only Admin can update. Sets the price for a movie
	Price float32 `json:"price"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Scene) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	// scene := &Scene{}

	// // Find by slug when posting new scene
	// err := connection.Collection("scene").FindOne(bson.M{"slug": s.Slug}, scene)

	// if err == nil {
	// 	retval = append(retval, fmt.Errorf("This document is not unique"))
	// }
	return retval
}
