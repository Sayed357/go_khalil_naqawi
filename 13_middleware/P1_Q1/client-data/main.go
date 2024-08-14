package main

import (
	"client-data/database"
	"client-data/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Initdatabase()

	e := echo.New()

	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
