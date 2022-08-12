package templates

import (
	"encoding/json"
	"log"
	"net/http"
)

type Animal struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Habitat string `json:"habitat"`
	Age     int    `json:"age"`
}

func CreateAnimal(w http.ResponseWriter, r *http.Request) {

	animal := Animal{}

	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		log.Fatal("Failed to decode animal. Check that animal attribute types are correct")
	}

	log.Printf("Persisting %s animal record", animal.Name)

	json.NewEncoder(w).Encode(&animal)
}
