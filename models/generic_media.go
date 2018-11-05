package models

// MediaContent --
type MediaContent struct {
	Images struct {
		Detailpage            string `json:"detailpage"`
		Traythumbnail         string `json:"traythumbnail"`
		Trayfeaturedthumbnail string `json:"trayfeaturedthumbnail"`
		Mobilethumbnail       string `json:"mobilethumbnail"`
		Coverimage            struct {
			Selected  string   `json:"selected"`
			Available []string `json:"available"`
		} `json:"coverimage"`
	} `json:"images"`
	Extras []struct {
		URL       string `json:"url"`
		Published bool   `json:"published"`
	} `json:"extras"`
	Thumbnails struct {
		Prefix string `json:"prefix"`
		Count  int    `json:"count"`
		Format string `json:"format"`
	} `json:"thumbnails"`
}

// Trailer --
type Trailer struct {
	Title  string `json:"title"`
	Length int    `json:"length"`
	Image  string `json:"image"`
	Path   string `json:"path"`
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
