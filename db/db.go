package db

import (
	"book_api/pkg/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Init() *gorm.DB {
	db, err := gorm.Open(os.Getenv("DB_TYPE"), fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	))

	if err != nil {
		log.Fatal(err, "unable to connect to database")
	}
	log.Printf("database setupdone")
	db.Debug().AutoMigrate(models.Book{})
	return db
}
