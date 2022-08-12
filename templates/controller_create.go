package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func CreateAnimal(w http.ResponseWriter, r *http.Request) {

	animal := Animal{}

	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		log.Fatal("Failed to decode animal. Check that animal attribute types are correct")
	}

	log.Printf("Persisting %s animal record", animal.Name)

	//Persist data here

	json.NewEncoder(w).Encode(&animal)
}
