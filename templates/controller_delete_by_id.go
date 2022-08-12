package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func DeleteAnimalById(w http.ResponseWriter, r *http.Request) {

	animalID := chi.URLParam(r, "animalID")

	log.Printf("Deleting record for animal id %s", animalID)

	animalUpdates := Animal{}

	//Get animal to update into currentAnimalRecord var

	json.NewDecoder(r.Body).Decode(&animalUpdates)
}
