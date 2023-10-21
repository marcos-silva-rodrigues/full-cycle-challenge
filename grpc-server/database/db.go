package database

import (
	"log"

	"github.com/full-cycle-challenge/model"
	"github.com/jinzhu/gorm"
	_ "gorm.io/driver/sqlite"
)

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic(err)
	}

	db.LogMode(true)

	db.AutoMigrate(&model.Product{})

	return db
}
