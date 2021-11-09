package db

import (
	"github.com/Nathapot/go-stock/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDB -> Call this method to get db
func GetDB() *gorm.DB {
	return db
}

// SetupDB() -> Setup database for sharing to all api
func SetupDB() {
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&models.User{})
	db = database
}
