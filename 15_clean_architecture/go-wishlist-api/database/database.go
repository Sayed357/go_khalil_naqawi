package database

import (
	"go-wishlist-api/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// connect to the database
func InitDatabase() {
	var err error

	DB, err = gorm.Open(sqlite.Open("advertising_db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Wishlist{})
}
