package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/api/animal", CreateAnimal)
	router.Get("/api/animal", ListAnimals)
	router.Get("/api/animal/{id}", GetAnimalById)
	router.Put("/api/animal/{id}", UpdateAnimals)

	log.Print("Animals API Server listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
