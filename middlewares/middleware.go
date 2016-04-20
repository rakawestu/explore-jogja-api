package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AccessKey access key
var AccessKey string

// CheckHeaders check header before continuing process
func CheckHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
		   For now we just check if current request has valid
		   parameter from header.
		*/
		headers := c.Request.Header
		accessKey := headers.Get("Explore-Jogja-Auth")
		if accessKey != AccessKey {
			c.AbortWithStatus(http.StatusUnauthorized)
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Invalid access key or access key not found."})
		}

		c.Next()

	}
}
