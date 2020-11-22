package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/notify-me/server/handler"
	"github.com/mcbattirola/notify-me/server/models"
)

const DB_NAME = "testdb.db"

func main() {
	setupServer().Run()
}

// The engine with all endpoints is now extracted from the main function
func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.PingEndpoint)
	r.GET("/register/:id", handler.RegisterEndpoint)

	models.ConnectDataBase(DB_NAME)

	return r
}
