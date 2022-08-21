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
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	log.Printf("Persisting %s animal record", animal.Name)

	animalOut, err := DbInstanceRepository.CreateAnimal(r.Context(), animal)
	if err != nil {
		log.Printf("Failed to create animal %s : %s", animal.Name, err)
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&animalOut)
	log.Printf("Animal %s has been created", animal.ID)
}
