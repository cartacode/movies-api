package main

import (
	"fmt"
	"log"

	"github.com/VuliTv/go-movie-api/dbh"
	"github.com/VuliTv/go-movie-api/models"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	mDB, dbError := dbh.NewMongoDBConnection("convert")

	if dbError != nil {
		log.Fatal(dbError)
	}
	results := mDB.Collection("scene").Find(bson.M{"volume": bson.M{"$exists": true}})

	scene := &models.Scene{}
	for results.Next(scene) {
		fmt.Println(scene.Volume)
	}

}
