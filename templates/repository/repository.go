package repository

import (
	"context"
	"templates/models"
)

type Repository interface {
	Create(ctx context.Context, animal models.Animal) (*models.Animal, error)
	Get(ctx context.Context, attribute string) (*models.Animal, error)
	GetAll(ctx context.Context) (*models.Animals, error)
	Update(ctx context.Context, attribute string) (*models.Animal, error)
	Delete(ctx context.Context, attribute string) (*models.Animal, error)
}
