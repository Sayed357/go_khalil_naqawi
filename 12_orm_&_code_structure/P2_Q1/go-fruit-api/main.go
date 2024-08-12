package main

import (
	"go-fruit-api/controllers"
	"go-fruit-api/database"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	router := gin.Default()

	database.Initdatabase()

	router.GET("/api/v1/fruits", controllers.GetFruits)
	router.GET("/api/v1/fruits/:id", controllers.GetFruitByID)
	router.POST("/api/v1/fruits", controllers.CreateFruit)
	//router.DELETE("/api/v1/fruits/:id", controllers.DeleteFruit)

	router.Run(":1323")
}
