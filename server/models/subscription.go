package models

type Subscription struct {
	DeviceID string `json:"device_id" binding:"required" gorm:"TYPE:integer REFERENCES devices"`
	SenderID uint   `json:"sender_id" binding:"required" gorm:"TYPE:integer REFERENCES senders"`
}

func CreateSubscription(subscription Subscription) (Subscription, error) {
	result := DB.Create(&subscription)

	return subscription, result.Error
}

func GetSubscriptionsBySenderID(senderID uint) ([]Subscription, error) {
	subscriptions := make([]Subscription, 0)

	result := DB.Where(&Subscription{
		SenderID: senderID,
	}).Find(&subscriptions)

	return subscriptions, result.Error
}

func ListSubscriptions() []Subscription {
	var subscriptions []Subscription

	DB.Find(&subscriptions)

	return subscriptions
}
