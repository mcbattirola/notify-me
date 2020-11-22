package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase(dbName string) {
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Printf("\n\nConnected to database %s.\n\n", dbName)

	database.AutoMigrate(&Device{})

	DB = database
}
