package denormalized

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"
	"github.com/VuliTv/go-movie-api/app/star"
	"github.com/VuliTv/go-movie-api/app/volume"
	"github.com/VuliTv/go-movie-api/libs/requests"
	"gopkg.in/mgo.v2/bson"
)

// Scenes --
func Scenes(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}

	query, err := requests.QuerySanatizer(r.URL.Query())
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}
	log.Debugw("query running", "Q", query)

	hash := fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	data := redisHandler.Get(hash)
	if data.Err() == nil {
		resp := data.Val()
		log.Info("Loaded response from cache")
		requests.ReturnAPIOK(w, []byte(resp))
		return
	}
	results := mongoHandler.Collection("scene").Find(query)
	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	pagination, err := results.Paginate(perPage, page)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	scene := &scene.Model{}

	// Add the found results
	for results.Next(&scene) {

		dScene := &Scene{Model: *scene}

		// Volume
		dScene.Volume = addDenormalizedData("volume", scene.Volume)
		dScene.Series = addDenormalizedData("volume", scene.Series)
		dScene.Information.Studio = addDenormalizedData("volume", scene.Information.Studio)

		// Stars
		dScene.Information.Stars = addDenormalizedDataFromSlice("star", scene.Information.Stars)
		dScene.Information.Director = addDenormalizedDataFromSlice("star", scene.Information.Director)

		dScene.Information.Length = scene.Information.Length
		dScene.Information.Quality = scene.Information.Quality
		dScene.Information.Year = scene.Information.Year

		// Append all
		retval = append(retval, dScene)
	}
	response := requests.JSONPaginationResponse{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}
	log.Info("Loaded response from db")
	hash = fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	if ok := redisHandler.Set(hash, rs, time.Second*10); ok != nil {
		log.Info("set response to cache")
	}

	requests.ReturnAPIOK(w, rs)
}

// Stars --
func Stars(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}

	query, err := requests.QuerySanatizer(r.URL.Query())
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}
	log.Debugw("query running", "Q", query)

	hash := fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	data := redisHandler.Get(hash)
	if data.Err() == nil {
		resp := data.Val()
		log.Info("Loaded response from cache")
		requests.ReturnAPIOK(w, []byte(resp))
		return
	}
	results := mongoHandler.Collection("star").Find(query)
	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	pagination, err := results.Paginate(perPage, page)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	star := &star.Model{}

	// Add the found results
	for results.Next(&star) {

		dStar := &Star{Model: *star}

		dStar.Studios = addDenormalizedDataFromSlice("studio", star.Studios)
		dStar.Scenes = addDenormalizedDataFromSlice("scene", star.Scenes)
		dStar.Movies = addDenormalizedDataFromSlice("movie", star.Movies)
		dStar.Volumes = addDenormalizedDataFromSlice("volume", star.Volumes)

		// Append all
		retval = append(retval, dStar)
	}
	response := requests.JSONPaginationResponse{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}
	log.Info("Loaded response from db")
	hash = fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	if ok := redisHandler.Set(hash, rs, time.Second*10); ok != nil {
		log.Info("set response to cache")
	}

	requests.ReturnAPIOK(w, rs)
}

