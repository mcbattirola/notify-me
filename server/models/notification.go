package models

import (
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Title    string
	Body     string
	SenderID uint `json:"sender_id" binding:"required" gorm:"TYPE:integer REFERENCES senders"`
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
