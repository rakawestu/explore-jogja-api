package main

import (
	"github.com/gin-gonic/gin"
)

// APIRoutes is endpoint group for API
func APIRoutes(router *gin.RouterGroup) {
	router.GET("/places", GetAllPlacesHandler)
	router.GET("/places/:id", GetSinglePlaceHandler)
	router.POST("/places", InsertPlaceHandler)
	router.DELETE("/places/:id", DeletePlaceHandler)
	router.PUT("/places/:id", UpdatePlaceHandler)
}
