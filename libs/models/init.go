package models

import (
	"fmt"

	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"
	"github.com/VuliTv/go-movie-api/app/series"
	"github.com/VuliTv/go-movie-api/app/star"
	"github.com/VuliTv/go-movie-api/app/studio"
	"github.com/VuliTv/go-movie-api/app/volume"
	"github.com/VuliTv/go-movie-api/libs/logging"
)

var log = logging.GetProdLog()

// CLOUDFRONT .. Our cloudfront URL
var CLOUDFRONT = "https://cdn.vuli.tv"

// ModelByCollection --
func ModelByCollection(collection string) (interface{}, error) {

	switch collection {
	case "movie":
		model := &movie.Model{}
		return model, nil

	case "series":
		model := &series.Model{}
		return model, nil

	case "star":
		model := &star.Model{}
		return model, nil

	case "scene":
		model := &scene.Model{}
		return model, nil

	case "volume":
		model := &volume.Model{}
		return model, nil

	case "studio":
		model := &studio.Model{}
		return model, nil

	}
	err := fmt.Errorf("No collection found")
	return "", err
}
