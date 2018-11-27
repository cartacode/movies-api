package models

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

// Images --
type Images struct {
	Landscape  string `json:"landscape"`
	Portrait   string `json:"portrait"`
	Banner     string `json:"banner"`
	DetailPage string `json:"detailpage"`
}

// Extras --
type Extras struct {
	URL       string `json:"url"`
	Published bool   `json:"published"`
}

// Thumbnails --
type Thumbnails struct {
	Prefix string `json:"prefix"`
	Count  int    `json:"count"`
	Format string `json:"format"`
}

// Trailer --
type Trailer struct {
	DynamoDBId string `json:"dynamodbid"`
	Title      string `json:"title"`
}

// URL .. Get the url if available from the dynamoDbId
func (t *Trailer) URL() string {

	if t.DynamoDBId == "" {
		return ""
	}
	url := fmt.Sprintf("%s/%s/hls/%s.m3u8", CLOUDFRONT, t.DynamoDBId, t.Title)
	return url
}

// MediaInformation --
type MediaInformation struct {
	Director []*bson.ObjectId `json:"director"`

	Studio *bson.ObjectId `json:"studio"`

	// List of Mongo ObjectId for the Stars in this movie. Embeddable
	Stars []*bson.ObjectId `json:"Stars"`

	// Total movie length in seconds
	Length int32 `json:"length"`

	// List of available qualities for the video
	Quality []int `json:"quality"`

	Year string `json:"year"`
}

// BestQuality .. Get the best quality video available from slice
func (m *MediaInformation) BestQuality() int {

	highest := 480
	for i := range m.Quality {
		if i > highest {
			highest = i
		}
	}
	return highest
}

// Performance --
type Performance struct {

	// Calculated externally and maintained here
	Rank int32 `json:"rank"`

	// Calculated by user input. Only increases.
	Upvotes int32 `json:"upvotes"`

	// Calculated by user input. Only decreases.
	Downvotes int32 `json:"downvotes"`

	// Calculated by user input. Only decreases.
	Favorites int64 `json:"favorites"`

	// Calculated by user view. Only increases.
	Views int32 `json:"views"`
}

// Chapter --
type Chapter struct {
	Name      string `json:"name"`
	Timestamp int    `json:"timestamp"`
}
