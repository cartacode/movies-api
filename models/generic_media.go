package models

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

// MediaInformation --
type MediaInformation struct {
	Director []string `json:"director"`

	Studio string `json:"studio"`

	// List of Mongo ObjectId for the Stars in this movie. Embeddable
	Stars []string `json:"Stars"`

	// Total movie length in seconds
	Length int32 `json:"length"`

	// List of available qualities for the video
	Quality []int `json:"quality"`
}

// Performance --
type Performance struct {

	// Calculated externally and maintained here
	Rank int32 `json:"rank"`

	// Calculated by user input. Only increases.
	Upvotes int32 `json:"upvotes"`

	// Calculated by user input. Only decreases.
	Downvotes int32 `json:"downvotes"`

	// Calculated by user view. Only increases.
	Views int32 `json:"views"`
}

// Chapter --
type Chapter struct {
	Name      string `json:"name"`
	Timestamp int    `json:"timestamp"`
}
