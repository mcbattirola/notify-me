package models

type Subscription struct {
	DeviceID string `json:"device_id" binding:"required" gorm:"TYPE:integer REFERENCES devices"`
	SenderID uint   `json:"sender_id" binding:"required" gorm:"TYPE:integer REFERENCES senders"`
}

func CreateSubscription(subscription Subscription) (Subscription, error) {
	result := DB.Create(&subscription)

	return subscription, result.Error
}
