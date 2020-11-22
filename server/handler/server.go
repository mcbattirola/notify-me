package handler

import (
	"notify-me/platform/models"

	"github.com/gin-gonic/gin"
)

func SetupServer(dbName string) *gin.Engine {
	r := gin.Default()

	r.GET("/api/ping", pingGet())
	r.GET("/api/v1/register/:token", register())

	models.ConnectDataBase(dbName)

	return r
}
