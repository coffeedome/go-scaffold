package main

import (
	"encoding/json"
	"log"
	"net/http"
	"templates/models"
)

func DeleteAnimal(w http.ResponseWriter, r *http.Request) {

	var animal models.Animal

	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	log.Printf("Deleting record for animal id %s", animal.ID)

	animalOut, err := DbInstanceRepository.DeleteAnimal(r.Context(), animal)
	if err != nil {
		log.Printf("Failed to delete animal record: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(&animalOut)
	log.Printf("Animal id %s has been deleted", animal.ID)

}
