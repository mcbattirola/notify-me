package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/notify-me/server/handler"
	"github.com/mcbattirola/notify-me/server/models"
)

const dbName = "prod.db"

func main() {
	setupServer().Run()
}

func setupServer() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handler.PingEndpoint)
	r.GET("/register/:id", handler.RegisterEndpoint)
	r.GET("/sender", handler.GetSenders)
	r.POST("/sender", handler.PostSender)
	r.POST("/subscribe", handler.Subscribe)
	r.POST("/notification", handler.PostNotification)

	models.ConnectDataBase(dbName)

	return r
}
