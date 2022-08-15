package main

import (
	"encoding/json"
	"net/http"
)

func UpdateAnimals(w http.ResponseWriter, r *http.Request) {

	animalUpdates := Animal{}

	//Get animal(s) to update into currentAnimalRecord var

	json.NewDecoder(r.Body).Decode(&animalUpdates)
}
