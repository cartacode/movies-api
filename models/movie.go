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
)

// Movie Document
//
// A Movie we can show. Contains all information and streaming data
//
// swagger:model movie movieDocument
type Movie struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Title for this movie
	Title string `json:"title"`

	// DynamoDBId
	DynamoDBId string `json:"dynamoId"`

	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// Chapter definitions for this movie
	Chapters []Chapter `json:"chapters"`

	// Media information
	Images     Images     `json:"images"`
	Extras     []Extras   `json:"extras"`
	Thumbnails Thumbnails `json:"thumbnails"`
	Trailer    Trailer    `json:"trailer"`

	// MovieInformation --
	Information MediaInformation `json:"information"`

	// Media Performance
	Performance Performance `json:"performance"`

	// List of Tags
	Tags []string `json:"tags"`

	// Read only value. Only Admin can update. Sets the price for a movie
	Price float64 `json:"price"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`
}

// Validate --
func (m *Movie) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	// movie := &Movie{}

	// Check for studio
	if m.Information.Studio == nil {
		retval = append(retval, fmt.Errorf("studio cannot be empty"))

	} else {
		if !m.Information.Studio.Valid() {
			retval = append(retval, fmt.Errorf("not a valid studio objectId"))
		}
	}

	// Check for bad IDs
	for i, e := range m.Information.Director {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("director id is not valid in position: "+strconv.Itoa(i)))
		}
	}

	// Check for bad IDs
	for i, e := range m.Information.Stars {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("star id is not valid in position: "+strconv.Itoa(i)))
		}
	}
	// Find by slug when posting new movie
	// err := connection.Collection("movie").FindOne(bson.M{"slug": m.Slug}, movie)

	// if err == nil {
	// retval = append(retval, fmt.Errorf("This document is not unique"))
	// }

	// s.Price = math.Ceil(s.Pr*100)/100
	log.Debugw("error saving volume", "error", retval)
	return retval
}
