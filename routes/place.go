package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/handlers"
)

// PlaceRoutes is endpoint group for API
func PlaceRoutes(router *gin.RouterGroup) {
	router.GET("/places", handlers.GetAllPlacesHandler)
	router.GET("/places/:id", handlers.GetSinglePlaceHandler)
	router.POST("/places", handlers.InsertPlaceHandler)
	router.DELETE("/places/:id", handlers.DeletePlaceHandler)
	router.PUT("/places/:id", handlers.UpdatePlaceHandler)
	router.GET("/categories", handlers.GetAllCategoriesHandler)
	router.POST("/categories", handlers.InsertCategoryHandler)
}
