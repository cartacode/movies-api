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

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Movie Document
//
// A Movie we can show. Contains all information and streaming data
//
// swagger:model
type Movie struct {
	bongo.DocumentBase `bson:",inline"`
	MediaContent       `bson:",inline"`
	// List of Categories
	Category []string `json:"category"`

	// MovieInformation --
	Information struct {
		Director []string `json:"director"`

		Studio string `json:"studio"`

		// List of Mongo ObjectID for the Performers in this movie. Embeddable
		Performers []string `json:"performers"`
	} `json:"information"`

	// DynamoDBId

	DynamoDBId string `json:"dynamoId"`
	// Ordered list of movie location(s)
	Playlist []string `json:"playlist"`

	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	// Calculated by user view. Only increases.
	Views int32 `json:"views"`

	// Read only value. Only Admin can update. Sets the price for a movie
	Price float32 `json:"price"`

	// Unique Title for this movie
	Title string `json:"title"`

	// Calculated externally and maintained here
	Rank int32 `json:"rank"`

	// True/False. Has someone reviewed this movie
	Reviewed bool `json:"reviewed"`

	// Total movie length in seconds
	Length int32 `json:"length"`

	// Calculated by user input. Only increases.
	Upvotes int32 `json:"upvotes"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// Calculated by user input. Only decreases.
	Downvotes int32 `json:"downvotes"`

	// Free list of tag strings
	Tags []string `json:"tags"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (s *Movie) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	movie := &Movie{}

	// Find by slug when posting new movie
	err := connection.Collection("movie").FindOne(bson.M{"slug": s.Slug}, movie)

	if err == nil {
		retval = append(retval, fmt.Errorf("This document is not unique"))
	}

	return retval
}
