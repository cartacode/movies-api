package models

// MediaContent --
type MediaContent struct {
	Images     []string `json:"images"`
	Extras     []string `json:"extras"`
	Trailers   []string `json:"trailers"`
	CoverImage string   `json:"cover-image"`
	
}
