package controllers

import (
	"go-wishlist-api/models"
	usecases "go-wishlist-api/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WishlistController struct {
	UseCase *usecases.WishlistUseCase
}

func (wc *WishlistController) GetAll(c echo.Context) error {
	wishlists, err := wc.UseCase.GetAllWishlists()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to retrieve wishlists"})
	}
	return c.JSON(http.StatusOK, echo.Map{"data": wishlists})
}

func (wc *WishlistController) Create(c echo.Context) error {
	var input models.Wishlist
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	if err := wc.UseCase.CreateWishlist(input); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Unable to create wishlist"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"data": input})
}
