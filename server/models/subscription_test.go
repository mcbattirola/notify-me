package models

import "testing"

func TestCreateSubscription(t *testing.T) {
	subscription := Subscription{
		DeviceID: "1",
		SenderID: 1,
	}

	_, err := CreateSubscription(subscription)

	if err != nil {
		t.Fatalf("Expected no errors, got %s", err.Error())
	}
}
