package models

import "context"

type AnimalRepository interface {
	CreateAnimal(ctx context.Context, animal Animal) (*Animal, error)
	GetAnimalById(ctx context.Context, animalId string) (*Animal, error)
	GetAnimals(ctx context.Context) ([]*Animal, error)
	UpdateAnimal(ctx context.Context, animalUpdates Animal) (*Animal, error)
	DeleteAnimal(ctx context.Context, animal Animal) (*Animal, error)
}
