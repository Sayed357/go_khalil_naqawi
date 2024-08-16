package main

import (
	"managing_data/config"
	"managing_data/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := config.SetupDatabase()
	defer db.Close()

	routes.RegisterRoutes(r)

	r.Run(":1323")
}
