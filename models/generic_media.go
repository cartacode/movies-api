package models

// MediaContent --
type MediaContent struct {
	Thumbnails struct {
		Prefix string `json:"prefix"`
		Count  int    `json:"count"`
		Format string `json:"format"`
	} `json:"thumbnails"`
	Images struct {
		DetailPage            string `json:"detail_page"`
		TrayThumbnail         string `json:"tray_thumbnail"`
		TrayFeaturedThumbnail string `json:"tray_featured_thumbnail"`
		MobileThumbnail       string `json:"mobile_thumbnail"`
	} `json:"images"`
	Extras     []string `json:"extras"`
	Trailers   []string `json:"trailers"`
	CoverImage string   `json:"cover-image"`
}

// MediaInformation --
type MediaInformation struct {
	Director []string `json:"director"`

	Studio string `json:"studio"`

	// List of Mongo ObjectID for the Performers in this movie. Embeddable
	Performers []string `json:"performers"`

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
