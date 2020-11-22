package main

import (
	"notify-me/handler"
)

const DB_NAME string = "notifyme"

func main() {
	handler.SetupServer(DB_NAME).Run()
}
