package models

import "testing"

func TestCreateNotification(t *testing.T) {
	notification := Notification{
		Title: "Title",
		Body:  "Body",
	}

	createdNotification, err := CreateNotification(notification)

	// No errors
	if err != nil {
		t.Fatalf("Expected no errors, got %s", err.Error())
	}

	// The notification is persisted
	newNotification := GetNotificationById(createdNotification.ID)

	if !notificationFieldsMatch(notification, newNotification) {
		t.Fatalf("Expected to find the notification with same fields, but got a different one")
	}

}

func notificationFieldsMatch(notification1 Notification, notification Notification) bool {
	return (notification1.Title == notification.Title) && (notification1.Body == notification.Body)
}
