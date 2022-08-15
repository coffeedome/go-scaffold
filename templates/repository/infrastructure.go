package repository

import (
	"strconv"
	"templates/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AnimalGorm struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Type    string `gorm:"column:type"`
	Habitat string `gorm:"column:habitat"`
	Age     int    `gorm:"column:age"`
}

func (animalGorm *AnimalGorm) ToEntity() (*models.Animal, error) {

	animalGormId := strconv.FormatUint(uint64(animalGorm.ID), 10)

	parsed, err := uuid.Parse(animalGormId)
	if err != nil {
		return &models.Animal{}, err
	}

	return &models.Animal{
		ID:      parsed,
		Name:    animalGorm.Name,
		Type:    animalGorm.Type,
		Habitat: animalGorm.Habitat,
		Age:     animalGorm.Age,
	}, nil
}
