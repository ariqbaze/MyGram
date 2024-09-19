package database

import (
	"MyGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", "localhost", "postgres", "postgres", "postgres", "5432")

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Success connected to database")
	db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
