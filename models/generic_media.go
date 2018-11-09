package models

// Images --
type Images struct {
	Detailpage            string   `json:"detailpage"`
	TrayThumbnail         string   `json:"traythumbnail"`
	TrayFeaturedThumbnail string   `json:"trayfeaturedthumbnail"`
	MobileThumbnail       string   `json:"mobilethumbnail"`
	CoverImage            string   `json:"coverimage"`
	BackgroundImage       string   `json:"backgroundimage"`
	Available             []string `json:"available"`
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
	URL       string `json:"url"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}

// MediaInformation --
type MediaInformation struct {
	Director []string `json:"director"`

	Studio string `json:"studio"`

	// List of Mongo ObjectID for the Stars in this movie. Embeddable
	Stars []string `json:"Stars"`

	// Total movie length in seconds
	Length int32 `json:"length"`
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
