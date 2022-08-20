package main

import (
	"encoding/json"
	"log"
	"net/http"
	"templates/models"
)

func CreateAnimal(w http.ResponseWriter, r *http.Request) {

	var animal models.Animal

	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		log.Fatal("Failed to decode animal. Check that animal attribute types are correct")
	}

	log.Printf("Persisting %s animal record", animal.Name)

	DbInstanceRepository.CreateAnimal(r.Context(), animal)

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&animal)
	log.Printf("Animal %s has been created", animal.ID)
}
