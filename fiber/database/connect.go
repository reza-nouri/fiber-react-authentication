package database

import (
	"fiberAuthentication/models"

	"gorm.io/driver/postgres"
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var DB *gorm.DB

// github.com/mattn/go-sqlite3
func Connect() {
	dsn := "connection string"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to database")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Could not auto migrate")
	}

	DB = db
}
