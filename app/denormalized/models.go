package denormalized

import (
	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"
	"github.com/VuliTv/go-movie-api/app/star"
	"github.com/VuliTv/go-movie-api/app/volume"
	"gopkg.in/mgo.v2/bson"
)

type Scene struct {
	scene.Model `json:",inline"`
	Volume      *ModelStub `json:"volume"`
	Series      *ModelStub `json:"series"`
	// Information --
	Information Information `json:"information"`
}

type Movie struct {
	movie.Model `json:",inline"`

	// Information --
	Information Information `json:"information"`
}

type Star struct {
	star.Model `json:",inline"`

	Studios []*ModelStub `json:"studios"`

	Scenes []*ModelStub `json:"scenes"`

	Movies []*ModelStub `json:"movies"`

	Volumes []*ModelStub `json:"volumes"`
}

type Volume struct {
	volume.Model `json:",inline"`

	Series *ModelStub `json:"series"`

	Scenes []*ModelStub `json:"scenes"`

	Information Information `json:"information"`
}

type Information struct {
	media.Information
	Director []*ModelStub `json:"director"`

	Studio *ModelStub `json:"studio"`

	// List of Mongo ObjectId for the Stars in this movie. Embeddable
	Stars []*ModelStub `json:"stars"`

	// Total movie length in seconds
	Length int32 `json:"length"`

	// List of available qualities for the video
	Quality []int `json:"quality"`

	Year string `json:"year"`
}
type ModelStub struct {
	Id     *bson.ObjectId `json:"_id"`
	Title  string         `json:"title,omitempty"`
	Name   string         `json:"name,omitempty"`
	Slug   string         `json:"slug,omitempty"`
	Images media.Images   `json:"images"`
}
