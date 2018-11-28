package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"
	"github.com/VuliTv/go-movie-api/app/series"
	"github.com/VuliTv/go-movie-api/app/star"
	"github.com/VuliTv/go-movie-api/app/volume"
	"github.com/VuliTv/go-movie-api/app/webdata"
	"github.com/VuliTv/go-movie-api/libs/requests"
)

// DataMovieTray ..
// fetches a customer profile from Authorize.net
func DataMovieTray(w http.ResponseWriter, r *http.Request) {

	// populate the just_for_you tray when logged in
	query := make(map[string]interface{})
	results := connection.Collection("movie").Collection().Find(query)

	movie := movie.Model{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	movieData := &webdata.FrontEndDataRequestResponse{}
	for list.Next(&movie) {
		trendingData := &webdata.Trending{}
		trendingData.Name = movie.Title
		trendingData.ImageURL = movie.Images.Landscape
		trendingData.Year = movie.Information.Year
		trendingData.Quality = movie.Information.BestQuality()
		trendingData.Length = movie.Information.Length
		trendingData.Description = movie.Description
		trendingData.TrailerURL = movie.Trailer.URL()

		trendingData.ID = movie.Id.Hex()

		// Add the data
		movieData.Trending = append(movieData.Trending, trendingData)
		if authUser.ObjectID != "" {
			justForYouData := &webdata.JustForYou{}
			justForYouData.Name = movie.Title
			justForYouData.ImageURL = movie.Images.Landscape
			justForYouData.Year = movie.Information.Year
			justForYouData.Quality = movie.Information.BestQuality()
			justForYouData.Length = movie.Information.Length
			justForYouData.Description = movie.Description
			justForYouData.TrailerURL = movie.Trailer.URL()

			justForYouData.ID = movie.Id.Hex()

			// Add the data
			movieData.JustForYou = append(movieData.JustForYou, justForYouData)
		}
	}

	js, err := json.Marshal(movieData)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}

// DataSeriesTray ..
// fetches a customer profile from Authorize.net
func DataSeriesTray(w http.ResponseWriter, r *http.Request) {

	// populate the just_for_you tray when logged in
	query := make(map[string]interface{})
	results := connection.Collection("series").Collection().Find(query)

	series := series.Model{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	seriesData := &webdata.FrontEndDataRequestResponse{}
	for list.Next(&series) {
		trendingData := &webdata.Trending{}
		trendingData.Name = series.Title
		trendingData.Description = series.Description

		trendingData.ID = series.Id.Hex()

		// Add the data
		seriesData.Trending = append(seriesData.Trending, trendingData)
		if authUser.ObjectID != "" {
			justForYouData := &webdata.JustForYou{}
			justForYouData.Name = series.Title
			justForYouData.Description = series.Description

			justForYouData.ID = series.Id.Hex()

			// Add the data
			seriesData.JustForYou = append(seriesData.JustForYou, justForYouData)
		}
	}

	js, err := json.Marshal(seriesData)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}

// DataVolumeTray ..
// fetches a customer profile from Authorize.net
func DataVolumeTray(w http.ResponseWriter, r *http.Request) {

	// populate the just_for_you tray when logged in
	query := make(map[string]interface{})
	results := connection.Collection("volume").Collection().Find(query)

	volume := volume.Model{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	volumeData := &webdata.FrontEndDataRequestResponse{}
	for list.Next(&volume) {
		trendingData := &webdata.Trending{}
		trendingData.Name = volume.Title
		trendingData.ImageURL = volume.Images.Landscape
		trendingData.Year = volume.Information.Year
		trendingData.Quality = volume.Information.BestQuality()
		trendingData.Length = volume.Information.Length
		trendingData.Description = volume.Description
		trendingData.TrailerURL = volume.Trailer.URL()
		trendingData.NumberOfScenes = len(volume.Scenes)
		trendingData.ID = volume.Id.Hex()

		// Add the data
		volumeData.Trending = append(volumeData.Trending, trendingData)
		if authUser.ObjectID != "" {
			justForYouData := &webdata.JustForYou{}
			justForYouData.Name = volume.Title
			justForYouData.ImageURL = volume.Images.Landscape
			justForYouData.Year = volume.Information.Year
			justForYouData.Quality = volume.Information.BestQuality()
			justForYouData.Length = volume.Information.Length
			justForYouData.Description = volume.Description
			justForYouData.TrailerURL = volume.Trailer.URL()
			justForYouData.NumberOfScenes = len(volume.Scenes)
			justForYouData.ID = volume.Id.Hex()

			// Add the data
			volumeData.JustForYou = append(volumeData.JustForYou, justForYouData)
		}
	}

	js, err := json.Marshal(volumeData)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}

// DataSceneTray ..
// fetches a customer profile from Authorize.net
func DataSceneTray(w http.ResponseWriter, r *http.Request) {

	// populate the just_for_you tray when logged in
	query := make(map[string]interface{})
	results := connection.Collection("scene").Collection().Find(query)

	scene := scene.Model{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	sceneData := &webdata.FrontEndDataRequestResponse{}
	for list.Next(&scene) {
		trendingData := &webdata.Trending{}
		trendingData.Name = scene.Title
		trendingData.ImageURL = scene.Images.Landscape
		trendingData.Year = scene.Information.Year
		trendingData.Quality = scene.Information.BestQuality()
		trendingData.Length = scene.Information.Length
		trendingData.Description = scene.Description
		trendingData.TrailerURL = scene.Trailer.URL()
		trendingData.ID = scene.Id.Hex()

		// Add the data
		sceneData.Trending = append(sceneData.Trending, trendingData)
		if authUser.ObjectID != "" {
			justForYouData := &webdata.JustForYou{}
			justForYouData.Name = scene.Title
			justForYouData.ImageURL = scene.Images.Landscape
			justForYouData.Year = scene.Information.Year
			justForYouData.Quality = scene.Information.BestQuality()
			justForYouData.Length = scene.Information.Length
			justForYouData.Description = scene.Description
			justForYouData.TrailerURL = scene.Trailer.URL()
			justForYouData.ID = scene.Id.Hex()

			// Add the data
			sceneData.JustForYou = append(sceneData.JustForYou, justForYouData)
		}
	}

	js, err := json.Marshal(sceneData)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}

// DataStarTray ..
// fetches a customer profile from Authorize.net
func DataStarTray(w http.ResponseWriter, r *http.Request) {

	// populate the just_for_you tray when logged in
	query := make(map[string]interface{})
	results := connection.Collection("star").Collection().Find(query)

	star := star.Model{}
	list := results.Limit(20).Sort("-performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	starData := &webdata.FrontEndDataRequestResponse{}
	for list.Next(&star) {
		trendingData := &webdata.Trending{}
		trendingData.Name = star.Name
		trendingData.ImageURL = star.Images.Portrait
		trendingData.TagLine = star.Tagline
		trendingData.ID = star.Id.Hex()

		// Add the data
		starData.Trending = append(starData.Trending, trendingData)
		if authUser.ObjectID != "" {
			justForYouData := &webdata.JustForYou{}
			justForYouData.Name = star.Name
			justForYouData.ImageURL = star.Images.Portrait
			justForYouData.TagLine = star.Tagline
			justForYouData.ID = star.Id.Hex()

			// Add the data
			starData.JustForYou = append(starData.JustForYou, justForYouData)
		}
	}

	js, err := json.Marshal(starData)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}
