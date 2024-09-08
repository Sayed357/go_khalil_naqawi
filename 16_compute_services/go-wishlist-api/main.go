package main

import (
	"go-wishlist-api/controllers"
	"go-wishlist-api/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDatabase()

	e := echo.New()

	e.GET("/wishlists", controllers.GetAll)
	e.POST("/wishlists", controllers.Create)

	e.Logger.Fatal(e.Start(":1323"))
}
