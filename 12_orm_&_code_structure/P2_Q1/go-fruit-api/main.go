package main

import (
	"database/sql"
	"go-fruit-api/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	router := gin.Default()

	db, _ = sql.Open("sqlite3", "./fruits.db")

	router.GET("/api/v1/fruits", controllers.GetFruits)
	router.GET("/api/v1/fruits/:id", controllers.GetFruitByID)
	router.POST("/api/v1/fruits", controllers.CreateFruit)
	router.DELETE("/api/v1/fruits/:id", controllers.DeleteFruit)

	router.Run(":2313")
}
