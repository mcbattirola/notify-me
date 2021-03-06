package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDataBase(dbName string) {
	database, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Device{})
	database.AutoMigrate(&Sender{})
	database.AutoMigrate(&Subscription{})
	database.AutoMigrate(&Notification{})

	DB = database
}
