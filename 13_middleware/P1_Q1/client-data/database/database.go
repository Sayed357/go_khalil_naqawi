package database

import (
	"client-data/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = "root"
	DB_PASSWORD string = ""
	DB_NAME     string = "clientdata_db"
	DB_HOST     string = "localhost"
	DB_PORT     string = "3306"
)

func Initdatabase() {
	var err error

	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error connecting to the database: %s\n", err)
	}

	log.Println("successfully connected to the database")

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Content{})

	if err != nil {
		log.Fatalf("error when migrating: %s\n", err)
	}

	log.Println("migration successful")
}
