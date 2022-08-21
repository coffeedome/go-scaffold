package main

import (
	"fmt"
	"log"
	"net/http"
	"templates/config"
	"templates/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var DbInstanceRepository repository.PostgresRepository

func main() {

	config.LoadConfig()

	DbInstanceRepository.InitPostgresRepository(config.AppConfig.DbConnectionString)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(middleware.Logger)

	router.Post("/api/animal", CreateAnimal)
	router.Get("/api/animal", GetAnimals)
	router.Get("/api/animal/{id}", GetAnimalById)
	router.Put("/api/animal/{id}", UpdateAnimals)

	log.Printf("Animals API Server listening on port %d...", config.AppConfig.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.AppConfig.ApiPort), router))
}
