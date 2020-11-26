package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/notify-me/server/models"
	fcm "github.com/mcbattirola/notify-me/server/services"
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

	// get senderID by IP
	clientIP := c.ClientIP()
	sender, err := models.GetSenderByIp(clientIP)

	// return error if sender doesnt exist
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	notification, err := models.CreateNotification(models.Notification{
		Title:    input.Title,
		Body:     input.Body,
		SenderID: sender.ID,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get all subscription with sender id = notification.SenderID
	subscriptions, _ := models.GetSubscriptionsBySenderID(notification.SenderID)

	for _, subscription := range subscriptions {
		go fcm.SendPushNotification(notification, subscription.DeviceID)
	}

	c.JSON(200, gin.H{
		"data": notification,
	})
}
