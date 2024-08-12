package main

import (
	"client-data-api/database"
	"client-data-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
