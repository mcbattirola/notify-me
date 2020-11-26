package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mcbattirola/notify-me/server/models"
)

func Subscribe(c *gin.Context) {

	var input models.Subscription

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	subscription, err := models.CreateSubscription(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]models.Subscription{
		"data": subscription,
	})
}

func GetSubscriptions(c *gin.Context) {
	senders := models.ListSubscriptions()
	c.JSON(200, gin.H{
		"data": senders,
	})
}
