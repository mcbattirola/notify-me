package models

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ConnectDataBase("models-test.db")

	os.Exit(m.Run())
}
