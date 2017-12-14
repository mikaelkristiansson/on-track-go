package exercise

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/now"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "on-track-service"

// DOCNAME the name of the document
const DOCNAME = "exercise"

// GetExercisesInYear returns the list of Exercises within a year
func (r Repository) GetExercisesInYear(user string) Exercises {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	query := bson.M{"user": user, "created_at": bson.M{
		"$gt": now.BeginningOfYear(),
		"$lt": time.Now(),
	}}
	results := Exercises{}
	if err := c.Find(query).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// AddExercise inserts an Exercise in the DB
func (r Repository) AddExercise(exercise Exercise) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	exercise.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(exercise)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
