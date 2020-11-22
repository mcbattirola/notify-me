package handler

import (
	"net/http"

	"github.com/mcbattirola/notify-me/server/models"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoint(c *gin.Context) {
	id := c.Param("id")
	// add in db

	device := models.Device{
		ID:         id,
		Is_Enabled: true,
	}

	result := models.DB.Create(&device)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"error": result.Error.Error(),
		})
	} else {
		c.JSON(http.StatusOK, map[string]string{
			"ID": id,
		})
	}

	// return ok or error msmg
}
