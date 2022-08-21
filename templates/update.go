package main

import (
	"encoding/json"
	"log"
	"net/http"
	"templates/models"
)

func UpdateAnimals(w http.ResponseWriter, r *http.Request) {

	var animal models.Animal

	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	animalOut, err := DbInstanceRepository.UpdateAnimal(r.Context(), animal)
	if err != nil {
		log.Printf("Failed to update animal record: %s", animal.ID)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	log.Printf("Animal id %s updated", animal.ID)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&animalOut)
}
