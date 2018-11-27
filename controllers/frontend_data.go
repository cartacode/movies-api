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

	movie := models.Scene{}
	list := results.Limit(20).Sort("performance.rank").Iter()

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

	}

	// Get auth user information
	var _, err = requests.GetAuthUser(r)
	if err == nil {
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

	js, err := json.Marshal(movieData)
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

	}

	// Get auth user information
	var _, err = requests.GetAuthUser(r)
	if err == nil {
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

	js, err := json.Marshal(sceneData)
	if err != nil {
		log.Warn(requests.ReturnAPIError(w, http.StatusBadRequest, err.Error()))
		return
	}
	requests.ReturnAPIOK(w, js)
}
