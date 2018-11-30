package denormalized

import (
	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/VuliTv/go-movie-api/app/scene"
	"gopkg.in/mgo.v2/bson"
)

type Scene struct {
	scene.Model `json:",inline"`
	Volume      *ModelStub `json:"volume"`
	Series      *ModelStub `json:"series"`
	// Information --
	Information struct {
		Director []*ModelStub `json:"director"`

		Studio *ModelStub `json:"studio"`

		// List of Mongo ObjectId for the Stars in this movie. Embeddable
		Stars []*ModelStub `json:"stars"`

		// Total movie length in seconds
		Length int32 `json:"length"`

		// List of available qualities for the video
		Quality []int `json:"quality"`

		Year string `json:"year"`
	} `json:"information"`
}

type ModelStub struct {
	ID     bson.ObjectId `json:"_id"`
	Title  string        `json:"title,omitempty"`
	Name   string        `json:"name,omitempty"`
	Slug   string        `json:"slug,omitempty"`
	Images media.Images  `json:"images"`
}
