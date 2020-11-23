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
		if !sender.Is_Enabled {
			t.Fatalf("Expected no disabled senders")
		}
	}
}
