package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Food struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}

var foods = []Food{}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hi lovely",
		})
	})

	e.GET("/api/v1/foods", getAllFoods)
	e.GET("/api/v1/foods/:id", getFoodByID)
	e.POST("/api/v1/foods", addFood)
	e.PUT("/api/v1/foods/:id", updateFood)
	e.DELETE("/api/v1/foods/:id", deleteFood)

	e.Logger.Fatal(e.Start(":1323"))
}

func getAllFoods(c echo.Context) error {
	return c.JSON(http.StatusOK, foods)
}

func getFoodByID(c echo.Context) error {
	id := c.Param("id")
	for _, food := range foods {
		if food.ID == id {
			return c.JSON(http.StatusOK, food)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "the seeking food not found"})

}

func addFood(c echo.Context) error {
	var food Food
	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	food.ID = uuid.New().String()
	foods = append(foods, food)
	return c.JSON(http.StatusCreated, food)
}

func updateFood(c echo.Context) error {
	id := c.Param("id")
	var updatedFood Food
	if err := c.Bind(&updatedFood); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "INvalid request"})
	}
	for i, food := range foods {
		if food.ID == id {
			foods[i].Name = updatedFood.Name
			foods[i].Price = updatedFood.Price
			foods[i].Description = updatedFood.Description
			return c.JSON(http.StatusOK, foods[i])
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Food not found"})
}

func deleteFood(c echo.Context) error {
	id := c.Param("id")
	for i, food := range foods {
		if food.ID == id {
			foods = append(foods[:i], foods[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}
	return c.JSON(http.StatusNotFound, map[string]string{"message": "Food not found"})
}
