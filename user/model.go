package user

import "time"

//User model
type User struct {
	Username  string    `json:"username" required:"true"`
	Password  string    `bson:",omitempty"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

//Let's try an index
// index := mgo.Index{
// 	Key: []string{"uniqid"},
// 	Unique: true,
// }
// err = coll.EnsureIndex(index)
// if err != nil {
// 	log.Fatal(err)
// }
