package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/notify-me/server/models"
)

type SenderDTO struct {
	IP string `json:"IP" binding:"required"`
}

func GetSenders(c *gin.Context) {
	senders := models.ListSenders()
	c.JSON(200, gin.H{
		"data": senders,
	})
}

func PostSender(c *gin.Context) {
	var input SenderDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sender, err := models.CreateSender(models.Sender{
		IP:        input.IP,
		IsEnabled: true,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": sender,
	})
}
