package main

import (
	"fmt"
	"log"
	"net/http"
	"templates/config"
	"templates/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var DbInstanceRepository repository.PostgresRepository

func main() {

	config.LoadConfig()

	DbInstanceRepository.InitPostgresRepository(config.AppConfig.DbConnectionString)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Post("/api/animal", CreateAnimal)
	router.Get("/api/animal", GetAnimals)
	router.Get("/api/animal/{id}", GetAnimalById)
	router.Put("/api/animal/{id}", UpdateAnimals)

	log.Printf("Animals API Server listening on port %d...", config.AppConfig.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.ApiPort), router))
}
