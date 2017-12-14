package exercise

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// Index GET /exercises
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	user := "232124124"                                //r.Body
	exercises := c.Repository.GetExercisesInYear(user) // list of all exercises within a year
	log.Println(exercises)
	data, _ := json.Marshal(exercises)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddExercise POST /
func (c *Controller) AddExercise(w http.ResponseWriter, r *http.Request) {

	var exercise Exercise
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Fatalln("Error AddExercise", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddExercise", err)
	}
	if err := json.Unmarshal(body, &exercise); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddExercise unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	//exercise.User = r.user._id
	success := c.Repository.AddExercise(exercise) // adds the exercise to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}
