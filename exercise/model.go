package exercise

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Exercise connected to a user
type Exercise struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	User      bson.ObjectId `json:"user"`
	Type      string        `json:"type"`
	CreatedAt time.Time     `json:"created_at"`
	DeletedAt time.Time     `json:"deleted_at"`
}

//Exercises is an array of Exercise
type Exercises []Exercise
