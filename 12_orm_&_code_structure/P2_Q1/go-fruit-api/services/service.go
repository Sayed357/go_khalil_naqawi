package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"fruit-api/models"
	"net/http"

	"github.com/google/uuid"
)

func FetchNutritionalData(fruitName string) (models.NutritionalData, error) {
	var nutritionalData models.NutritionalData
	resp, err := http.Get(`https://www.fruityvice.com/api/fruit/` + fruitName)
	if err != nil {
		return nutritionalData, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nutritionalData, fmt.Errorf("failed to fetch nutritional data: status code %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&nutritionalData); err != nil {
		return nutritionalData, err
	}

	return nutritionalData, nil
}

func AddFruit(db *sql.DB, name string, price int) (models.Fruit, error) {
	nutritionalData, err := FetchNutritionalData(name)
	if err != nil {
		return models.Fruit{}, err
	}

	fruit := models.Fruit{
		ID:              uuid.New(),
		Name:            name,
		Price:           price,
		NutritionalData: nutritionalData,
	}

	_, err = db.Exec("INSERT INTO fruits (id, name, price, calories, fat, sugar, carbohydrates, protein) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		fruit.ID, fruit.Name, fruit.Price, fruit.NutritionalData.Calories, fruit.NutritionalData.Fat, fruit.NutritionalData.Sugar, fruit.NutritionalData.Carbohydrates, fruit.NutritionalData.Protein)
	if err != nil {
		return fruit, err
	}

	return fruit, nil
}
