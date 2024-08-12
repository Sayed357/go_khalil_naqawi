package controllers

import (
	"database/sql"
	"fruit-api/models"
	"fruit-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFruits(c *gin.Context, db *sql.DB) {
	var fruits []models.Fruit
	rows, err := db.Query("SELECT * FROM fruits")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query fruits"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var fruit models.Fruit
		if err := rows.Scan(&fruit.ID, &fruit.Name, &fruit.Price, &fruit.NutritionalData.Calories, &fruit.NutritionalData.Fat, &fruit.NutritionalData.Sugar, &fruit.NutritionalData.Carbohydrates, &fruit.NutritionalData.Protein); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan fruit"})
			return
		}
		fruits = append(fruits, fruit)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred during rows iteration"})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

func GetFruitByID(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	var fruit models.Fruit
	row := db.QueryRow("SELECT * FROM fruits WHERE id = ?", id)
	err := row.Scan(&fruit.ID, &fruit.Name, &fruit.Price, &fruit.NutritionalData.Calories, &fruit.NutritionalData.Fat, &fruit.NutritionalData.Sugar, &fruit.NutritionalData.Carbohydrates, &fruit.NutritionalData.Protein)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "fruit not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve fruit"})
		}
		return
	}

	c.JSON(http.StatusOK, fruit)
}

func CreateFruit(c *gin.Context, db *sql.DB) {
	var fruitInput struct {
		Name  string `json:"name" binding:"required"`
		Price int    `json:"price" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&fruitInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fruit, err := services.AddFruit(db, fruitInput.Name, fruitInput.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, fruit)
}
