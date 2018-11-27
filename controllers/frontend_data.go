package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VuliTv/go-movie-api/libs/requests"
	"github.com/VuliTv/go-movie-api/models"
)

// DataMovieTray ..
// fetches a customer profile from Authorize.net
func DataMovieTray(w http.ResponseWriter, r *http.Request) {

	// populate the just_for_you tray when logged in
	query := make(map[string]interface{})
	results := connection.Collection("movie").Collection().Find(query)

	movie := models.Movie{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	movieData := &models.FrontEndDataRequestResponse{}
	for list.Next(&movie) {
		trendingData := &models.Trending{}
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
			justForYouData := &models.JustForYou{}
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

	series := models.Movie{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	seriesData := &models.FrontEndDataRequestResponse{}
	for list.Next(&series) {
		trendingData := &models.Trending{}
		trendingData.Name = series.Title
		trendingData.ImageURL = series.Images.Landscape
		trendingData.Year = series.Information.Year
		trendingData.Quality = series.Information.BestQuality()
		trendingData.Length = series.Information.Length
		trendingData.Description = series.Description
		trendingData.TrailerURL = series.Trailer.URL()

		trendingData.ID = series.Id.Hex()

		// Add the data
		seriesData.Trending = append(seriesData.Trending, trendingData)
		if authUser.ObjectID != "" {
			justForYouData := &models.JustForYou{}
			justForYouData.Name = series.Title
			justForYouData.ImageURL = series.Images.Landscape
			justForYouData.Year = series.Information.Year
			justForYouData.Quality = series.Information.BestQuality()
			justForYouData.Length = series.Information.Length
			justForYouData.Description = series.Description
			justForYouData.TrailerURL = series.Trailer.URL()

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

	volume := models.Volume{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	volumeData := &models.FrontEndDataRequestResponse{}
	for list.Next(&volume) {
		trendingData := &models.Trending{}
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
			justForYouData := &models.JustForYou{}
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

	scene := models.Scene{}
	list := results.Limit(20).Sort("performance.rank").Iter()
	// Get auth user information
	authUser, _ := requests.GetAuthUser(r)

	sceneData := &models.FrontEndDataRequestResponse{}
	for list.Next(&scene) {
		trendingData := &models.Trending{}
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
			justForYouData := &models.JustForYou{}
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
