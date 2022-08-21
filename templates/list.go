package main

import (
	"encoding/json"
	"log"
	"net/http"
	"templates/models"
)

func GetAnimals(w http.ResponseWriter, r *http.Request) {

	var animals []models.Animal

	err := json.NewDecoder(r.Body).Decode(&animals)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	animalsOut, err := DbInstanceRepository.GetAnimals(r.Context())
	if err != nil {
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&animalsOut)
	log.Print("Retrieved all animals")

}
