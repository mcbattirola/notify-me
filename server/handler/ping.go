package handler

import "github.com/gin-gonic/gin"

func PingEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
