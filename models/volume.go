/*
 * Vuli API
 *
 * Vuli Movie Delivery API
 *
 * API version: 3

 */

package models

import (
	"fmt"
	"strconv"

	"github.com/go-bongo/bongo"
	"gopkg.in/mgo.v2/bson"
)

// Volume Document
//
// A volume can be associated with a Series
//
// swagger:model
type Volume struct {
	bongo.DocumentBase `bson:",inline"`

	// Unique Title for this scene
	Title string `json:"title"`

	// Unique Slug for this scene. Made of <title><studio> lowercase and character stripped
	Slug string `json:"slug"`

	// Description of this scene if it has one. Not required
	Description string `json:"description"`

	// Series this volume is in
	Series *bson.ObjectId `json:"series"`

	// Series this volume is in
	Scenes []bson.ObjectId `json:"scenes"`

	// MovieInformation --
	Information MediaInformation `json:"information"`

	// Volume Viewing Performance
	Performance Performance `json:"performance"`

	// Media information
	Images     Images     `json:"images"`
	Extras     []Extras   `json:"extras"`
	Thumbnails Thumbnails `json:"thumbnails"`
	Trailer    Trailer    `json:"trailer"`

	// List of Tags
	Tags []string `json:"tags"`

	// Read only value. Only Admin can update. Sets the price for a the volume which supersedes the scene price
	Price float64 `json:"price"`

	// True/False. Is it available on the site or not
	IsPublished bool `json:"is_published"`

	diffTracker *bongo.DiffTracker
}

// GetDiffTracker ..
func (v *Volume) GetDiffTracker() *bongo.DiffTracker {
	if v.diffTracker == nil {
		v.diffTracker = bongo.NewDiffTracker(v)
	}

	return v.diffTracker

}

// Validate --
func (v *Volume) Validate(*bongo.Collection) []error {

	retval := make([]error, 0)
	// volume := &Volume{}

	// Check for series
	if v.Series != nil {
		if !v.Series.Valid() {
			retval = append(retval, fmt.Errorf("not a valid series objectId"))
		}
	}

	// Check for studio
	if v.Information.Studio == nil {
		retval = append(retval, fmt.Errorf("studio cannot be empty"))

	} else {
		if !v.Information.Studio.Valid() {
			retval = append(retval, fmt.Errorf("not a valid studio objectId"))
		}
	}

	// Check for bad IDs
	for i, e := range v.Information.Director {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("director id is not valid in position: "+strconv.Itoa(i)))
		}
	}

	// Check for bad IDs
	for i, e := range v.Information.Stars {
		if !e.Valid() {
			retval = append(retval, fmt.Errorf("star id is not valid in position: "+strconv.Itoa(i)))
		}
	}
	// Find by slug when posting new volume
	// err := connection.Collection("volume").FindOne(bson.M{"slug": v.Slug}, volume)

	// if err == nil {
	// retval = append(retval, fmt.Errorf("this document is not unique (via slug)"))
	// }

	// log.Debugw("error saving volume", "error", retval)
	return retval
}

// AfterSave ..
func (v *Volume) AfterSave(*bongo.Collection) error {

	/*  Series CASCADE */
	if v.Series != nil {
		if err := connection.Collection("series").Collection().Update(bson.M{"_id": v.Series}, bson.M{"$push": bson.M{"volumes": v.Id}}); err != nil {
			log.Errorw("cascade failure on series add sceneId",
				"series_id", v.Series,
				"volume_id", v.Id,
				"error", err,
			)
			return fmt.Errorf("cascade failure")
		}
	}

	return nil

}

// AfterDelete ..
func (v *Volume) AfterDelete(*bongo.Collection) error {

	/*  Series CASCADE */
	if v.Series != nil {
		if err := connection.Collection("series").Collection().Update(bson.M{"_id": v.Series}, bson.M{"$pull": bson.M{"volumes": v.Id}}); err != nil {
			log.Errorw("cascade failure on series add sceneId",
				"series_id", v.Series,
				"volume_id", v.Id,
				"error", err,
			)
			return fmt.Errorf("cascade failure")
		}
	}
	return nil

}
