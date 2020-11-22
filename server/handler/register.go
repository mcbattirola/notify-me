package handler

import (
	"net/http"
	"notify-me/platform/models"

	"github.com/gin-gonic/gin"
)

func register() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")

		// add in db
		device := models.Device{Token: token, Is_enabled: true}

		result := models.DB.Create(&device)

		if result.Error != nil {
			c.JSON(http.StatusBadRequest, map[string]string{
				"error": result.Error.Error(),
			})
		} else {
			c.JSON(http.StatusOK, map[string]string{
				"token": token,
			})
		}
	}
}
