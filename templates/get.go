package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func GetAnimalById(w http.ResponseWriter, r *http.Request) {

	animalID := chi.URLParam(r, "id")

	log.Printf("Getting record for animal id %s", animalID)

	animal, err := DbInstanceRepository.GetAnimalById(r.Context(), animalID)
	if err != nil {
		w.WriteHeader(http.StatusAccepted)

	}

	json.NewEncoder(w).Encode(&animal)

}
