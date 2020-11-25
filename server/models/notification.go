package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Title string
	Body  string
}

func CreateNotification(notification Notification) (Notification, error) {
	result := DB.Create(&notification)

	return notification, result.Error
}

func GetNotificationById(id uint) Notification {
	notification := Notification{}
	DB.First(&notification, id)
	return notification
}
