package models

import (
	"github.com/google/uuid"
)

type Animal struct {
	ID      uuid.UUID
	Name    string
	Type    string
	Habitat string
	Age     int
}

type Animals []Animal
