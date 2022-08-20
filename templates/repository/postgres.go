package repository

import (
	"context"
	"templates/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	connection *gorm.DB
}

func (repository *PostgresRepository) InitPostgresRepository(dbConnectionString string) error {
	var err error

	repository.connection, err = gorm.Open(postgres.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		return err
	}

	return nil
}

func (repository *PostgresRepository) CreateAnimal(ctx context.Context, animal models.Animal) (*models.Animal, error) {
	tx := repository.connection

	err := tx.Create(&animal).Error
	if err != nil {
		return nil, err
	}

	return &animal, nil
}

func (repository *PostgresRepository) GetAnimalById(ctx context.Context, animalId string) (*models.Animal, error) {
	tx := repository.connection

	animal := models.Animal{
		ID: animalId,
	}

	err := tx.First(&animal).Error
	if err != nil {
		return nil, err
	}

	return &animal, nil
}

func (repository *PostgresRepository) GetAnimals(ctx context.Context) ([]*models.Animal, error) {
	var animals []*models.Animal
	tx := repository.connection

	err := tx.Find(&animals).Error
	if err != nil {
		return nil, err
	}

	return animals, nil
}

func (repository *PostgresRepository) UpdateAnimal(ctx context.Context, animalUpdates models.Animal) (*models.Animal, error) {
	tx := repository.connection

	var animalToUpdate models.Animal
	animalToUpdate.ID = animalUpdates.ID

	err := tx.First(&animalToUpdate).Error
	if err != nil {
		return nil, err
	}

	animalToUpdate.Name = animalUpdates.Name
	animalToUpdate.Type = animalUpdates.Type
	animalToUpdate.Habitat = animalUpdates.Habitat
	animalToUpdate.Age = animalUpdates.Age

	err = tx.Save(&animalToUpdate).Error
	if err != nil {
		return nil, err
	}

	return &animalToUpdate, nil
}

func (repository *PostgresRepository) DeleteAnimal(ctx context.Context, animal models.Animal) (*models.Animal, error) {
	tx := repository.connection

	err := tx.Delete(&animal).Error
	if err != nil {
		return nil, err
	}

	return &animal, nil
}
