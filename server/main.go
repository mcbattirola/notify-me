package main

import (
	"notify-me/handler"
)

const DB_NAME string = "notifyme.db"

func main() {
	handler.SetupServer(DB_NAME).Run()
}
