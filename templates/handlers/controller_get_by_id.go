package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func GetAnimalById(w http.ResponseWriter, r *http.Request) {

	animalID := chi.URLParam(r, "animalID")

	log.Printf("Getting record for animal id %s", animalID)

	animal := Animal{}

	//Get data here

	json.NewEncoder(w).Encode(&animal)

}
