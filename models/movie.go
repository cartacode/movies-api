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

// Movie Document
//
// A Movie we can show. Contains all information and streaming data
//
// swagger:model
type Movie struct {
	bongo.DocumentBase `bson:",inline"`

	// Media information
	Images     Images     `json:"images"`
	Extras     []Extras   `json:"extras"`
	Thumbnails Thumbnails `json:"thumbnails"`
	Trailers   []Trailer  `json:"trailers"`

	Chapters []int `json:"chapters"`

	Quality []int `json:"quality"`

	// MovieInformation --
	Information MediaInformation `json:"information"`

	// Media Performance
	Performance Performance `json:"performance"`

	// List of Categories
	Category []string `json:"category"`

	// Unique Title for this movie
	Title string `json:"title"`

	// DynamoDBId
	DynamoDBId string `json:"dynamoId"`

	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	// Read only value. Only Admin can update. Sets the price for a movie
	Price float32 `json:"price"`

	// True/False. Has someone reviewed this movie
	Reviewed bool `json:"reviewed"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Movie) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	// movie := &Movie{}

	// Find by slug when posting new movie
	// err := connection.Collection("movie").FindOne(bson.M{"slug": s.Slug}, movie)

	// if err == nil {
	// retval = append(retval, fmt.Errorf("This document is not unique"))
	// }

	return retval
}
