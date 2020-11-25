package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/notify-me/server/models"
)

type NotificationDTO struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

func PostNotification(c *gin.Context) {
	var input NotificationDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification, err := models.CreateNotification(models.Notification{
		Title: input.Title,
		Body:  input.Body,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"data": notification,
	})
}
