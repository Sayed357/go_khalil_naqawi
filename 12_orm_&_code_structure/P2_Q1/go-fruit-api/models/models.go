package models

import (
	"github.com/google/uuid"
)

type NutritionalData struct {
	Calories      int `json:"calories"`
	Fat           int `json:"fat"`
	Sugar         int `json:"sugar"`
	Carbohadrates int `json:"carbohydrates"`
	Protein       int `json:"protein"`
}

type Fruit struct {
	ID              uuid.UUID       `json:"id"`
	Name            string          `json:"name"`
	Price           int             `json:"price"`
	NutritionalData NutritionalData `json:"nutritional_data"`
}
