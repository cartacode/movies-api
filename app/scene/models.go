/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package scene

import (
	"fmt"
	"strconv"

	"github.com/VuliTv/go-movie-api/app/media"
	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Model Document
//
// A scene can be associated with a volume, and series
//
// swagger:model
type Model struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Title for this movie
	Title string `json:"title"`

	// Unique Slug for this movie. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// DynamoDBId this will go away. Used right now to set media information
	DynamoDBId string `json:"dynamoId"`
	// Description of this movie if it has one. Not required
	Description string `json:"description"`

	// Media information
	Images     media.Images     `json:"images"`
	Extras     []media.Extras   `json:"extras"`
	Thumbnails media.Thumbnails `json:"thumbnails"`
	Trailer    media.Trailer    `json:"trailer"`

	// MovieInformation --
	Information media.MediaInformation `json:"information"`

	Chapters []media.Chapter `json:"chapters"`

	// Media Performance
	Performance media.Performance `json:"performance"`

	// Volume this scene is in. Not all scenes have volumes
	Volume *bson.ObjectId `json:"volume"`

	// Some scenes can have no volumes but a series (best of/star profile)
	Series *bson.ObjectId `json:"series"`

	// List of Tags
	Tags []string `json:"tags"`

	// Read only value. Only Admin can update. Sets the price for a movie
	Price float64 `json:"price"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`

	diffTracker *bongo.DiffTracker
}

// GetDiffTracker ..
func (s *Model) GetDiffTracker() *bongo.DiffTracker {
	if s.diffTracker == nil {
		s.diffTracker = bongo.NewDiffTracker(s)
	}

	return s.diffTracker
}

// Validate --
func (s *Model) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	// scene := &Model{}

	// Check for series
	if s.Volume == nil {
		retval = append(retval, fmt.Errorf("volume cannot be empty"))

	} else {
		if !s.Volume.Valid() {
			retval = append(retval, fmt.Errorf("not a valid volume objectId"))
		}
	}

	// Check for studio
	if s.Information.Studio == nil {
		retval = append(retval, fmt.Errorf("studio cannot be empty"))

	} else {
		if !s.Information.Studio.Valid() {
			retval = append(retval, fmt.Errorf("not a valid studio objectId"))
		}
	}

	// Check for bad IDs
	for i, e := range s.Information.Director {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("director id is not valid in position: "+strconv.Itoa(i)))
		}
	}

	// Check for bad IDs
	for i, e := range s.Information.Stars {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("star id is not valid in position: "+strconv.Itoa(i)))
		}
	}
	// Find by slug when posting new scene
	// err := connection.Collection("scene").FindOne(bson.M{"slug": s.Slug}, scene)

	// if err == nil {
	// 	retval = append(retval, fmt.Errorf("This document is not unique"))
	// }
	return retval
}

// AfterSave ..
func (s *Model) AfterSave(*bongo.Collection) error {

	/*  VOLUME CASCADE */
	if err := connection.Collection("volume").Collection().Update(bson.M{"_id": s.Volume}, bson.M{"$push": bson.M{"scenes": s.Id}}); err != nil {
		log.Errorw("cascade failure on volume add sceneId",
			"volume_id", s.Volume,
			"scene_id", s.Id,
			"error", err,
		)
		return fmt.Errorf("cascade failure")
	}

	/*  Stars CASCADE */
	if s.Information.Stars != nil {
		for i, starID := range s.Information.Stars {
			if err := connection.Collection("star").Collection().Update(bson.M{"_id": starID}, bson.M{"$push": bson.M{"scenes": s.Id}}); err != nil {
				log.Errorw("cascade failure on star add sceneId",
					"index", i,
					"star_id", starID,
					"scene_id", s.Id,
					"error", err,
				)
				return fmt.Errorf("cascade failure")
			}
		}
	}

	/*  Studio CASCADE */
	if s.Information.Studio != nil {
		if err := connection.Collection("studio").Collection().Update(bson.M{"_id": s.Information.Studio}, bson.M{"$push": bson.M{"scenes": s.Id}}); err != nil {
			log.Errorw("cascade failure on studio add sceneId",
				"studio_id", s.Volume,
				"scene_id", s.Id,
				"error", err,
			)
			return fmt.Errorf("cascade failure")
		}
	}

	if s.Information.Director != nil {
		for i, directorID := range s.Information.Director {
			if err := connection.Collection("star").Collection().Update(bson.M{"_id": directorID}, bson.M{"$push": bson.M{"scenes": s.Id}}); err != nil {
				log.Errorw("cascade failure on star add sceneId",
					"index", i,
					"star_id", directorID,
					"scene_id", s.Id,
					"error", err,
				)
				return fmt.Errorf("cascade failure")
			}
		}
	}

	return nil

}

// AfterDelete ..
func (s *Model) AfterDelete(*bongo.Collection) error {

	/*  VOLUME CASCADE */
	if err := connection.Collection("volume").Collection().Update(bson.M{"_id": s.Volume}, bson.M{"$pull": bson.M{"scenes": s.Id}}); err != nil {
		log.Errorw("cascade failure on volume add sceneId",
			"volume_id", s.Volume,
			"scene_id", s.Id,
			"error", err,
		)
		return fmt.Errorf("cascade failure")
	}

	/*  Stars CASCADE */
	if s.Information.Stars != nil {
		for i, starID := range s.Information.Stars {
			if err := connection.Collection("star").Collection().Update(bson.M{"_id": starID}, bson.M{"$pull": bson.M{"scenes": s.Id}}); err != nil {
				log.Errorw("cascade failure on star add sceneId",
					"index", i,
					"star_id", starID,
					"scene_id", s.Id,
					"error", err,
				)
				return fmt.Errorf("cascade failure")
			}
		}
	}

	/*  Studio CASCADE */
	if s.Information.Studio != nil {
		if err := connection.Collection("studio").Collection().Update(bson.M{"_id": s.Information.Studio}, bson.M{"$pull": bson.M{"scenes": s.Id}}); err != nil {
			log.Errorw("cascade failure on studio add sceneId",
				"studio_id", s.Volume,
				"scene_id", s.Id,
				"error", err,
			)
			return fmt.Errorf("cascade failure")
		}
	}

	if s.Information.Director != nil {
		for i, directorID := range s.Information.Director {
			if err := connection.Collection("star").Collection().Update(bson.M{"_id": directorID}, bson.M{"$pull": bson.M{"scenes": s.Id}}); err != nil {
				log.Errorw("cascade failure on star add sceneId",
					"index", i,
					"star_id", directorID,
					"scene_id", s.Id,
					"error", err,
				)
				return fmt.Errorf("cascade failure")
			}
		}
	}

	return nil

}
