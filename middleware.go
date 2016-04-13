package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHeaders check header before continuing process
func CheckHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
		   For now we just check if current request has valid
		   parameter from header.  We're not
		   check the value, just check if parameter exists or not.
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
