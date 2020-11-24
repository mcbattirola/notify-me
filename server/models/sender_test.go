package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	println("test main")
	ConnectDataBase("sender_test.db")

	os.Exit(m.Run())
}

func TestListSenders(t *testing.T) {
	// It should return only enabled senders
	senders := ListSenders()
	for _, sender := range senders {
		if !sender.IsEnabled {
			t.Fatalf("Expected no disabled senders")
		}
	}
}

func TestCreateSender(t *testing.T) {
	sender := Sender{
		ID:        1,
		IsEnabled: true,
		IP:        "127.0.0.1",
	}

	_, err := CreateSender(sender)

	// No errors
	if err != nil {
		t.Fatalf("Expected no errors, got %s", err.Error())
	}

	// The sender is persisted
	newSender := GetSenderById(sender.ID)

	if !fieldsMatch(sender, newSender) {
		t.Fatalf("Expected to find the sender with same fields, but got a different one")
	}

}

func fieldsMatch(sender1 Sender, sender2 Sender) bool {
	return (sender1.ID == sender2.ID) && (sender1.IP == sender2.IP) && (sender1.IsEnabled == sender2.IsEnabled)
}
