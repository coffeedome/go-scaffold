package main

import (
	"encoding/json"
	"net/http"
)

func ListAnimals(w http.ResponseWriter, r *http.Request) {

	animals := []Animal{}

	//Get data here

	json.NewEncoder(w).Encode(&animals)

}
