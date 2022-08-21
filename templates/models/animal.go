package models

type Animal struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Habitat string `json:"habitat"`
	Age     int    `json:"age"`
}
