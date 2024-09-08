package controllers

import (
	"go-wishlist-api/database"
	"go-wishlist-api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAll(c echo.Context) error {
	var wishlists []models.Wishlist

	database.DB.Find(&wishlists)

	return c.JSON(http.StatusOK, echo.Map{
		"data": wishlists,
	})
}

func Create(c echo.Context) error {
	var input models.Wishlist

	c.Bind(&input)

	database.DB.Create(&input)

	return c.JSON(http.StatusCreated, echo.Map{
		"data": input,
	})
}
