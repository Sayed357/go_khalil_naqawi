package main

import (
	"go-wishlist-api/controllers"
	"go-wishlist-api/database"
	"go-wishlist-api/repository"
	"go-wishlist-api/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()

	wishlistRepo := &repository.GormWishlistRepository{DB: database.DB}
	wishlistUseCase := &usecases.WishlistUseCase{Repo: wishlistRepo}
	wishlistController := &controllers.WishlistController{UseCase: wishlistUseCase}

	e := echo.New()

	e.GET("/wishlists", wishlistController.GetAll)
	e.POST("/wishlists", wishlistController.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