// Movies --
func Movies(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	data := redisHandler.Get(hash)
	if data.Err() == nil {
		resp := data.Val()
		log.Info("Loaded response from cache")
		requests.ReturnAPIOK(w, []byte(resp))
		return
	}
	results := mongoHandler.Collection("movie").Find(nil)
	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	pagination, err := results.Paginate(perPage, page)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	movie := &movie.Model{}

	// Add the found results
	for results.Next(&movie) {

		dMovie := &Movie{Model: *movie}

		// Studio
		if movie.Information.Studio != nil {
			studio := &ModelStub{}
			if err := mongoHandler.Collection("studio").FindById(*movie.Information.Studio, &studio); err != nil {
				log.Warn(err)
				dMovie.Information.Studio = nil
			} else {

				dMovie.Information.Studio = studio
			}
		}

		// Stars
		if movie.Information.Stars != nil {
			for _, star := range movie.Information.Stars {
				starStub := &ModelStub{}
				if err := mongoHandler.Collection("star").FindById(*star, &starStub); err != nil {
					log.Warn(err)
				} else {
					dMovie.Information.Stars = append(dMovie.Information.Stars, starStub)
				}
			}

		}

		if movie.Information.Director != nil {
			for _, director := range movie.Information.Director {
				directorStub := &ModelStub{}
				if err := mongoHandler.Collection("star").FindById(*director, &directorStub); err != nil {
					log.Warn(err)
				} else {
					dMovie.Information.Director = append(dMovie.Information.Director, directorStub)
				}
			}

		}
		dMovie.Information.Length = movie.Information.Length
		dMovie.Information.Quality = movie.Information.Quality
		dMovie.Information.Year = movie.Information.Year

		// Append all
		retval = append(retval, dMovie)
	}
	response := requests.JSONPaginationResponse{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}
	log.Info("Loaded response from db")
	hash = fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	if ok := redisHandler.Set(hash, rs, time.Second*10); ok != nil {
		log.Info("set response to cache")
	}

	requests.ReturnAPIOK(w, rs)
}

func Volumes(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}

	query, err := requests.QuerySanatizer(r.URL.Query())
	if err != nil {
		log.Error(requests.ReturnAPIError(w, http.StatusInternalServerError, err.Error()))
		return
	}
	log.Debugw("query running", "Q", query)

	hash := fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	data := redisHandler.Get(hash)
	if data.Err() == nil {
		resp := data.Val()
		log.Info("Loaded response from cache")
		requests.ReturnAPIOK(w, []byte(resp))
		return
	}
	results := mongoHandler.Collection("volume").Find(query)
	// Get pagination information
	perPage, page := requests.GetPaginationInfo(r)
	pagination, err := results.Paginate(perPage, page)

	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}

	// Get which page we are on to skip
	// results.Query.Skip(page * perpage)

	volume := &volume.Model{}

	// Add the found results
	for results.Next(&volume) {

		dVolume := &Volume{Model: *volume}

		// Volume
		dVolume.Scenes = addDenormalizedDataFromSlice("scene", volume.Scenes)
		dVolume.Series = addDenormalizedData("volume", volume.Series)
		dVolume.Information.Studio = addDenormalizedData("volume", volume.Information.Studio)

		// Stars
		dVolume.Information.Stars = addDenormalizedDataFromSlice("star", volume.Information.Stars)
		dVolume.Information.Director = addDenormalizedDataFromSlice("star", volume.Information.Director)

		dVolume.Information.Length = volume.Information.Length
		dVolume.Information.Quality = volume.Information.Quality
		dVolume.Information.Year = volume.Information.Year

		// Append all
		retval = append(retval, dVolume)
	}
	response := requests.JSONPaginationResponse{
		Results:       retval,
		TotalResults:  pagination.TotalRecords,
		RecordsOnPage: pagination.RecordsOnPage,
		Page:          pagination.Current,
		TotalPages:    pagination.TotalPages,
	}

	// Turn it into a json and serve it up
	rs, err := json.Marshal(response)
	if err != nil {
		requests.ReturnAPIError(w, http.StatusBadRequest, err.Error())
		log.Error(err)
		return
	}
	log.Info("Loaded response from db")
	hash = fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	if ok := redisHandler.Set(hash, rs, time.Second*10); ok != nil {
		log.Info("set response to cache")
	}

	requests.ReturnAPIOK(w, rs)
}
func addDenormalizedDataFromSlice(collection string, objectIDS []*bson.ObjectId) []*ModelStub {

	if objectIDS == nil {
		return nil
	}
	retval := []*ModelStub{}
	for _, object := range objectIDS {
		objectStub := &ModelStub{}
		if err := mongoHandler.Collection(collection).FindById(*object, &objectStub); err != nil {
			log.Warn(err)
		} else {
			objectStub.Id = object
			retval = append(retval, objectStub)
		}
	}

	return retval
}

func addDenormalizedData(collection string, objectId *bson.ObjectId) *ModelStub {

	if objectId == nil {
		return nil
	}
	retval := &ModelStub{}
	if err := mongoHandler.Collection(collection).FindById(*objectId, &retval); err != nil {
		log.Warn(err)

	}

	return retval
}
