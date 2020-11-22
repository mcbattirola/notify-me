package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func register() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		// add in db

		// return ok or error msmg
		c.JSON(http.StatusOK, map[string]string{
			"ok": token,
		})
	}
}
