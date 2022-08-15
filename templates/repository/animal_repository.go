package repository

import (
	"templates/models"

	"gorm.io/gorm"
)

type AnimalRepository struct {
	connection *gorm.DB
}

func (animalRepository *AnimalRepository) Create() (*models.Animal, error) {
	tx := animalRepository.connection.Begin()
	defer func() {
		if animalRepository := recover(); animalRepository != nil {
			tx.Rollback()
	}

	if err != nil {
		return nil, err
	}
}
