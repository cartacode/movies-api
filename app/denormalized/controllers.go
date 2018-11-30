package denormalized

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/VuliTv/go-movie-api/app/movie"
	"github.com/VuliTv/go-movie-api/app/scene"
	"github.com/VuliTv/go-movie-api/libs/requests"
)

// Scenes --
func Scenes(w http.ResponseWriter, r *http.Request) {

	var retval []interface{}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(r.URL.String())))
	data := redisHandler.Get(hash)
	if data.Err() == nil {
		resp := data.Val()
		log.Info("Loaded response from cache")
		requests.ReturnAPIOK(w, []byte(resp))
		return
	}
	results := mongoHandler.Collection("scene").Find(nil)
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
		if scene.Volume != nil {
			vol := &ModelStub{}
			if err := mongoHandler.Collection("volume").FindById(*scene.Volume, &vol); err != nil {
				dScene.Volume = nil
				log.Warn(err)

			}
			dScene.Volume = vol
		}

		//Series
		if scene.Volume != nil {
			series := &ModelStub{}
			if err := mongoHandler.Collection("series").FindById(*scene.Volume, &series); err != nil {
				log.Warn(err)
				dScene.Series = nil
			} else {

				dScene.Series = series
			}
		}

		// Studio
		if scene.Information.Studio != nil {
			studio := &ModelStub{}
			if err := mongoHandler.Collection("studio").FindById(*scene.Information.Studio, &studio); err != nil {
				log.Warn(err)
				dScene.Information.Studio = nil
			} else {

				dScene.Information.Studio = studio
			}
		}

		// Stars
		if scene.Information.Stars != nil {
			for _, star := range scene.Information.Stars {
				starStub := &ModelStub{}
				if err := mongoHandler.Collection("star").FindById(*star, &starStub); err != nil {
					log.Warn(err)
				} else {
					dScene.Information.Stars = append(dScene.Information.Stars, starStub)
				}
			}

		}

		if scene.Information.Director != nil {
			for _, director := range scene.Information.Director {
				directorStub := &ModelStub{}
				if err := mongoHandler.Collection("star").FindById(*director, &directorStub); err != nil {
					log.Warn(err)
				} else {
					dScene.Information.Director = append(dScene.Information.Director, directorStub)
				}
			}

		}
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
