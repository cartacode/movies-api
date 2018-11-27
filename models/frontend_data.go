package models

// FrontEndDataRequestResponse --
type FrontEndDataRequestResponse struct {
	Trending   []*Trending   `json:"trending"`
	JustForYou []*JustForYou `json:"just_for_you,omitempty"`
}

// Trending --
type Trending struct {
	Name           string `json:"name"`
	ImageURL       string `json:"imageUrl"`
	Year           string `json:"year"`
	Quality        int    `json:"quality"`
	Length         int32  `json:"length"`
	Description    string `json:"description"`
	TrailerURL     string `json:"trailerUrl"`
	NumberOfScenes int    `json:"number_of_scenes,omitempty"`
	ID             string `json:"_id"`
	VuliOriginal   bool   `json:"vuliOriginal"`
}

// JustForYou --
type JustForYou struct {
	Name           string `json:"name"`
	ImageURL       string `json:"imageUrl"`
	Year           string `json:"year"`
	Quality        int    `json:"quality"`
	Length         int32  `json:"length"`
	Description    string `json:"description"`
	TrailerURL     string `json:"trailerUrl"`
	NumberOfScenes int    `json:"number_of_scenes,omitempty"`
	ID             string `json:"_id"`
	VuliOriginal   bool   `json:"vuliOriginal"`
}
