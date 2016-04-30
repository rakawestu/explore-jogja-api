package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rakawestu/explore-jogja-api/handlers"
)

// ReviewRoutes is endpoint group for API
func ReviewRoutes(router *gin.RouterGroup) {
	router.GET("/places/:id/reviews", handlers.GetPlaceReviews)
	router.POST("/places/:id/reviews", handlers.InsertReviewForPlace)
}
