package media

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
	Title     bool   `json:"title"`
	Published bool   `json:"published"`
	ImageURL  bool   `json:"imageUrl"`
}

// Thumbnails --
type Thumbnails struct {
	Prefix string `json:"prefix"`
	Count  int    `json:"count"`
	Format string `json:"format"`
}

func (t *Thumbnails) Upload() {

}

// Trailer --
type Trailer struct {
	DynamoDBId string `json:"dynamodbid"`
	Title      string `json:"title"`
}

// CLOUDFRONT .. Our cloudfront URL
var CLOUDFRONT = "https://cdn.vuli.tv"

// URL .. Get the url if available from the dynamoDbId
func (t *Trailer) URL() string {

	if t.DynamoDBId == "" {
		return ""
	}
	url := fmt.Sprintf("%s/%s/hls/%s.m3u8", CLOUDFRONT, t.DynamoDBId, t.Title)
	return url
}

// Information --
type Information struct {
	Director []*bson.ObjectId `json:"director"`

	Studio *bson.ObjectId `json:"studio"`

	// List of Mongo ObjectId for the Stars in this movie. Embeddable
	Stars []*bson.ObjectId `json:"stars"`

	// Total movie length in seconds
	Length int32 `json:"length"`

	// List of available qualities for the video
	Quality []int `json:"quality"`

	Year string `json:"year"`
}

// BestQuality .. Get the best quality video available from slice
func (m *Information) BestQuality() int {

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

// DynamoRecord --
type DynamoRecord struct {
	AbrBucket      string   `json:"abrBucket"`
	Dash           []int    `json:"dash"`
	DashPlaylist   string   `json:"dashPlaylist"`
	DashURL        string   `json:"dashUrl"`
	EcodeJobID     string   `json:"ecodeJobId"`
	EndTime        string   `json:"EndTime"`
	FrameCapture   bool     `json:"frameCapture"`
	FrameHeight    int      `json:"frameHeight"`
	FrameWdith     int      `json:"frameWdith"`
	GUID           string   `json:"guid"`
	Hls            []int    `json:"hls"`
	HlsPlaylist    string   `json:"hlsPlaylist"`
	HlsURL         string   `json:"hlsUrl"`
	Mp4            []int    `json:"mp4"`
	Mp4Bucket      string   `json:"mp4Bucket"`
	Mp4Outputs     []string `json:"mp4Outputs"`
	SrcBucket      string   `json:"srcBucket"`
	SrcHeight      int      `json:"srcHeight"`
	SrcMediainfo   string   `json:"srcMediainfo"`
	SrcVideo       string   `json:"srcVideo"`
	SrcWidth       int      `json:"srcWidth"`
	StartTime      string   `json:"startTime"`
	WorkflowStatus string   `json:"workflowStatus"`
}
